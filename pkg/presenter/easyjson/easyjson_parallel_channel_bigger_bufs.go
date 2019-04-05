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

type CandidateParallelChannelBiggerBuf struct{}

var _ presenter.CandidateHandlerFunc = (*CandidateParallelChannelBiggerBuf)(nil)

func (CandidateParallelChannelBiggerBuf) HandlerFunc(options []*presenter.Option) (http.HandlerFunc, error) {
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

		intBuf := NewGzipBuffer2(w)
		intBuf.Write([]byte(`{"data": {"hotelX": {"search": {"options": [`))

		parallelism := 8
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
					_, err := jw.DumpTo(intBuf)
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

		intBuf.Close()
		return
	}, nil
}

func (CandidateParallelChannelBiggerBuf) UnmarshalOptions(b []byte) ([]*presenter.Option, error) {
	return presenter.JSONUnmarshalOptions(b)
}

type GzipBuffer2 struct {
	w    http.ResponseWriter
	c    chan *bytes.Buffer
	done chan struct{}
	mu   sync.Mutex
	buf  *bytes.Buffer
}

func NewGzipBuffer2(w http.ResponseWriter) *GzipBuffer2 {
	gzb := &GzipBuffer2{
		w:    w,
		c:    make(chan *bytes.Buffer, 32),
		done: make(chan struct{}),
		buf:  GetBufferChan(),
	}

	go func() {
		for b := range gzb.c {
			io.Copy(w, b)
			PutBufferChan(b)
		}
		gzb.done <- struct{}{}
	}()

	return gzb
}

func (gz *GzipBuffer2) Write(p []byte) (int, error) {
	gz.mu.Lock()
	gz.buf.Write(p)
	if gz.buf.Len() >= 203000 {
		gz.c <- gz.buf
		gz.buf = GetBufferChan()
	}
	gz.mu.Unlock()
	return len(p), nil
}

func (gz *GzipBuffer2) Close() error {
	if gz.buf.Len() > 0 {
		gz.c <- gz.buf
	}
	close(gz.c)
	<-gz.done
	return nil
}
