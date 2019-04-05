package presenterbenchmark

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime/pprof"
	"sync"
	"testing"
	"time"

	"github.com/travelgateX/presenters-benchmark/pkg/presenter"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter/easyjson"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter/easyjsonmapping"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter/ffjson"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter/gophers"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter/gqlgen"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter/jsoniter"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter/protobuf"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter/simplejson"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter/stdjson"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter/stdjsonmapping"
)

var funcs = []struct {
	Name      string
	Candidate presenter.CandidateHandlerFunc
}{
	{"gophers", gophers.Candidate{}},
	{"gqlgen mapping", gqlgen.Candidate{}},
	//{"gqlgen service models", gqlgensm.Candidate{}}, // broken
	{"protobuf mapping", protobuf.Candidate{}},
	{"std json", stdjson.Candidate{}},
	{"std json mapping", stdjsonmapping.Candidate{}},
	{"ffjson mapping", ffjson.Candidate{}},
	{"simplejson", simplejson.Candidate{}},
	{"jsoniter", jsoniter.Candidate{}},
	{"easyjson", easyjson.Candidate{}},
	{"easyjson mapping", easyjsonmapping.Candidate{}},
}

var funcEasyJSON = []struct {
	Name      string
	Candidate presenter.CandidateHandlerFunc
}{
	{"easyjson", easyjson.Candidate{}},
	{"easyjson parallel", easyjson.CandidateParallel{}},
	{"easyjson parallel (int buf)", easyjson.CandidateParallelGzip{}},
	{"easyjson parallel (int buf channel)", easyjson.CandidateParallelChannel{}},
	{"easyjson parallel (int bigger buf channel)", easyjson.CandidateParallelChannelBiggerBuf{}},
}

var funcEasyJSONChannel = []struct {
	Name      string
	Candidate presenter.CandidateChannelHandlerFunc
}{
	{"easyjson sequential chan wait", easyjson.Candidate{}},
	{"easyjson sequential chan per option", easyjson.CandidateParallel{}},
}

// Variables:
// - Options
// - ResolveScale
// - Candidate Implementation
func BenchmarkSequential(b *testing.B) {
	body := presenter.NewSearchGraphQLRequester().SearchGraphQLRequest(presenter.ResolveScaleHigh)
	optionGen := presenter.NewOptionsGen()
	for _, f := range funcEasyJSON {
		time.Sleep(2 * time.Second)
		for optNumber := 1; optNumber <= 65536; optNumber *= 2 {
			hf, err := f.Candidate.HandlerFunc(optionGen.Gen(optNumber))
			if err != nil {
				b.Fatalf("Error creating Handler: %v", err)
			}
			b.Run(fmt.Sprintf("%s/%d", f.Name, optNumber), func(b *testing.B) {
				if false {
					memProfile, err := os.Create("example-mem.prof")
					if err != nil {
						b.Fatal(err)
					}
					defer func() {
						if err := pprof.WriteHeapProfile(memProfile); err != nil {
							b.Fatal(err)
						}
						memProfile.Close()
					}()
				}
				for i := 0; i < b.N; i++ {
					req, err := http.NewRequest("POST", "/status", bytes.NewReader(body))
					if err != nil {
						b.Fatal(err)
					}

					// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
					rr := httptest.NewRecorder()
					handler := http.HandlerFunc(hf)

					// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
					// directly and pass in our Request and ResponseRecorder.
					handler.ServeHTTP(rr, req)

					// Check the status code is what we expect.
					if status := rr.Code; status != http.StatusOK {
						b.Errorf("handler returned wrong status code: got %v want %v, body: %s", status, http.StatusOK, rr.Body.String())
					}
				}
			})
		}
	}
}

func BenchmarkParallel(b *testing.B) {
	body := presenter.NewSearchGraphQLRequester().SearchGraphQLRequest(presenter.ResolveScaleHigh)
	optionGen := presenter.NewOptionsGen()
	for _, f := range funcEasyJSON {
		time.Sleep(2 * time.Second)
		for optNumber := 4096; optNumber <= 65536; optNumber *= 2 {
			hf, err := f.Candidate.HandlerFunc(optionGen.Gen(optNumber))
			if err != nil {
				b.Fatalf("Error creating Handler: %v", err)
			}
			b.Run(fmt.Sprintf("%s/%d", f.Name, optNumber), func(b *testing.B) {
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						req, err := http.NewRequest("POST", "/status", bytes.NewReader(body))
						if err != nil {
							b.Fatal(err)
						}

						// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
						rr := httptest.NewRecorder()
						handler := http.HandlerFunc(hf)

						// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
						// directly and pass in our Request and ResponseRecorder.
						handler.ServeHTTP(rr, req)

						// Check the status code is what we expect.
						if status := rr.Code; status != http.StatusOK {
							b.Errorf("handler returned wrong status code: got %v want %v, body: %s", status, http.StatusOK, rr.Body.String())
						}
					}
				})
			})
		}
	}
}

func BenchmarkSequentialGzip(b *testing.B) {
	body := presenter.NewSearchGraphQLRequester().SearchGraphQLRequest(presenter.ResolveScaleHigh)
	optionGen := presenter.NewOptionsGen()
	for _, f := range funcEasyJSON {
		time.Sleep(2 * time.Second)
		for optNumber := 4096; optNumber <= 65536; optNumber *= 2 {
			hf, err := f.Candidate.HandlerFunc(optionGen.Gen(optNumber))
			if err != nil {
				b.Fatalf("Error creating Handler: %v", err)
			}
			if false {
				memProfile, err := os.Create("example-mem-gzip_seq.prof")
				if err != nil {
					b.Fatal(err)
				}
				defer func() {
					if err := pprof.WriteHeapProfile(memProfile); err != nil {
						b.Fatal(err)
					}
					memProfile.Close()
				}()
			}
			b.Run(fmt.Sprintf("%s/%d", f.Name, optNumber), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					req, err := http.NewRequest("POST", "/status", bytes.NewReader(body))
					if err != nil {
						b.Fatal(err)
					}

					// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
					rr := httptest.NewRecorder()
					handler := http.HandlerFunc(hf)

					// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
					// directly and pass in our Request and ResponseRecorder.
					w, close := presenter.GzipCompress(rr)
					handler.ServeHTTP(w, req)
					close()

					// Check the status code is what we expect.
					if status := rr.Code; status != http.StatusOK {
						b.Errorf("handler returned wrong status code: got %v want %v, body: %s", status, http.StatusOK, rr.Body.String())
					}
				}
			})
		}
	}
}

func BenchmarkParallelGzip(b *testing.B) {
	body := presenter.NewSearchGraphQLRequester().SearchGraphQLRequest(presenter.ResolveScaleHigh)
	optionGen := presenter.NewOptionsGen()
	for _, f := range funcEasyJSON {
		time.Sleep(10 * time.Second)
		for optNumber := 4096; optNumber <= 65536; optNumber *= 2 {
			hf, err := f.Candidate.HandlerFunc(optionGen.Gen(optNumber))
			if err != nil {
				b.Fatalf("Error creating Handler: %v", err)
			}
			b.Run(fmt.Sprintf("%s/%d", f.Name, optNumber), func(b *testing.B) {
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						req, err := http.NewRequest("POST", "/status", bytes.NewReader(body))
						if err != nil {
							b.Fatal(err)
						}

						// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
						rr := httptest.NewRecorder()
						handler := http.HandlerFunc(hf)

						// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
						// directly and pass in our Request and ResponseRecorder.
						w, close := presenter.GzipCompress(rr)
						handler.ServeHTTP(w, req)
						close()

						// Check the status code is what we expect.
						if status := rr.Code; status != http.StatusOK {
							b.Errorf("handler returned wrong status code: got %v want %v, body: %s", status, http.StatusOK, rr.Body.String())
						}
					}
				})
			})
		}
	}
}

func BenchmarkSequentialChan(b *testing.B) {
	body := presenter.NewSearchGraphQLRequester().SearchGraphQLRequest(presenter.ResolveScaleHigh)
	optionGen := presenter.NewOptionsGen()
	for _, f := range funcEasyJSONChannel {
		time.Sleep(2 * time.Second)
		for optNumber := 64; optNumber <= 65536; optNumber *= 2 {
			opts := optionGen.Gen(optNumber)

			b.Run(fmt.Sprintf("%s/%d", f.Name, optNumber), func(b *testing.B) {
				if false {
					memProfile, err := os.Create("example-mem.prof")
					if err != nil {
						b.Fatal(err)
					}
					defer func() {
						if err := pprof.WriteHeapProfile(memProfile); err != nil {
							b.Fatal(err)
						}
						memProfile.Close()
					}()
				}
				for i := 0; i < b.N; i++ {
					optsC := make(chan *presenter.Option, optNumber)
					for _, opt := range opts {
						optsC <- opt
					}
					close(optsC)
					hf, err := f.Candidate.ChannelHandlerFunc(optsC)
					if err != nil {
						b.Fatalf("Error creating Handler: %v", err)
					}
					req, err := http.NewRequest("POST", "/status", bytes.NewReader(body))
					if err != nil {
						b.Fatal(err)
					}

					// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
					rr := httptest.NewRecorder()
					handler := http.HandlerFunc(hf)

					// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
					// directly and pass in our Request and ResponseRecorder.
					handler.ServeHTTP(rr, req)

					// Check the status code is what we expect.
					if status := rr.Code; status != http.StatusOK {
						b.Errorf("handler returned wrong status code: got %v want %v, body: %s", status, http.StatusOK, rr.Body.String())
					}
				}
			})
		}
	}
}

var final interface{}

func BenchmarkIterateOptions(b *testing.B) {
	optionGen := presenter.NewOptionsGen()
	for optNumber := 1; optNumber <= 65536; optNumber *= 2 {
		opts := optionGen.Gen(optNumber)
		b.Run(fmt.Sprintf("%s/%d", "iterate options", optNumber), func(b *testing.B) {
			var ret *presenter.Option
			for i := 0; i < b.N; i++ {
				for _, opt := range opts {
					ret = opt
				}
			}
			final = ret
		})
	}
}

func BenchmarkOptionChannel(b *testing.B) {
	optionGen := presenter.NewOptionsGen()
	for optNumber := 1; optNumber <= 65536; optNumber *= 2 {
		opts := optionGen.Gen(optNumber)
		b.Run(fmt.Sprintf("%s/%d", "iterate options", optNumber), func(b *testing.B) {
			ch := make(chan *presenter.Option, optNumber)
			wg := sync.WaitGroup{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				for range ch {
				}
			}()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, opt := range opts {
					ch <- opt
				}
			}
			close(ch)
			wg.Wait()
		})
	}
}
