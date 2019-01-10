# Presenters

This is a repository to compare different types of presenters in means of performance.
We are measuring the serialization of Options to an HTTP Response Body, where Options are always the same mock. All implementations must return the same response.


## Implementations

	- Graphql
		- gqlgen
			- mapping: from application models to presentation models then serialize
			- from service models: resolve directly from service models
		- gophers
			- mapping (the only way)
	- REST
		- JSON
			- mapping
			- from service models


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
PC: Ubuntu 16.04, 7.7 GiB mem, Intel Core i5-3230M CPU @ 2.60GHz × 4

Request: [high resolve scale](benchmark/resolveScale_high.txt)

Response: N times [option](benchmark/option.json)


- Sequential benchmarks

```
BenchmarkSequential/gophers/1-4         	    1000	   1451941 ns/op	  352200 B/op	    1912 allocs/op
BenchmarkSequential/gophers/2-4         	    1000	   1563888 ns/op	  387502 B/op	    2596 allocs/op
BenchmarkSequential/gophers/4-4         	    1000	   2012883 ns/op	  459856 B/op	    3962 allocs/op
BenchmarkSequential/gophers/8-4         	     500	   2831767 ns/op	  606062 B/op	    6693 allocs/op
BenchmarkSequential/gophers/16-4        	     300	   4160711 ns/op	  899982 B/op	   12153 allocs/op
BenchmarkSequential/gophers/32-4        	     200	   7452941 ns/op	 1497587 B/op	   23074 allocs/op
BenchmarkSequential/gophers/64-4        	     100	  14099009 ns/op	 2739470 B/op	   44914 allocs/op
BenchmarkSequential/gophers/128-4       	     100	  27575092 ns/op	 5112244 B/op	   88592 allocs/op
BenchmarkSequential/gophers/256-4       	      30	  54522222 ns/op	 9875169 B/op	  175948 allocs/op
BenchmarkSequential/gophers/512-4       	      20	 107578761 ns/op	19386161 B/op	  350656 allocs/op
BenchmarkSequential/gophers/1024-4      	      10	 196839044 ns/op	38434494 B/op	  700051 allocs/op
BenchmarkSequential/gophers/2048-4      	       3	 351876251 ns/op	76515384 B/op	 1398826 allocs/op
BenchmarkSequential/gophers/4096-4      	       2	 656139423 ns/op	152679768 B/op	 2796398 allocs/op
BenchmarkSequential/gophers/8192-4      	       1	1316191332 ns/op	304981128 B/op	 5591453 allocs/op
BenchmarkSequential/gophers/16384-4     	       1	2485446165 ns/op	609525256 B/op	11181546 allocs/op
BenchmarkSequential/gophers/32768-4     	       1	4897353396 ns/op	1218636856 B/op	22361827 allocs/op
BenchmarkSequential/gophers/65536-4     	       1	9826655318 ns/op	2436812200 B/op	44722211 allocs/op
BenchmarkSequential/gqlgen_mapping/1-4  	    2000	    739517 ns/op	   84775 B/op	    1288 allocs/op
BenchmarkSequential/gqlgen_mapping/2-4  	    2000	    852009 ns/op	  129343 B/op	    2111 allocs/op
BenchmarkSequential/gqlgen_mapping/4-4  	    2000	   1036159 ns/op	  218237 B/op	    3755 allocs/op
BenchmarkSequential/gqlgen_mapping/8-4  	    1000	   1602967 ns/op	  398569 B/op	    7042 allocs/op
BenchmarkSequential/gqlgen_mapping/16-4 	     500	   2745586 ns/op	  762611 B/op	   13614 allocs/op
BenchmarkSequential/gqlgen_mapping/32-4 	     300	   4966324 ns/op	 1492646 B/op	   26766 allocs/op
BenchmarkSequential/gqlgen_mapping/64-4 	     200	   9621530 ns/op	 2965830 B/op	   53110 allocs/op
BenchmarkSequential/gqlgen_mapping/128-4         	     100	  20574390 ns/op	 5867585 B/op	  105855 allocs/op
BenchmarkSequential/gqlgen_mapping/256-4         	      30	  45215303 ns/op	11688660 B/op	  211346 allocs/op
BenchmarkSequential/gqlgen_mapping/512-4         	      20	  93131232 ns/op	23304975 B/op	  422008 allocs/op
BenchmarkSequential/gqlgen_mapping/1024-4        	       5	 230332448 ns/op	46665992 B/op	  844221 allocs/op
BenchmarkSequential/gqlgen_mapping/2048-4        	       3	 525808672 ns/op	93186344 B/op	 1688205 allocs/op
BenchmarkSequential/gqlgen_mapping/4096-4        	       1	1540528072 ns/op	186417408 B/op	 3375830 allocs/op
BenchmarkSequential/gqlgen_mapping/8192-4        	       1	2869176853 ns/op	375336864 B/op	 6759476 allocs/op
BenchmarkSequential/gqlgen_mapping/16384-4       	       1	4511867405 ns/op	748286376 B/op	13504919 allocs/op
BenchmarkSequential/gqlgen_mapping/32768-4       	       1	7762207350 ns/op	1491152304 B/op	26993144 allocs/op
BenchmarkSequential/gqlgen_mapping/65536-4       	       1	17979894118 ns/op	2974402360 B/op	53905120 allocs/op
BenchmarkSequential/gqlgen_service_models/1-4    	    2000	    718550 ns/op	   83399 B/op	    1277 allocs/op
BenchmarkSequential/gqlgen_service_models/2-4    	    2000	    818727 ns/op	  126266 B/op	    2090 allocs/op
BenchmarkSequential/gqlgen_service_models/4-4    	    2000	   1225783 ns/op	  211569 B/op	    3715 allocs/op
BenchmarkSequential/gqlgen_service_models/8-4    	    1000	   2110830 ns/op	  383040 B/op	    6965 allocs/op
BenchmarkSequential/gqlgen_service_models/16-4   	     300	   5158427 ns/op	  728950 B/op	   13463 allocs/op
BenchmarkSequential/gqlgen_service_models/32-4   	     200	   8359651 ns/op	 1427820 B/op	   26460 allocs/op
BenchmarkSequential/gqlgen_service_models/64-4   	     100	  15946061 ns/op	 2844169 B/op	   52465 allocs/op
BenchmarkSequential/gqlgen_service_models/128-4  	      50	  27825090 ns/op	 5703355 B/op	  104516 allocs/op
BenchmarkSequential/gqlgen_service_models/256-4  	      20	  71622082 ns/op	11377691 B/op	  208754 allocs/op
BenchmarkSequential/gqlgen_service_models/512-4  	      10	 154239764 ns/op	22789916 B/op	  417226 allocs/op
BenchmarkSequential/gqlgen_service_models/1024-4 	       5	 248064692 ns/op	45548985 B/op	  834776 allocs/op
BenchmarkSequential/gqlgen_service_models/2048-4 	       2	 553738419 ns/op	91042808 B/op	 1669509 allocs/op
BenchmarkSequential/gqlgen_service_models/4096-4 	       1	1145676857 ns/op	182214344 B/op	 3340972 allocs/op
BenchmarkSequential/gqlgen_service_models/8192-4 	       1	2175438137 ns/op	364206672 B/op	 6679491 allocs/op
BenchmarkSequential/gqlgen_service_models/16384-4         	       1	4580355494 ns/op	727940224 B/op	13353432 allocs/op
BenchmarkSequential/gqlgen_service_models/32768-4         	       1	10952827201 ns/op	1454449856 B/op	26688192 allocs/op
BenchmarkSequential/gqlgen_service_models/65536-4         	       1	19232133199 ns/op	2908143920 B/op	53354275 allocs/op
BenchmarkSequential/rest_json_service_models/1-4          	   10000	    220663 ns/op	   40556 B/op	     399 allocs/op
BenchmarkSequential/rest_json_service_models/2-4          	   10000	    213073 ns/op	   43143 B/op	     399 allocs/op
BenchmarkSequential/rest_json_service_models/4-4          	   10000	    248513 ns/op	   47788 B/op	     399 allocs/op
BenchmarkSequential/rest_json_service_models/8-4          	    5000	    344532 ns/op	   56878 B/op	     399 allocs/op
BenchmarkSequential/rest_json_service_models/16-4         	    3000	    510095 ns/op	   79503 B/op	     399 allocs/op
BenchmarkSequential/rest_json_service_models/32-4         	    2000	   1126190 ns/op	  113243 B/op	     399 allocs/op
BenchmarkSequential/rest_json_service_models/64-4         	    1000	   1948932 ns/op	  179710 B/op	     399 allocs/op
BenchmarkSequential/rest_json_service_models/128-4        	     500	   2746052 ns/op	  320482 B/op	     399 allocs/op
BenchmarkSequential/rest_json_service_models/256-4        	     300	   5279649 ns/op	  611647 B/op	     399 allocs/op
BenchmarkSequential/rest_json_service_models/512-4        	     100	  10045755 ns/op	 1241718 B/op	     399 allocs/op
BenchmarkSequential/rest_json_service_models/1024-4       	     100	  19359041 ns/op	 2436470 B/op	     399 allocs/op
BenchmarkSequential/rest_json_service_models/2048-4       	      30	  42286811 ns/op	 5407928 B/op	     401 allocs/op
BenchmarkSequential/rest_json_service_models/4096-4       	      20	  85679704 ns/op	11743544 B/op	     402 allocs/op
BenchmarkSequential/rest_json_service_models/8192-4       	      10	 251322070 ns/op	21498775 B/op	     401 allocs/op
BenchmarkSequential/rest_json_service_models/16384-4      	       2	 506971459 ns/op	73904812 B/op	     412 allocs/op
BenchmarkSequential/rest_json_service_models/32768-4      	       2	 821247658 ns/op	225111292 B/op	     423 allocs/op
BenchmarkSequential/rest_json_service_models/65536-4      	       1	1657592469 ns/op	450121216 B/op	     425 allocs/op
BenchmarkSequential/rest_json_mapping/1-4                 	   10000	    255082 ns/op	   41916 B/op	     408 allocs/op
BenchmarkSequential/rest_json_mapping/2-4                 	    5000	    265825 ns/op	   45875 B/op	     416 allocs/op
BenchmarkSequential/rest_json_mapping/4-4                 	   10000	    281583 ns/op	   53230 B/op	     432 allocs/op
BenchmarkSequential/rest_json_mapping/8-4                 	    5000	    380589 ns/op	   67562 B/op	     464 allocs/op
BenchmarkSequential/rest_json_mapping/16-4                	    3000	    603017 ns/op	  100935 B/op	     528 allocs/op
BenchmarkSequential/rest_json_mapping/32-4                	    2000	   1216774 ns/op	  147736 B/op	     656 allocs/op
BenchmarkSequential/rest_json_mapping/64-4                	     500	   2360335 ns/op	  259649 B/op	     912 allocs/op
BenchmarkSequential/rest_json_mapping/128-4               	     300	   4030472 ns/op	  482784 B/op	    1424 allocs/op
BenchmarkSequential/rest_json_mapping/256-4               	     200	   7693714 ns/op	  953235 B/op	    2448 allocs/op
BenchmarkSequential/rest_json_mapping/512-4               	     100	  13451408 ns/op	 1927002 B/op	    4497 allocs/op
BenchmarkSequential/rest_json_mapping/1024-4              	     100	  24244902 ns/op	 4426965 B/op	    8595 allocs/op
BenchmarkSequential/rest_json_mapping/2048-4              	      30	  48864469 ns/op	 9762991 B/op	   16790 allocs/op
BenchmarkSequential/rest_json_mapping/4096-4              	      10	 125946166 ns/op	19468227 B/op	   33175 allocs/op
BenchmarkSequential/rest_json_mapping/8192-4              	      10	 184279407 ns/op	46477464 B/op	   65947 allocs/op
BenchmarkSequential/rest_json_mapping/16384-4             	       3	 356043357 ns/op	80220904 B/op	  131480 allocs/op
BenchmarkSequential/rest_json_mapping/32768-4             	       2	 706879172 ns/op	185688216 B/op	  262556 allocs/op
BenchmarkSequential/rest_json_mapping/65536-4             	       1	1474367522 ns/op	523144112 B/op	  524714 allocs/op
```

![Time sequential](benchmark/time_seq.jpg?raw=true "Title")

![Bytes sequential](benchmark/bytes_seq.jpg?raw=true "Title")

![Allocs sequential](benchmark/allocs_seq.jpg?raw=true "Title")

- Parallel benchmarks

```
BenchmarkParallel/gophers/1-4         	    2000	    620547 ns/op	  355187 B/op	    1915 allocs/op
BenchmarkParallel/gophers/2-4         	    2000	    735243 ns/op	  392269 B/op	    2598 allocs/op
BenchmarkParallel/gophers/4-4         	    2000	    908721 ns/op	  467313 B/op	    3965 allocs/op
BenchmarkParallel/gophers/8-4         	    1000	   1320748 ns/op	  617164 B/op	    6696 allocs/op
BenchmarkParallel/gophers/16-4        	    1000	   2046836 ns/op	  919867 B/op	   12158 allocs/op
BenchmarkParallel/gophers/32-4        	     500	   3564853 ns/op	 1527481 B/op	   23079 allocs/op
BenchmarkParallel/gophers/64-4        	     200	   6565882 ns/op	 2737354 B/op	   44923 allocs/op
BenchmarkParallel/gophers/128-4       	     100	  12226520 ns/op	 5110264 B/op	   88607 allocs/op
BenchmarkParallel/gophers/256-4       	     100	  22011785 ns/op	 9875158 B/op	  175959 allocs/op
BenchmarkParallel/gophers/512-4       	      30	  41652129 ns/op	19358817 B/op	  350656 allocs/op
BenchmarkParallel/gophers/1024-4      	      20	  82673359 ns/op	38441269 B/op	  700059 allocs/op
BenchmarkParallel/gophers/2048-4      	      10	 169930150 ns/op	76516611 B/op	 1398836 allocs/op
BenchmarkParallel/gophers/4096-4      	       5	 353817328 ns/op	152663491 B/op	 2796322 allocs/op
BenchmarkParallel/gophers/8192-4      	       1	1311121753 ns/op	304969656 B/op	 5591413 allocs/op
BenchmarkParallel/gophers/16384-4     	       1	2560326657 ns/op	609481688 B/op	11181341 allocs/op
BenchmarkParallel/gophers/32768-4     	       1	4926283359 ns/op	1218615144 B/op	22361723 allocs/op
BenchmarkParallel/gophers/65536-4     	       1	9650229549 ns/op	2436815528 B/op	44722238 allocs/op
BenchmarkParallel/gqlgen_mapping/1-4  	   10000	    199749 ns/op	   84752 B/op	    1288 allocs/op
BenchmarkParallel/gqlgen_mapping/2-4  	    5000	    327173 ns/op	  129247 B/op	    2111 allocs/op
BenchmarkParallel/gqlgen_mapping/4-4  	    3000	    552745 ns/op	  218151 B/op	    3755 allocs/op
BenchmarkParallel/gqlgen_mapping/8-4  	    2000	   1048612 ns/op	  400877 B/op	    7044 allocs/op
BenchmarkParallel/gqlgen_mapping/16-4 	    1000	   2107456 ns/op	  768201 B/op	   13631 allocs/op
BenchmarkParallel/gqlgen_mapping/32-4 	     300	   4731365 ns/op	 1500058 B/op	   26807 allocs/op
BenchmarkParallel/gqlgen_mapping/64-4 	     200	   9564012 ns/op	 2962794 B/op	   53157 allocs/op
BenchmarkParallel/gqlgen_mapping/128-4         	     100	  19505478 ns/op	 5851783 B/op	  105847 allocs/op
BenchmarkParallel/gqlgen_mapping/256-4         	      30	  42019089 ns/op	11696329 B/op	  211366 allocs/op
BenchmarkParallel/gqlgen_mapping/512-4         	      20	  78081294 ns/op	23240824 B/op	  421840 allocs/op
BenchmarkParallel/gqlgen_mapping/1024-4        	      10	 178140111 ns/op	46750911 B/op	  844381 allocs/op
BenchmarkParallel/gqlgen_mapping/2048-4        	       2	 511157368 ns/op	93106512 B/op	 1687028 allocs/op
BenchmarkParallel/gqlgen_mapping/4096-4        	       1	1015767437 ns/op	186431344 B/op	 3377180 allocs/op
BenchmarkParallel/gqlgen_mapping/8192-4        	       1	2143728186 ns/op	372501216 B/op	 6746802 allocs/op
BenchmarkParallel/gqlgen_mapping/16384-4       	       1	5110984750 ns/op	746821480 B/op	13501416 allocs/op
BenchmarkParallel/gqlgen_mapping/32768-4       	       1	7890208410 ns/op	1491574568 B/op	26977949 allocs/op
BenchmarkParallel/gqlgen_mapping/65536-4       	       1	16465038921 ns/op	2976943488 B/op	53918930 allocs/op
BenchmarkParallel/gqlgen_service_models/1-4    	   10000	    279856 ns/op	   83374 B/op	    1277 allocs/op
BenchmarkParallel/gqlgen_service_models/2-4    	    2000	    536347 ns/op	  126193 B/op	    2090 allocs/op
BenchmarkParallel/gqlgen_service_models/4-4    	    2000	    948314 ns/op	  211413 B/op	    3715 allocs/op
BenchmarkParallel/gqlgen_service_models/8-4    	    2000	   1454049 ns/op	  382423 B/op	    6965 allocs/op
BenchmarkParallel/gqlgen_service_models/16-4   	     500	   2770080 ns/op	  726395 B/op	   13465 allocs/op
BenchmarkParallel/gqlgen_service_models/32-4   	     200	   8908655 ns/op	 1425410 B/op	   26469 allocs/op
BenchmarkParallel/gqlgen_service_models/64-4   	     100	  16798154 ns/op	 2840005 B/op	   52502 allocs/op
BenchmarkParallel/gqlgen_service_models/128-4  	      50	  32039046 ns/op	 5711524 B/op	  104641 allocs/op
BenchmarkParallel/gqlgen_service_models/256-4  	      20	  62619995 ns/op	11449112 B/op	  209024 allocs/op
BenchmarkParallel/gqlgen_service_models/512-4  	      10	 133560484 ns/op	22706757 B/op	  417465 allocs/op
BenchmarkParallel/gqlgen_service_models/1024-4 	       5	 222398081 ns/op	45583044 B/op	  835160 allocs/op
BenchmarkParallel/gqlgen_service_models/2048-4 	       2	 510075610 ns/op	91093348 B/op	 1670031 allocs/op
BenchmarkParallel/gqlgen_service_models/4096-4 	       2	 849268673 ns/op	182149444 B/op	 3339839 allocs/op
BenchmarkParallel/gqlgen_service_models/8192-4 	       1	2143680101 ns/op	364233928 B/op	 6679727 allocs/op
BenchmarkParallel/gqlgen_service_models/16384-4         	       1	5375373960 ns/op	727806504 B/op	13351489 allocs/op
BenchmarkParallel/gqlgen_service_models/32768-4         	       1	13463517542 ns/op	1453716744 B/op	26678039 allocs/op
BenchmarkParallel/gqlgen_service_models/65536-4         	       1	16131221543 ns/op	2909032368 B/op	53355032 allocs/op
BenchmarkParallel/rest_json_service_models/1-4          	   20000	     96574 ns/op	   40563 B/op	     399 allocs/op
BenchmarkParallel/rest_json_service_models/2-4          	   10000	    113189 ns/op	   43171 B/op	     399 allocs/op
BenchmarkParallel/rest_json_service_models/4-4          	   10000	    134728 ns/op	   47869 B/op	     399 allocs/op
BenchmarkParallel/rest_json_service_models/8-4          	   10000	    150720 ns/op	   57041 B/op	     399 allocs/op
BenchmarkParallel/rest_json_service_models/16-4         	   10000	    276788 ns/op	   79897 B/op	     399 allocs/op
BenchmarkParallel/rest_json_service_models/32-4         	    3000	    467675 ns/op	  113940 B/op	     399 allocs/op
BenchmarkParallel/rest_json_service_models/64-4         	    2000	    646213 ns/op	  181133 B/op	     399 allocs/op
BenchmarkParallel/rest_json_service_models/128-4        	    1000	   1207533 ns/op	  327914 B/op	     399 allocs/op
BenchmarkParallel/rest_json_service_models/256-4        	    1000	   2659832 ns/op	  623489 B/op	     399 allocs/op
BenchmarkParallel/rest_json_service_models/512-4        	     300	   5718982 ns/op	 1225445 B/op	     399 allocs/op
BenchmarkParallel/rest_json_service_models/1024-4       	     100	  10427760 ns/op	 2630972 B/op	     400 allocs/op
BenchmarkParallel/rest_json_service_models/2048-4       	      50	  21247801 ns/op	 5407866 B/op	     401 allocs/op
BenchmarkParallel/rest_json_service_models/4096-4       	      30	  40291984 ns/op	11420457 B/op	     402 allocs/op
BenchmarkParallel/rest_json_service_models/8192-4       	      20	  79473150 ns/op	25370700 B/op	     403 allocs/op
BenchmarkParallel/rest_json_service_models/16384-4      	      10	 142333275 ns/op	66164750 B/op	     408 allocs/op
BenchmarkParallel/rest_json_service_models/32768-4      	       3	 335053860 ns/op	225111101 B/op	     424 allocs/op
BenchmarkParallel/rest_json_service_models/65536-4      	       1	1388714867 ns/op	450121424 B/op	     433 allocs/op
BenchmarkParallel/rest_json_mapping/1-4                 	   10000	    129356 ns/op	   41929 B/op	     408 allocs/op
BenchmarkParallel/rest_json_mapping/2-4                 	   10000	    128994 ns/op	   45907 B/op	     416 allocs/op
BenchmarkParallel/rest_json_mapping/4-4                 	   10000	    148024 ns/op	   53316 B/op	     432 allocs/op
BenchmarkParallel/rest_json_mapping/8-4                 	   10000	    184745 ns/op	   67774 B/op	     464 allocs/op
BenchmarkParallel/rest_json_mapping/16-4                	    5000	    304801 ns/op	  101117 B/op	     528 allocs/op
BenchmarkParallel/rest_json_mapping/32-4                	    3000	    481408 ns/op	  148723 B/op	     656 allocs/op
BenchmarkParallel/rest_json_mapping/64-4                	    2000	    760389 ns/op	  261225 B/op	     912 allocs/op
BenchmarkParallel/rest_json_mapping/128-4               	    1000	   1367161 ns/op	  497221 B/op	    1424 allocs/op
BenchmarkParallel/rest_json_mapping/256-4               	     500	   2597331 ns/op	  978779 B/op	    2449 allocs/op
BenchmarkParallel/rest_json_mapping/512-4               	     200	   5759142 ns/op	 1999123 B/op	    4497 allocs/op
BenchmarkParallel/rest_json_mapping/1024-4              	     100	  11777844 ns/op	 4474713 B/op	    8595 allocs/op
BenchmarkParallel/rest_json_mapping/2048-4              	      50	  20681919 ns/op	 9381580 B/op	   16789 allocs/op
BenchmarkParallel/rest_json_mapping/4096-4              	      30	  42236347 ns/op	23273251 B/op	   33177 allocs/op
BenchmarkParallel/rest_json_mapping/8192-4              	      20	  94801016 ns/op	46477139 B/op	   65946 allocs/op
BenchmarkParallel/rest_json_mapping/16384-4             	       5	 213897084 ns/op	115669144 B/op	  131491 allocs/op
BenchmarkParallel/rest_json_mapping/32768-4             	       3	 421200659 ns/op	261622253 B/op	  262569 allocs/op
BenchmarkParallel/rest_json_mapping/65536-4             	       1	1529631577 ns/op	523144320 B/op	  524722 allocs/op
```

![Time Parallel](benchmark/time_par.jpg?raw=true "Title")

Request: [medium resolve scale](benchmark/resolveScale_medium.txt)

```
BenchmarkSequential/gophers/1-4         	    2000	    790771 ns/op	  203665 B/op	    1291 allocs/op
BenchmarkSequential/gophers/2-4         	    2000	    934699 ns/op	  228071 B/op	    1674 allocs/op
BenchmarkSequential/gophers/4-4         	    1000	   1111242 ns/op	  264803 B/op	    2438 allocs/op
BenchmarkSequential/gophers/8-4         	    1000	   1714341 ns/op	  352819 B/op	    3967 allocs/op
BenchmarkSequential/gophers/16-4        	     500	   2653273 ns/op	  529935 B/op	    7025 allocs/op
BenchmarkSequential/gophers/32-4        	     500	   4751551 ns/op	  892567 B/op	   13139 allocs/op
BenchmarkSequential/gophers/64-4        	     200	   8301849 ns/op	 1633136 B/op	   25365 allocs/op
BenchmarkSequential/gophers/128-4       	     100	  16663703 ns/op	 3156155 B/op	   49819 allocs/op
BenchmarkSequential/gophers/256-4       	      50	  31835216 ns/op	 6108886 B/op	   98724 allocs/op
BenchmarkSequential/gophers/512-4       	      20	  64605595 ns/op	12029508 B/op	  196537 allocs/op
BenchmarkSequential/gophers/1024-4      	      10	 118805587 ns/op	23810060 B/op	  392133 allocs/op
BenchmarkSequential/gophers/2048-4      	       5	 207807100 ns/op	47406424 B/op	  783312 allocs/op
BenchmarkSequential/gophers/4096-4      	       3	 399519163 ns/op	94546224 B/op	 1565661 allocs/op
BenchmarkSequential/gophers/8192-4      	       2	 764445071 ns/op	188825668 B/op	 3130359 allocs/op
BenchmarkSequential/gophers/16384-4     	       1	1523483938 ns/op	377378960 B/op	 6259761 allocs/op
BenchmarkSequential/gophers/32768-4     	       1	3015585860 ns/op	754462640 B/op	12518505 allocs/op
BenchmarkSequential/gophers/65536-4     	       1	6154044163 ns/op	1508618192 B/op	25036004 allocs/op
BenchmarkSequential/gqlgen_mapping/1-4  	    5000	    441716 ns/op	   64382 B/op	     933 allocs/op
BenchmarkSequential/gqlgen_mapping/2-4  	    3000	    586492 ns/op	   86771 B/op	    1400 allocs/op
BenchmarkSequential/gqlgen_mapping/4-4  	    2000	    686304 ns/op	  138017 B/op	    2336 allocs/op
BenchmarkSequential/gqlgen_mapping/8-4  	    2000	   1140001 ns/op	  242003 B/op	    4206 allocs/op
BenchmarkSequential/gqlgen_mapping/16-4 	    1000	   1825560 ns/op	  452133 B/op	    7946 allocs/op
BenchmarkSequential/gqlgen_mapping/32-4 	     500	   3287627 ns/op	  878065 B/op	   15426 allocs/op
BenchmarkSequential/gqlgen_mapping/64-4 	     200	   6072255 ns/op	 1749192 B/op	   30407 allocs/op
BenchmarkSequential/gqlgen_mapping/128-4         	     100	  12888064 ns/op	 3482742 B/op	   60428 allocs/op
BenchmarkSequential/gqlgen_mapping/256-4         	      50	  31698256 ns/op	 6894832 B/op	  120462 allocs/op
BenchmarkSequential/gqlgen_mapping/512-4         	      20	  66516260 ns/op	13704911 B/op	  240367 allocs/op
BenchmarkSequential/gqlgen_mapping/1024-4        	      10	 147448101 ns/op	27330003 B/op	  479893 allocs/op
BenchmarkSequential/gqlgen_mapping/2048-4        	       3	 471148219 ns/op	54674205 B/op	  960034 allocs/op
BenchmarkSequential/gqlgen_mapping/4096-4        	       2	 774712699 ns/op	109293564 B/op	 1920025 allocs/op
BenchmarkSequential/gqlgen_mapping/8192-4        	       1	2128044208 ns/op	220803424 B/op	 3853381 allocs/op
BenchmarkSequential/gqlgen_mapping/16384-4       	       1	3732238101 ns/op	436410864 B/op	 7664316 allocs/op
BenchmarkSequential/gqlgen_mapping/32768-4       	       1	7605906681 ns/op	875188328 B/op	15347296 allocs/op
BenchmarkSequential/gqlgen_mapping/65536-4       	       1	12155896309 ns/op	1739032288 B/op	30627412 allocs/op
BenchmarkSequential/gqlgen_service_models/1-4    	    3000	    563933 ns/op	   63031 B/op	     925 allocs/op
BenchmarkSequential/gqlgen_service_models/2-4    	    3000	    698798 ns/op	   83995 B/op	    1386 allocs/op
BenchmarkSequential/gqlgen_service_models/4-4    	    2000	    960205 ns/op	  132178 B/op	    2309 allocs/op
BenchmarkSequential/gqlgen_service_models/8-4    	    1000	   1537787 ns/op	  228981 B/op	    4154 allocs/op
BenchmarkSequential/gqlgen_service_models/16-4   	    1000	   2322814 ns/op	  424140 B/op	    7844 allocs/op
BenchmarkSequential/gqlgen_service_models/32-4   	     300	   4993229 ns/op	  819963 B/op	   15224 allocs/op
BenchmarkSequential/gqlgen_service_models/64-4   	     200	   9609885 ns/op	 1634983 B/op	   29990 allocs/op
BenchmarkSequential/gqlgen_service_models/128-4  	     100	  29965286 ns/op	 3273267 B/op	   59512 allocs/op
BenchmarkSequential/gqlgen_service_models/256-4  	      50	  43058310 ns/op	 6527845 B/op	  118647 allocs/op
BenchmarkSequential/gqlgen_service_models/512-4  	      20	  66978092 ns/op	13066435 B/op	  236755 allocs/op
BenchmarkSequential/gqlgen_service_models/1024-4 	      20	 187532782 ns/op	26077622 B/op	  473136 allocs/op
BenchmarkSequential/gqlgen_service_models/2048-4 	       5	 328759698 ns/op	52180600 B/op	  947103 allocs/op
BenchmarkSequential/gqlgen_service_models/4096-4 	       2	 708818364 ns/op	104686376 B/op	 1899273 allocs/op
BenchmarkSequential/gqlgen_service_models/8192-4 	       1	1181237669 ns/op	209186080 B/op	 3797007 allocs/op
BenchmarkSequential/gqlgen_service_models/16384-4         	       1	2557000090 ns/op	418152528 B/op	 7591798 allocs/op
BenchmarkSequential/gqlgen_service_models/32768-4         	       1	4566109082 ns/op	834725112 B/op	15163281 allocs/op
BenchmarkSequential/gqlgen_service_models/65536-4         	       1	12575426640 ns/op	1668572104 B/op	30291510 allocs/op
```

## Run the profiler

```
cd presenters/benchmark/
go test -benchmem -cpuprofile restsm65536.pprof -test.benchtime 10s -run=^$ rfc/presenters/benchmark -bench BenchmarkCandidate_rest_servicemodels_65536_high
mv benchmark.test restsm65536.test
go tool pprof -http=:6060 restsm65536.test restsm65536.pprof
```

- Rest service models

![Rest service models](benchmark/restsm65536.png?raw=true "Title")

- Gophers

![Gophers](benchmark/gophers65536.png?raw=true "Title")

- Gqlgen

![Gqlgen](benchmark/gqlgen65536.png?raw=true "Title")
