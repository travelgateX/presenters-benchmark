package easyjson

import (
	"bytes"
	"encoding/json"
	"github.com/mailru/easyjson/buffer"
	"github.com/mailru/easyjson/jwriter"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter"
	"io"
	"net/http"
	"sync"
)

type CandidateParallelGzip struct{}

var _ presenter.CandidateHandlerFunc = (*CandidateParallelGzip)(nil)

func (CandidateParallelGzip) HandlerFunc(options []*presenter.Option) (http.HandlerFunc, error) {
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

		intBuf := GetBuffer()
		intBuf.Write([]byte(`{"data": {"hotelX": {"search": {"options": [`))

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
					_, err := jw.DumpTo(intBuf)
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
		_, err := jw.DumpTo(intBuf)
		if err != nil {
			panic(err)
		}

		_, err = intBuf.Write([]byte(`],"errors": {"code": "","type": "","description": ""}}}}}`))
		if err != nil {
			panic(err)
		}

		io.Copy(w, intBuf)
		PutBuffer(intBuf)
		return
	}, nil
}

func (CandidateParallelGzip) UnmarshalOptions(b []byte) ([]*presenter.Option, error) {
	return presenter.JSONUnmarshalOptions(b)
}

var buffers = sync.Pool{
	// New is called when a new instance is needed
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func GetBuffer() *bytes.Buffer {
	return buffers.Get().(*bytes.Buffer)
}

// PutBuffer returns a buffer to the pool
func PutBuffer(buf *bytes.Buffer) {
	buf.Reset()
	buffers.Put(buf)
}
