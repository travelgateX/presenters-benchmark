package rest

import (
	"encoding/json"
	"hub-aggregator/common/kit/routing"
	"hub-aggregator/common/stats"
	"net/http"
	"rfc/presenters/pkg/presenter"
)

type Candidate struct{}

var _ presenter.CandidateHandlerFunc = (*Candidate)(nil)
var _ presenter.CandidateServer = (*Candidate)(nil)

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
		w.Write(prefix)
		err := json.NewEncoder(w).Encode(options)
		w.Write(sufix)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}, nil
}

func HandlerFunc(options []*presenter.Option, results chan<- presenter.OperationResult) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		w.Write(prefix)
		err := json.NewEncoder(w).Encode(options)
		w.Write(sufix)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		time.SerializeTime = stats.NewTimes(deserializeTime)
		time.TotalTime = stats.NewTimes(startTime)
		results <- time
	}
}

var prefix = []byte(`{"data": {"hotelX": {"search": {"options": `)
var sufix = []byte(`,"errors": {"code": "","type": "","description": ""}}}}}`)
