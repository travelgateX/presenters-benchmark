package jsonitermapping

import (
	"testing"
	"presenters-benchmark/pkg/presenter"
)

func TestCandidate(t *testing.T) {
	presenter.TestCandidateHandleFunc(t, Candidate{})
}


