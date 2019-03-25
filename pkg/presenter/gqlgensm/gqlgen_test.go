package gqlgensm

import (
	"github.com/travelgateX/presenters-benchmark/pkg/presenter"
	"testing"
)

// gqlgen is not able to print nullable arrays as null, it prints "[]" and doesn't pass the test
func TestCandidate(t *testing.T) {
	presenter.TestCandidateHandleFunc(t, Candidate{})
}
