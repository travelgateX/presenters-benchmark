package rest_simplejson

import (
	"io/ioutil"
	"net/http"
	"presenters-benchmark/pkg/presenter"

	simplejson "github.com/likexian/simplejson-go"
)

type Candidate struct{}

var _ presenter.CandidateHandlerFunc = (*Candidate)(nil)

func (Candidate) HandlerFunc(options []*presenter.Option) (http.HandlerFunc, error) {
	return func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, err := ioutil.ReadAll(r.Body)
		bodyString := string(bodyBytes)

		//part: deserialize
		if _, err := simplejson.Loads(bodyString); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//part serialize
		jsonObject := simplejson.New(options)
		if nil == jsonObject {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write(prefix)
		jsonObjectSerialized, err := jsonObject.Dumps()
		if nil != err {
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
