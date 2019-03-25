package stdjson

import (
	"encoding/json"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter"
	"net/http"
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
