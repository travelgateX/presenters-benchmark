package protobuf

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"net/http"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter"
)

type Candidate struct{}

var _ presenter.CandidateHandlerFunc = (*Candidate)(nil)

func (c Candidate) HandlerFunc(options []*presenter.Option) (http.HandlerFunc, error) {
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

		b, err := c.MarshalSearchReply(NewResponse(options))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(b)
	}, nil
}

func (Candidate) MarshalSearchReply(sr *SearchReply) ([]byte, error) {
	return proto.Marshal(sr)
}

func (Candidate) UnmarshalOptions(b []byte) ([]*presenter.Option, error) {
	var sr SearchReply
	err := proto.Unmarshal(b, &sr)
	if err != nil {
		return nil, err
	}

	// json.NewEncoder(os.Stdout).Encode(sr)
	// TODO: reverse parsing?....
	return nil, nil
}
