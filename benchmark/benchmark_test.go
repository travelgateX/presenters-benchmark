package presenterbenchmark

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"presenters-benchmark/pkg/presenter"
	"presenters-benchmark/pkg/presenter/gophers"
	"presenters-benchmark/pkg/presenter/gqlgen"
	"presenters-benchmark/pkg/presenter/gqlgensm"
	"presenters-benchmark/pkg/presenter/protobuf"
	"presenters-benchmark/pkg/presenter/rest"
	"presenters-benchmark/pkg/presenter/resteasyjson"
	"presenters-benchmark/pkg/presenter/restmapping"
	"testing"
	"time"
)

// Variables:
// - Options
// - ResolveScale
// - Candidate Implementation
func BenchmarkSequential(b *testing.B) {
	body := presenter.NewSearchGraphQLRequester().SearchGraphQLRequest(presenter.ResolveScaleHigh)
	optionGen := presenter.NewOptionsGen()
	for _, f := range funcs {
		time.Sleep(2 * time.Second)
		for optNumber := 1; optNumber <= 65536; optNumber *= 2 {
			hf, err := f.Candidate.HandlerFunc(optionGen.Gen(optNumber))
			if err != nil {
				b.Fatalf("Error creating Handler: %v", err)
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
	for _, f := range funcs {
		time.Sleep(2 * time.Second)
		for optNumber := 1; optNumber <= 65536; optNumber *= 2 {
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

var funcs = []struct {
	Name      string
	Candidate presenter.CandidateHandlerFunc
}{
	{"gophers", gophers.Candidate{}},
	{"gqlgen mapping", gqlgen.Candidate{}},
	{"gqlgen service models", gqlgensm.Candidate{}},
	{"rest json service models", rest.Candidate{}},
	{"rest json mapping", restmapping.Candidate{}},
	{"protobuf mapping", protobuf.Candidate{}},
	{"rest easyJson mapping", resteasyjson.Candidate{}},
}

func benchmarkCandidates(b *testing.B, cb candidateBenchmark) {
	options := presenter.NewOptionsGen().Gen(cb.OptionNumber)
	hf, err := cb.Candidate.HandlerFunc(options)
	if err != nil {
		b.Fatalf("Error creating Handler: %v", err)
	}
	body := presenter.NewSearchGraphQLRequester().SearchGraphQLRequest(cb.ResolveScale)
	b.ResetTimer()

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
			if status := rr.Code; status != cb.HTTPStatus {
				b.Errorf("handler returned wrong status code: got %v want %v, body: %s", status, cb.HTTPStatus, rr.Body.String())
			}
		}
	})
}

type candidateBenchmark struct {
	Candidate  presenter.CandidateHandlerFunc
	HTTPStatus int
	// OptionNumber are the number of options that each operation is going to return
	OptionNumber int
	// Scale of fields in the graph Request
	// low: options with 1 field
	// medium: options with half the fields
	// high: options with all the fields
	ResolveScale presenter.ResolveScale
}

// HIGH

func BenchmarkCandidate_gqlgen_1_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgen.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 1,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_gqlgen_servicemodels_1_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgensm.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 1,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_gophers_1_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gophers.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 1,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_rest_1_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    rest.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 1,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_rest_servicemodels_1_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    restmapping.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 1,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_gqlgen_2000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgen.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 2000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_gqlgen_servicemodels_2000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgensm.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 2000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_gophers_2000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gophers.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 2000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_rest_2000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    rest.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 2000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}
func BenchmarkCandidate_rest_servicemodels_2000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    restmapping.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 2000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_gqlgen_7000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgen.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 7000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_gqlgen_servicemodels_7000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgensm.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 7000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_gophers_7000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gophers.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 7000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}
func BenchmarkCandidate_rest_7000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    rest.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 7000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}
func BenchmarkCandidate_rest_servicemodels_7000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    restmapping.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 7000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}
func BenchmarkCandidate_gqlgen_20000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgen.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 20000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_gqlgen_servicemodels_20000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgensm.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 20000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_gophers_20000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gophers.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 20000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_rest_20000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    rest.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 20000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}
func BenchmarkCandidate_rest_servicemodels_20000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    restmapping.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 20000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_gqlgen_65536_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgen.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 65536,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_gqlgen_servicemodels_65536_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgensm.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 65536,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_gophers_65536_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gophers.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 65536,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_rest_65536_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    rest.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 65536,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}
func BenchmarkCandidate_rest_servicemodels_65536_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    restmapping.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 65536,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

// Medium

func BenchmarkCandidate_gqlgen_1_medium(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgen.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 1,
		ResolveScale: presenter.ResolveScaleMedium,
	})
}

func BenchmarkCandidate_gqlgen_servicemodels_1_medium(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgensm.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 1,
		ResolveScale: presenter.ResolveScaleMedium,
	})
}

func BenchmarkCandidate_gophers_1_medium(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gophers.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 1,
		ResolveScale: presenter.ResolveScaleMedium,
	})
}

func BenchmarkCandidate_gqlgen_7000_medium(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgen.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 7000,
		ResolveScale: presenter.ResolveScaleMedium,
	})
}

func BenchmarkCandidate_gqlgen_servicemodels_7000_medium(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgensm.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 7000,
		ResolveScale: presenter.ResolveScaleMedium,
	})
}

func BenchmarkCandidate_gophers_7000_medium(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gophers.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 7000,
		ResolveScale: presenter.ResolveScaleMedium,
	})
}

func BenchmarkCandidate_gqlgen_20000_medium(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgen.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 20000,
		ResolveScale: presenter.ResolveScaleMedium,
	})
}

func BenchmarkCandidate_gqlgen_servicemodels_20000_medium(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgensm.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 20000,
		ResolveScale: presenter.ResolveScaleMedium,
	})
}

func BenchmarkCandidate_gophers_20000_medium(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gophers.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 20000,
		ResolveScale: presenter.ResolveScaleMedium,
	})
}

// Low

func BenchmarkCandidate_gqlgen_1_low(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgen.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 1,
		ResolveScale: presenter.ResolveScaleLow,
	})
}

func BenchmarkCandidate_gqlgen_servicemodels_1_low(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgensm.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 1,
		ResolveScale: presenter.ResolveScaleLow,
	})
}

func BenchmarkCandidate_gophers_1_low(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gophers.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 1,
		ResolveScale: presenter.ResolveScaleLow,
	})
}

func BenchmarkCandidate_gqlgen_7000_low(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgen.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 7000,
		ResolveScale: presenter.ResolveScaleLow,
	})
}

func BenchmarkCandidate_gqlgen_servicemodels_7000_low(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgensm.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 7000,
		ResolveScale: presenter.ResolveScaleLow,
	})
}

func BenchmarkCandidate_gophers_7000_low(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gophers.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 7000,
		ResolveScale: presenter.ResolveScaleLow,
	})
}

func BenchmarkCandidate_rest_servicemodels_7000_low(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    restmapping.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 7000,
		ResolveScale: presenter.ResolveScaleLow,
	})
}

func BenchmarkCandidate_gqlgen_20000_low(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgen.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 20000,
		ResolveScale: presenter.ResolveScaleLow,
	})
}

func BenchmarkCandidate_gqlgen_servicemodels_20000_low(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgensm.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 20000,
		ResolveScale: presenter.ResolveScaleLow,
	})
}

func BenchmarkCandidate_gophers_20000_low(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gophers.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 20000,
		ResolveScale: presenter.ResolveScaleLow,
	})
}
