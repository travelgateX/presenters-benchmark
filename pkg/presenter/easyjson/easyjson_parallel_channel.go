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

type CandidateParallelChannel struct{}

var _ presenter.CandidateHandlerFunc = (*CandidateParallelChannel)(nil)

func (CandidateParallelChannel) HandlerFunc(options []*presenter.Option) (http.HandlerFunc, error) {
	buffer.Init(
		buffer.PoolConfig{
			StartSize:  2050,
			PooledSize: 2050,
			MaxSize:    2050 * 10,
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

		intBuf := NewGzipBuffer(w)
		intBuf.Write([]byte(`{"data": {"hotelX": {"search": {"options": [`))

		parallelism := 10
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

func (CandidateParallelChannel) UnmarshalOptions(b []byte) ([]*presenter.Option, error) {
	return presenter.JSONUnmarshalOptions(b)
}

type GzipBuffer struct {
	w    http.ResponseWriter
	c    chan *bytes.Buffer
	done chan struct{}
}

func NewGzipBuffer(w http.ResponseWriter) *GzipBuffer {
	gzb := &GzipBuffer{
		w:    w,
		c:    make(chan *bytes.Buffer, 4),
		done: make(chan struct{}),
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

func (gz *GzipBuffer) Write(p []byte) (int, error) {
	buf := GetBufferChan()
	buf.Write(p)
	gz.c <- buf
	return len(p), nil
}

func (gz *GzipBuffer) Close() error {
	close(gz.c)
	<-gz.done
	return nil
}

func GetBufferChan() *bytes.Buffer {
	return buffersChan.Get().(*bytes.Buffer)
}

// PutBuffer returns a buffer to the pool
func PutBufferChan(buf *bytes.Buffer) {
	buf.Reset()
	buffersChan.Put(buf)
}

var buffersChan = sync.Pool{
	// New is called when a new instance is needed
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}
