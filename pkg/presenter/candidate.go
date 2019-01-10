package presenter

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"hub-aggregator/common/kit/routing"
	"hub-aggregator/common/stats"
)

type CandidateHandlerFunc interface {
	HandlerFunc(options []*Option) (http.HandlerFunc, error)
}

type CandidateServer interface {
	NewServer(addr, pattern string, options []*Option, results chan<- OperationResult) (*routing.Server, error)
}

func NewGzipCandidateServer(addr, pattern string, handlerFunc http.HandlerFunc, mws ...routing.Middleware) *routing.Server {
	mws = append(mws, routing.GzipCompress())
	return routing.NewServer(
		addr,
		[]routing.Route{
			routing.Route{
				Name:              "Candidate",
				Method:            "POST",
				Pattern:           pattern,
				Middlewares:       mws,
				ServiceHandleFunc: handlerFunc,
			},
		},
	)
}

func TotalTimeMiddleware(results chan<- OperationResult) routing.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time := OperationResult{}
			startTime := stats.UtcNow()
			next.ServeHTTP(w, r)
			time.TotalTime = stats.NewTimes(startTime)
			results <- time
		})
	}
}

type Response struct {
	Data struct {
		HotelX struct {
			Search struct {
				Options []Option `json:"options"`
				Errors  struct {
					Code        string `json:"code"`
					Type        string `json:"type"`
					Description string `json:"description"`
				} `json:"errors"`
			} `json:"search"`
		} `json:"hotelX"`
	} `json:"data"`
}

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
	var res Response
	json.Unmarshal(rr.Body.Bytes(), &res)

	if len(opts) != len(res.Data.HotelX.Search.Options) {
		t.Fatalf("different len of options generated vs returned len: %v, %v", len(opts), len(res.Data.HotelX.Search.Options))
	}
	for i := range opts {
		if !optEquals(*opts[i], res.Data.HotelX.Search.Options[i]) {
			b1, err := json.Marshal(*opts[i])
			if err != nil {
				t.Errorf("error marshalling opt: %v", err)
			}
			b2, err := json.Marshal(res.Data.HotelX.Search.Options[i])
			if err != nil {
				t.Errorf("error marshalling opt: %v", err)
			}
			t.Fatalf("options are not serialized correctly: opt1: %s\nopt2: %s ", string(b1), string(b2))
		}
	}
}

func optEquals(opt1, opt2 Option) bool {
	return reflect.DeepEqual(opt1, opt2)
}
