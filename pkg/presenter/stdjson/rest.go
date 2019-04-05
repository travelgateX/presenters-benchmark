package stdjson

import (
	"encoding/json"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter"
	"net/http"
	"sync"
)

type Candidate struct{}

var _ presenter.CandidateHandlerFunc = (*Candidate)(nil)

func (Candidate) HandlerFunc(options []*presenter.Option) (http.HandlerFunc, error) {
	return func(w http.ResponseWriter, r *http.Request) {
		var params struct {
			Query         string                 `json:"query"`
			OperationName string                 `json:"operationName"`
			Variables     map[string]interface{} `json:"variables"`
		}

		// mandatory to check this in all example
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// deserialize and return option
		response := presenter.Response{}
		response.Data.HotelX.Search.Options = options
		b, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(b)
	}, nil
}

func (Candidate) UnmarshalOptions(b []byte) ([]*presenter.Option, error) {
	return presenter.JSONUnmarshalOptions(b)
}

type CandidateParallel struct{}

var _ presenter.CandidateHandlerFunc = (*CandidateParallel)(nil)

func (CandidateParallel) HandlerFunc(options []*presenter.Option) (http.HandlerFunc, error) {
	return func(w http.ResponseWriter, r *http.Request) {
		var params struct {
			Query         string                 `json:"query"`
			OperationName string                 `json:"operationName"`
			Variables     map[string]interface{} `json:"variables"`
		}
		// mandatory to check this in all example
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		buffer := &Buffer{mu: sync.Mutex{}, w: w}
		buffer.Write([]byte(`{"data": {"hotelX": {"search": {"options": [`))

		parallelism := 8
		wg := &sync.WaitGroup{}
		wg.Add(parallelism)
		optionsC := make(chan *presenter.Option, parallelism)
		for i := 0; i < parallelism; i++ {
			go func() {
				for opt := range optionsC {
					b, err := json.Marshal(opt)
					if err != nil {
						panic(err)
					}
					b = append(b, ',')
					buffer.Write(b)
				}
				wg.Done()
			}()
		}

		for _, opt := range options[:len(options)-1] {
			optionsC <- opt
		}
		close(optionsC)
		wg.Wait()
		opt := options[len(options)-1]

		b, err := json.Marshal(opt)
		if err != nil {
			panic(err)
		}

		buffer.Write(b)
		_, err = buffer.Write([]byte(`],"errors": {"code": "","type": "","description": ""}}}}}`))
		if err != nil {
			panic(err)
		}
		return
	}, nil
}

func (CandidateParallel) UnmarshalOptions(b []byte) ([]*presenter.Option, error) {
	return presenter.JSONUnmarshalOptions(b)
}

type Buffer struct {
	mu sync.Mutex
	w  http.ResponseWriter
}

func (b *Buffer) Write(p []byte) (int, error) {
	b.mu.Lock()
	i, err := b.w.Write(p)
	b.mu.Unlock()
	return i, err
}
