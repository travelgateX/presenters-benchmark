package easyjson

import (
	"encoding/json"
	"github.com/mailru/easyjson/buffer"
	"github.com/mailru/easyjson/jwriter"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter"
	"net/http"
	"sync"
)

type CandidateParallel struct{}

var _ presenter.CandidateHandlerFunc = (*CandidateParallel)(nil)

func (CandidateParallel) HandlerFunc(options []*presenter.Option) (http.HandlerFunc, error) {
	buffer.Init(
		buffer.PoolConfig{
			StartSize:  2030,
			PooledSize: 2030,
			MaxSize:    2030 * 10,
		},
	)
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

		w.Write([]byte(`{"data": {"hotelX": {"search": {"options": [`))

		parallelism := 8
		mutex := sync.Mutex{}
		wg := &sync.WaitGroup{}
		wg.Add(parallelism)
		optionsC := make(chan *presenter.Option, parallelism)
		for i := 0; i < parallelism; i++ {
			go func() {
				jw := &jwriter.Writer{}
				for opt := range optionsC {
					opt.MarshalEasyJSON(jw)
					if jw.Error != nil {
						panic(jw.Error)
					}
					jw.RawByte(',')
					mutex.Lock()
					_, err := jw.DumpTo(w)
					mutex.Unlock()
					if err != nil {
						panic(err)
					}
				}
				wg.Done()
			}()
		}

		for _, opt := range options[:len(options)-1] {
			optionsC <- opt
		}
		close(optionsC)
		wg.Wait()
		jw := &jwriter.Writer{}
		opt := options[len(options)-1]
		opt.MarshalEasyJSON(jw)
		if jw.Error != nil {
			panic(jw.Error)
		}
		_, err := jw.DumpTo(w)
		if err != nil {
			panic(err)
		}

		_, err = w.Write([]byte(`],"errors": {"code": "","type": "","description": ""}}}}}`))
		if err != nil {
			panic(err)
		}
		return
	}, nil
}

func (CandidateParallel) UnmarshalOptions(b []byte) ([]*presenter.Option, error) {
	return presenter.JSONUnmarshalOptions(b)
}

func (CandidateParallel) ChannelHandlerFunc(optionsC <-chan *presenter.Option) (http.HandlerFunc, error) {
	buffer.Init(
		buffer.PoolConfig{
			StartSize:  2030,
			PooledSize: 2030,
			MaxSize:    2030 * 10,
		},
	)
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"data": {"hotelX": {"search": {"options": [`))

		jw := &jwriter.Writer{}
		opt := <-optionsC
		opt.MarshalEasyJSON(jw)
		if jw.Error != nil {
			panic(jw.Error)
		}
		_, err := jw.DumpTo(w)
		if err != nil {
			panic(err)
		}

		for opt := range optionsC {
			jw.RawByte(',')
			opt.MarshalEasyJSON(jw)
			if jw.Error != nil {
				panic(jw.Error)
			}
			_, err := jw.DumpTo(w)
			if err != nil {
				panic(err)
			}
		}

		_, err = w.Write([]byte(`],"errors": {"code": "","type": "","description": ""}}}}}`))
		if err != nil {
			panic(err)
		}
		return
	}, nil
}
