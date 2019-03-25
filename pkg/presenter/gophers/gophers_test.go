package gophers

import (
	"github.com/travelgateX/presenters-benchmark/pkg/presenter"
	"testing"
)

func TestCandidate(t *testing.T) {
	presenter.TestCandidateHandleFunc(t, Candidate{})
}
