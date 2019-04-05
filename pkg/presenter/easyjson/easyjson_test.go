package easyjson

import (
	"github.com/travelgateX/presenters-benchmark/pkg/presenter"
	"testing"
)

func TestCandidate(t *testing.T) {
	presenter.TestCandidateHandleFunc(t, Candidate{})
}

func TestCandidate_gzip(t *testing.T) {
	presenter.TestCandidateHandleFunc_Gzip(t, Candidate{})
}

func TestCandidate_Channel(t *testing.T) {
	presenter.TestCandidateChannelHandleFunc(t, Candidate{})
}

func TestCandidateParallel(t *testing.T) {
	presenter.TestCandidateHandleFunc(t, CandidateParallel{})
}

func TestCandidateParallel_gzip(t *testing.T) {
	presenter.TestCandidateHandleFunc_Gzip(t, CandidateParallel{})
}

func TestCandidateParallel_Channel(t *testing.T) {
	presenter.TestCandidateChannelHandleFunc(t, CandidateParallel{})
}

func TestCandidateParallelGzip(t *testing.T) {
	presenter.TestCandidateHandleFunc(t, CandidateParallelGzip{})
}

func TestCandidateParallelGzip_Gzip(t *testing.T) {
	presenter.TestCandidateHandleFunc_Gzip(t, CandidateParallelGzip{})
}

func TestCandidateParallelChan(t *testing.T) {
	presenter.TestCandidateHandleFunc(t, CandidateParallelChannel{})
}

func TestCandidateParallelChan_Gzip(t *testing.T) {
	presenter.TestCandidateHandleFunc_Gzip(t, CandidateParallelChannel{})
}

func TestCandidateParallelChanBiggerBuf(t *testing.T) {
	presenter.TestCandidateHandleFunc(t, CandidateParallelChannelBiggerBuf{})
}

func TestCandidateParallelChanBiggerBuf_Gzip(t *testing.T) {
	presenter.TestCandidateHandleFunc_Gzip(t, CandidateParallelChannelBiggerBuf{})
}
