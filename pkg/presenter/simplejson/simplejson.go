package simplejson

import (
	"encoding/json"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter"
	"net/http"

	simplejson "github.com/likexian/simplejson-go"
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

		//part serialize
		jsonObject := simplejson.New(options)

		w.Write(prefix)
		jsonObjectSerialized, err := jsonObject.Dumps()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write([]byte(jsonObjectSerialized))
		w.Write(sufix)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}, nil
}

func (Candidate) UnmarshalOptions(b []byte) ([]*presenter.Option, error) {
	return presenter.JSONUnmarshalOptions(b)
}

var prefix = []byte(`{"data": {"hotelX": {"search": {"options": `)
var sufix = []byte(`,"errors": {"code": "","type": "","description": ""}}}}}`)
