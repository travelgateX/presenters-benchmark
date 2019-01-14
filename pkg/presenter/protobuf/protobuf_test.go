package protobuf

import (
	"presenters-benchmark/pkg/presenter"
	"testing"
)

// this would need reverse parsing to pass...
func TestCandidate(t *testing.T) {
	presenter.TestCandidateHandleFunc(t, Candidate{})
}
