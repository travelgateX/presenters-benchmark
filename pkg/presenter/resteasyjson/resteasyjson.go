package resteasyjson

import (
	"encoding/json"
	"github.com/mailru/easyjson/jwriter"
	"net/http"
	"presenters-benchmark/pkg/presenter"
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
		// TODO aplicar la libreria EasyJson
		jw := jwriter.Writer{NoEscapeHTML:true}
		response := NewResponse(options)
		response.MarshalEasyJSON(&jw)
		if jw.Error != nil {
			http.Error(w, jw.Error.Error(), http.StatusInternalServerError)
		}

		_, err := jw.DumpTo(w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}, nil
}

func (Candidate) UnmarshalOptions(b []byte) ([]*presenter.Option, error) {
	return presenter.JSONUnmarshalOptions(b)
}
