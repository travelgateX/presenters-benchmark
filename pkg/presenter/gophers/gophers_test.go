package gophers

import (
	"presenters-benchmark/pkg/presenter"
	"testing"
)

func TestCandidate(t *testing.T) {
	presenter.TestCandidateHandleFunc(t, Candidate{})
}
