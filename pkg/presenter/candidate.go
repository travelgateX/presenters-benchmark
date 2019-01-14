package presenter

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// CandidateHandlerFunc is the interface used to benchmark
type CandidateHandlerFunc interface {
	// HandlerFunc returned must write the given options in the response body
	HandlerFunc(options []*Option) (http.HandlerFunc, error)
	// UnmarshalOptions must take the bytes of a responseBody and unmarshal them to Options
	UnmarshalOptions([]byte) ([]*Option, error)
}

// JSONUnmarshalOptions is a valid UnmarshalOptions func for those implementations that write
// a JSON encoded 'Response' in the response body
func JSONUnmarshalOptions(b []byte) ([]*Option, error) {
	var res Response
	err := json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}
	return res.Data.HotelX.Search.Options, nil
}

// TestCandidateHandleFunc tests that a candidate encodes options correctly
func TestCandidateHandleFunc(t *testing.T, c CandidateHandlerFunc) {
	opts := NewOptionsGen().Gen(100)
	hf, err := c.HandlerFunc(opts)
	if err != nil {
		t.Fatalf("error creating candidate's handler func %v", err)
	}

	body := NewSearchGraphQLRequester().SearchGraphQLRequest(ResolveScaleHigh)
	req, err := http.NewRequest("POST", "/status", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("error in NewRequest %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hf)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v, body: %s", status, http.StatusOK, rr.Body.String())
	}

	// deserialize option, it must be the same as the generated
	var retOpts []*Option
	retOpts, err = c.UnmarshalOptions(rr.Body.Bytes())
	if err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}

	if len(opts) != len(retOpts) {
		t.Fatalf("different len of options generated vs returned len: %v, %v", len(opts), len(retOpts))
	}
	for i := range opts {
		if !OptionEquals(opts[i], retOpts[i]) {
			b1, err := json.Marshal(*retOpts[i])
			if err != nil {
				t.Errorf("error marshalling opt: %v", err)
			}
			b2, err := json.Marshal(retOpts[i])
			if err != nil {
				t.Errorf("error marshalling opt: %v", err)
			}
			t.Fatalf("options are not serialized correctly: opt1: %s\nopt2: %s ", string(b1), string(b2))
		}
	}
}
