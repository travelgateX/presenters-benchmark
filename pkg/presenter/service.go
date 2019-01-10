package presenter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"hub-aggregator/common/stats"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Input struct {
	// OptionNumber are the number of options that each operation is going to return
	OptionNumber int
	// Scale of fields in the graph Request
	// low: options with 1 field
	// medium: options with half the fields
	// high: options with all the fields
	ResolveScale ResolveScale
	// Number of operations executed in the benchmark
	OperationNumber int
	// Operations will be sent in this parallel magnitude
	Parallelism int
}

type Result struct {
	Input            Input
	Gzip             bool
	OperationResults []OperationResult
}

type OperationResult struct {
	// TODO: json serialize ellapsed
	TotalTime     *stats.Times `json:",omitempty"`
	GraphTime     *stats.Times `json:",omitempty"`
	SerializeTime *stats.Times `json:",omitempty"`
	Err           string       `json:",omitempty"`
}

func NewService(options OptionsGen, reqs SearchGraphQLRequester, logs Logger, candidate CandidateServer) *service {
	return &service{
		Options:                options,
		SearchGraphQLRequester: reqs,
		Logger:                 logs,
		Candidate:              candidate,
	}
}

type service struct {
	Options                OptionsGen
	SearchGraphQLRequester SearchGraphQLRequester
	Logger                 Logger
	Candidate              CandidateServer
}

func (s *service) HandlerFunc() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var acceptGzip bool
		acceptEncoding := strings.Split(r.Header.Get("Accept-Encoding"), ",")
		for _, item := range acceptEncoding {
			if strings.TrimSpace(item) == "gzip" {
				acceptGzip = true
				break
			}
		}

		var input Input
		json.NewDecoder(r.Body).Decode(&input)

		result, err := s.Serve(input, ":8081", "localhost", r.URL, acceptGzip)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(result)
	}
}

func (s *service) Serve(input Input, addr, host string, u *url.URL, acceptGzip bool) (Result, error) {
	operationResults := make(chan OperationResult, input.OperationNumber)
	// create candidate server
	options := s.Options.Gen(input.OptionNumber)
	candidateUrl := u.String() + "/candidate"
	s.Logger.Info(fmt.Sprintf("candidate up with addr: [%s] and url: [%s]", addr, candidateUrl))
	candidateServer, err := s.Candidate.NewServer(addr, candidateUrl, options, operationResults)
	if err != nil {
		return Result{}, err
	}

	// start listening the candidate server
	go func() {
		err := candidateServer.ListenAndServe()
		if err != nil {
			s.Logger.Info("candidate listen and serve error: " + err.Error())
		} else {
			s.Logger.Info("candidate server shutdown")
		}
	}()

	// send jobs
	jobs := make(chan struct{}, input.OperationNumber)
	for i := 0; i < input.OperationNumber; i++ {
		jobs <- struct{}{}
	}
	close(jobs)

	// send request for every job
	httpClient := http.DefaultClient
	graphQLRequest := s.SearchGraphQLRequester.SearchGraphQLRequest(input.ResolveScale)
	for i := 0; i < input.Parallelism; i++ {
		go func(i int) {
			s.Logger.Info(fmt.Sprintf("Worker [%v] running", i))
			for range jobs {
				s.Logger.Info(fmt.Sprintf("Worker [%v] got job", i))
				req, err := http.NewRequest("POST", "http://"+host+addr+candidateUrl, bytes.NewReader(graphQLRequest))
				if err != nil {
					s.Logger.Info(fmt.Sprintf("Worker [%v] job error creating request: %v", i, err))
					operationResults <- OperationResult{Err: err.Error()}
				}
				req.Header.Set("Content-Type", "application-json")
				if acceptGzip {
					req.Header.Set("Accept-Encoding", "gzip")
				}
				response, err := httpClient.Do(req)
				if err != nil {
					s.Logger.Info(fmt.Sprintf("Worker [%v] job error: %v", i, err))
					operationResults <- OperationResult{Err: err.Error()}
				} else {
					io.Copy(ioutil.Discard, response.Body)
					response.Body.Close()
				}
			}
		}(i)
	}

	result := Result{
		Input:            input,
		Gzip:             acceptGzip,
		OperationResults: make([]OperationResult, input.OperationNumber),
	}

	s.Logger.Info("waiting results")
	for i := 0; i < input.OperationNumber; i++ {
		result.OperationResults[i] = <-operationResults
	}
	close(operationResults)
	s.Logger.Info("finished waiting results")

	candidateServer.Shutdown(context.TODO())
	s.Logger.LogResult(result)
	return result, nil
}
