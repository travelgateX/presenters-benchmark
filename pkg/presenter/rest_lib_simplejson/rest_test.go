package rest_simplejson

import (
	"presenters-benchmark/pkg/presenter"
	"testing"
)

func TestCandidate(t *testing.T) {
	presenter.TestCandidateHandleFunc(t, Candidate{})
}
