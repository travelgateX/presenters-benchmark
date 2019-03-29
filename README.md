# Presenters 

This is a repository to compare different types of presenters in means of performance.
We are measuring the serialization of Options to an HTTP Response Body, where Options are always the same mock


## Implementations


| tech      | lib 		 | From service models* | Mapping* |
|:----------|:-----------|:---------------------:|:---------:|
| Graphql   | graphql-go | 		    x  		    |  	 ✓	   |
| 			| gqlgen     |          ✓		    |  	 ✓	   |
| JSON		| standard   |          ✓ 		    |  	 ✓	   |
| 		    | easyjson   |          ✓		    |  	 ✓	   |
| 		    | ffjson     |          ✓		    |  	 x	   |
| 		    | jsoniter   |          ✓		    | 	 x 	   |
| Protobuf  | standard   |        	x	   	    |  	 ✓	   |

*From service models: serialize directly from service models
*Mapping: from application models to presentation models then serialize

To add another implementation create its own package under 'pkg/presenter' and implement the interface:

```go
type CandidateHandlerFunc interface {
	HandlerFunc(options []*Option) (http.HandlerFunc, error)
}
```

## Run the tests

A valid candidate must pass tests in [candidate.go](pkg/presenter/candidate.go).

Usage example of the [gophers](pkg/presenter/gophers/gophers_test.go) implementation: 

```go
func TestCandidate(t *testing.T) {
	presenter.TestCandidateHandleFunc(t, Candidate{})
}
```

## Run the benchmarks

To run the benchmarks, run 'BenchmarkSequential'  or 'BenchmarkParallel' in [benchmark_test.go](benchmark/benchmark_test.go)

Use [this excel](benchmark/graphics.xlsx) to build graphics

### Sample 
PC: Intel(R) Core(TM) i7-3520M CPU @ 2.90GHz x 4 - 8,00 GB - Windows 10 Enterprise 1809

Request: [high resolve scale](benchmark/resolveScale_high.txt)

Response: N times [option](benchmark/option.json)


- Sequential benchmarks

| implementation      | options	   | time (ms/op) | mem (B/op) | allocs (allocs/op) |
|:--------------------|:-----------|---------------------:|---------:|---------:|
| graphql-go   		  | 65536	   | 13 308 997 |  	 2 981 868 976	   |  	45311713	   |
| gqlgen_mapping	  | 65536	   | 15 133 237 |  	 2 705 556 808	   |  	 43865858	   |
| std_json			  | 65536	   | 1 180 024 |  	 436 686 464	   |  	 425	   |
| std_json_mapping	  | 65536	   | 1 308 006 |  	 524 577 840	   |  	 524714	   |
| ffjson_mapping	  | 65536	   | 3 427 002 | 	 902 774 360 	   | 	 4063767 	   |
| jsoniter			  | 65536	   | 1 515 506 |  	 514 238 520	   |  	 5636635	   |
| easyjson			  | 65536	   | 807 003 |  	 343 538 248	   |  	 6829	   |
| easyjson_mapping	  | 65536	   | 1 021 005 |  	 435 678 392	   |  	 531883	   |
| protobuf_mapping	  | 65536	   | 678 511 |  	 267 425 264	   |  	1376657	   |

Full results can be found in ![benchmark_results.txt](benchmark/benchmark_results.txt)

![Time sequential](benchmark/time_seq.jpg?raw=true "Title")


## Run the profiler

```
cd presenters/benchmark/
go test -benchmem -cpuprofile restsm65536.pprof -test.benchtime 10s -run=^$ rfc/presenters/benchmark -bench BenchmarkCandidate_rest_servicemodels_65536_high
mv benchmark.test restsm65536.test
go tool pprof -http=:6060 restsm65536.test restsm65536.pprof
```

- Graphql-go

![Gophers](benchmark/gophers65536.png?raw=true "Title")