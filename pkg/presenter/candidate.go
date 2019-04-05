package presenter

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// CandidateHandlerFunc is the interface used to benchmark
type CandidateHandlerFunc interface {
	// HandlerFunc returned must write the given options in the response body
	HandlerFunc(options []*Option) (http.HandlerFunc, error)
	// UnmarshalOptions must take the bytes of a responseBody and unmarshal them to Options
	UnmarshalOptions([]byte) ([]*Option, error)
}

type CandidateChannelHandlerFunc interface {
	CandidateHandlerFunc
	ChannelHandlerFunc(options <-chan *Option) (http.HandlerFunc, error)
}

// JSONUnmarshalOptions is a valid UnmarshalOptions func for those implementations that write
// a JSON encoded 'Response' in the response body
func JSONUnmarshalOptions(b []byte) ([]*Option, error) {
	var res Response
	err := json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}
	return res.Data.HotelX.Search.Options, nil
}

// TestCandidateHandleFunc tests that a candidate encodes options correctly
func TestCandidateHandleFunc(t *testing.T, c CandidateHandlerFunc) {
	testCandidateHandleFunc(t, c, false)
}

func TestCandidateHandleFunc_Gzip(t *testing.T, c CandidateHandlerFunc) {
	testCandidateHandleFunc(t, c, true)
}

func testCandidateHandleFunc(t *testing.T, c CandidateHandlerFunc, gzipEncode bool) {
	opts := NewOptionsGen().Gen(200)
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
	if gzipEncode {
		w, close := GzipCompress(rr)
		handler.ServeHTTP(w, req)
		close()
	} else {
		handler.ServeHTTP(rr, req)
	}

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v, body: %s", status, http.StatusOK, rr.Body.String())
	}
	// t.Logf("response: %s", rr.Body.String())
	// deserialize option, it must be the same as the generated
	var retOpts []*Option
	var responseBody []byte
	if gzipEncode {
		reader, err := gzip.NewReader(rr.Body)
		if err != nil {
			t.Fatalf("gzip new reader error: %v", err)
		}
		responseBody, err = ioutil.ReadAll(reader)
		if err != nil {
			t.Fatalf("gzip read error: %v", err)
		}
	} else {
		responseBody = rr.Body.Bytes()
	}
	retOpts, err = c.UnmarshalOptions(responseBody)
	if err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}

	if len(opts) != len(retOpts) {
		t.Fatalf("different len of options generated vs returned len: %v, %v", len(opts), len(retOpts))
	}
	for i := range opts {
		if !OptionEquals(opts[i], retOpts[i]) {
			b1, err := json.Marshal(*retOpts[i])
			if err != nil {
				t.Errorf("error marshalling opt: %v", err)
			}
			b2, err := json.Marshal(retOpts[i])
			if err != nil {
				t.Errorf("error marshalling opt: %v", err)
			}
			t.Fatalf("options are not serialized correctly: opt1: %s\nopt2: %s ", string(b1), string(b2))
		}
	}
}

func TestCandidateChannelHandleFunc(t *testing.T, c CandidateChannelHandlerFunc) {
	opts := NewOptionsGen().Gen(200)
	optsC := make(chan *Option, 200)
	for _, opt := range opts {
		optsC <- opt
	}
	close(optsC)
	hf, err := c.ChannelHandlerFunc(optsC)
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
	if true {
		w, close := GzipCompress(rr)
		handler.ServeHTTP(w, req)
		close()
	} else {
		handler.ServeHTTP(rr, req)
	}

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v, body: %s", status, http.StatusOK, rr.Body.String())
	}
	// t.Logf("response: %s", rr.Body.String())
	// deserialize option, it must be the same as the generated
	var retOpts []*Option
	var responseBody []byte
	if true {
		reader, err := gzip.NewReader(rr.Body)
		if err != nil {
			t.Fatalf("gzip new reader error: %v", err)
		}
		responseBody, err = ioutil.ReadAll(reader)
		if err != nil {
			t.Fatalf("gzip read error: %v", err)
		}
	} else {
		responseBody = rr.Body.Bytes()
	}
	retOpts, err = c.UnmarshalOptions(responseBody)
	if err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}

	if len(opts) != len(retOpts) {
		t.Fatalf("different len of options generated vs returned len: %v, %v", len(opts), len(retOpts))
	}
	for i := range opts {
		if !OptionEquals(opts[i], retOpts[i]) {
			b1, err := json.Marshal(*retOpts[i])
			if err != nil {
				t.Errorf("error marshalling opt: %v", err)
			}
			b2, err := json.Marshal(retOpts[i])
			if err != nil {
				t.Errorf("error marshalling opt: %v", err)
			}
			t.Fatalf("options are not serialized correctly: opt1: %s\nopt2: %s ", string(b1), string(b2))
		}
	}
}

type compressResponseWriter struct {
	io.Writer
	http.ResponseWriter
	http.Hijacker
	http.Flusher
	http.CloseNotifier
}

func GzipCompress(w http.ResponseWriter) (http.ResponseWriter, func() error) {
	gw, _ := gzip.NewWriterLevel(w, -1)

	h, hok := w.(http.Hijacker)
	if !hok { /* w is not Hijacker... oh well... */
		h = nil
	}

	f, fok := w.(http.Flusher)
	if !fok {
		f = nil
	}

	cn, cnok := w.(http.CloseNotifier)
	if !cnok {
		cn = nil
	}

	return &compressResponseWriter{
		Writer:         gw,
		ResponseWriter: w,
		Hijacker:       h,
		Flusher:        f,
		CloseNotifier:  cn,
	}, gw.Close
}

func (w *compressResponseWriter) WriteHeader(c int) {
	w.ResponseWriter.Header().Del("Content-Length")
	w.ResponseWriter.WriteHeader(c)
}

func (w *compressResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *compressResponseWriter) Write(b []byte) (int, error) {
	h := w.ResponseWriter.Header()
	if h.Get("Content-Type") == "" {
		h.Set("Content-Type", http.DetectContentType(b))
	}
	h.Del("Content-Length")

	return w.Writer.Write(b)
}

type flusher interface {
	Flush() error
}

func (w *compressResponseWriter) Flush() {
	// Flush compressed data if compressor supports it.
	if f, ok := w.Writer.(flusher); ok {
		f.Flush()
	}
	// Flush HTTP response.
	if w.Flusher != nil {
		w.Flusher.Flush()
	}
}
