package presenterbenchmark

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/travelgateX/presenters-benchmark/pkg/presenter"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter/easyjson"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter/ffjson"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter/gophers"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter/gqlgen"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter/jsoniter"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter/stdjson"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter/stdjsonmapping"
)

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

/*func BenchmarkCandidate_gqlgen_servicemodels_1_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgensm.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 1,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}*/

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
		Candidate:    stdjson.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 1,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_rest_servicemodels_1_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    stdjsonmapping.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 1,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_jsoniter_1_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    jsoniter.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 1,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_ffjson_servicemodels_1_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    ffjson.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 1,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_easyjson_1_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    easyjson.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 1,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

//func BenchmarkCandidate_gqlgen_2000_high(b *testing.B) {
//	benchmarkCandidates(b, candidateBenchmark{
//		Candidate:    gqlgen.Candidate{},
//		HTTPStatus:   http.StatusOK,
//		OptionNumber: 2000,
//		ResolveScale: presenter.ResolveScaleHigh,
//	})
//}

/*func BenchmarkCandidate_gqlgen_servicemodels_2000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgensm.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 2000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}*/

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
		Candidate:    stdjson.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 2000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}
func BenchmarkCandidate_rest_servicemodels_2000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    stdjsonmapping.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 2000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_jsoniter_2000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    jsoniter.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 2000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_ffjson_servicemodels_2000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    ffjson.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 2000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

//func BenchmarkCandidate_gqlgen_7000_high(b *testing.B) {
//	benchmarkCandidates(b, candidateBenchmark{
//		Candidate:    gqlgen.Candidate{},
//		HTTPStatus:   http.StatusOK,
//		OptionNumber: 7000,
//		ResolveScale: presenter.ResolveScaleHigh,
//	})
//}

/*func BenchmarkCandidate_gqlgen_servicemodels_7000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgensm.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 7000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}
*/

func BenchmarkCandidate_gophers_7000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gophers.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 7000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}
func BenchmarkCandidate_rest_21000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    stdjson.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 100000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}
func BenchmarkCandidate_rest_21000_high_para(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    stdjson.CandidateParallel{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 100000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_rest_servicemodels_7000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    stdjsonmapping.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 7000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_ffjson_servicemodels_7000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    ffjson.Candidate{},
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

/*func BenchmarkCandidate_gqlgen_servicemodels_20000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgensm.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 20000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}*/

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
		Candidate:    stdjson.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 20000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}
func BenchmarkCandidate_rest_servicemodels_20000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    stdjsonmapping.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 20000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_jsoniter_20000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    jsoniter.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 20000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_ffjson_servicemodels_20000_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    ffjson.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 20000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_easyjson_10000_high_p_nobuff(b *testing.B) {
	benchmarkCandidatesSer(b, candidateBenchmark{
		Candidate:    easyjson.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 10000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_easyjson_10000_high_p_Parallel(b *testing.B) {
	benchmarkCandidatesSer(b, candidateBenchmark{
		Candidate:    easyjson.CandidateParallel{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 10000,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func benchmarkCandidatesSer(b *testing.B, cb candidateBenchmark) {
	options := presenter.NewOptionsGen().Gen(cb.OptionNumber)
	hf, err := cb.Candidate.HandlerFunc(options)
	if err != nil {
		b.Fatalf("Error creating Handler: %v", err)
	}
	body := presenter.NewSearchGraphQLRequester().SearchGraphQLRequest(cb.ResolveScale)
	b.ResetTimer()

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
		if status := rr.Code; status != cb.HTTPStatus {
			b.Errorf("handler returned wrong status code: got %v want %v, body: %s", status, cb.HTTPStatus, rr.Body.String())
		}
	}
}

//func BenchmarkCandidate_gqlgen_65536_high(b *testing.B) {
//	benchmarkCandidates(b, candidateBenchmark{
//		Candidate:    gqlgen.Candidate{},
//		HTTPStatus:   http.StatusOK,
//		OptionNumber: 65536,
//		ResolveScale: presenter.ResolveScaleHigh,
//	})
//}

/*func BenchmarkCandidate_gqlgen_servicemodels_65536_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgensm.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 65536,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}
*/

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
		Candidate:    stdjson.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 65536,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}
func BenchmarkCandidate_rest_servicemodels_65536_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    stdjsonmapping.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 65536,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}
func BenchmarkCandidate_ffjson_servicemodels_65536_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    ffjson.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 65536,
		ResolveScale: presenter.ResolveScaleHigh,
	})
}

func BenchmarkCandidate_jsoniter_65536_high(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    jsoniter.Candidate{},
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

/*func BenchmarkCandidate_gqlgen_servicemodels_1_medium(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgensm.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 1,
		ResolveScale: presenter.ResolveScaleMedium,
	})
}
*/

func BenchmarkCandidate_gophers_1_medium(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gophers.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 1,
		ResolveScale: presenter.ResolveScaleMedium,
	})
}

//func BenchmarkCandidate_gqlgen_7000_medium(b *testing.B) {
//	benchmarkCandidates(b, candidateBenchmark{
//		Candidate:    gqlgen.Candidate{},
//		HTTPStatus:   http.StatusOK,
//		OptionNumber: 7000,
//		ResolveScale: presenter.ResolveScaleMedium,
//	})
//}

/*func BenchmarkCandidate_gqlgen_servicemodels_7000_medium(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgensm.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 7000,
		ResolveScale: presenter.ResolveScaleMedium,
	})
}*/

func BenchmarkCandidate_gophers_7000_medium(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gophers.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 7000,
		ResolveScale: presenter.ResolveScaleMedium,
	})
}

//func BenchmarkCandidate_gqlgen_20000_medium(b *testing.B) {
//	benchmarkCandidates(b, candidateBenchmark{
//		Candidate:    gqlgen.Candidate{},
//		HTTPStatus:   http.StatusOK,
//		OptionNumber: 20000,
//		ResolveScale: presenter.ResolveScaleMedium,
//	})
//}

/*func BenchmarkCandidate_gqlgen_servicemodels_20000_medium(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgensm.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 20000,
		ResolveScale: presenter.ResolveScaleMedium,
	})
}
*/

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

/*func BenchmarkCandidate_gqlgen_servicemodels_1_low(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgensm.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 1,
		ResolveScale: presenter.ResolveScaleLow,
	})
}*/

func BenchmarkCandidate_gophers_1_low(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gophers.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 1,
		ResolveScale: presenter.ResolveScaleLow,
	})
}

//func BenchmarkCandidate_gqlgen_7000_low(b *testing.B) {
//	benchmarkCandidates(b, candidateBenchmark{
//		Candidate:    gqlgen.Candidate{},
//		HTTPStatus:   http.StatusOK,
//		OptionNumber: 7000,
//		ResolveScale: presenter.ResolveScaleLow,
//	})
//}

/*func BenchmarkCandidate_gqlgen_servicemodels_7000_low(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gqlgensm.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 7000,
		ResolveScale: presenter.ResolveScaleLow,
	})
}
*/

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
		Candidate:    stdjsonmapping.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 7000,
		ResolveScale: presenter.ResolveScaleLow,
	})
}

func BenchmarkCandidate_ffjson_servicemodels_7000_low(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    ffjson.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 7000,
		ResolveScale: presenter.ResolveScaleLow,
	})
}

//func BenchmarkCandidate_gqlgen_20000_low(b *testing.B) {
//	benchmarkCandidates(b, candidateBenchmark{
//		Candidate:    gqlgen.Candidate{},
//		HTTPStatus:   http.StatusOK,
//		OptionNumber: 20000,
//		ResolveScale: presenter.ResolveScaleLow,
//	})
//}

//func BenchmarkCandidate_gqlgen_servicemodels_20000_low(b *testing.B) {
//	benchmarkCandidates(b, candidateBenchmark{
//		Candidate:    gqlgensm.Candidate{},
//		HTTPStatus:   http.StatusOK,
//		OptionNumber: 20000,
//		ResolveScale: presenter.ResolveScaleLow,
//	})
//}

func BenchmarkCandidate_gophers_20000_low(b *testing.B) {
	benchmarkCandidates(b, candidateBenchmark{
		Candidate:    gophers.Candidate{},
		HTTPStatus:   http.StatusOK,
		OptionNumber: 20000,
		ResolveScale: presenter.ResolveScaleLow,
	})
}
