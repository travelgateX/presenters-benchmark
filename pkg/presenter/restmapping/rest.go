package restmapping

import (
	"encoding/json"
	"hub-aggregator/common/kit/routing"
	"hub-aggregator/common/stats"
	"net/http"
	"rfc/presenters/pkg/presenter"
)

type Candidate struct{}

var _ presenter.CandidateServer = (*Candidate)(nil)
var _ presenter.CandidateHandlerFunc = (*Candidate)(nil)

func (Candidate) NewServer(addr, pattern string, options []*presenter.Option, results chan<- presenter.OperationResult) (*routing.Server, error) {
	return presenter.NewGzipCandidateServer(
		addr,
		pattern,
		HandlerFunc(options, results),
	), nil
}

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
		err := json.NewEncoder(w).Encode(NewResponse(options))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}, nil
}

func HandlerFunc(options []*presenter.Option, results chan<- presenter.OperationResult) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(NewResponse(options))
		time := presenter.OperationResult{}
		startTime := stats.UtcNow()
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
		deserializeTime := stats.UtcNow()
		// deserialize and return option
		err := json.NewEncoder(w).Encode(options)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		time.SerializeTime = stats.NewTimes(deserializeTime)
		time.TotalTime = stats.NewTimes(startTime)
		results <- time
	}
}
