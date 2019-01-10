package gophers

import (
	"rfc/presenters/pkg/presenter"
	"testing"
)

func TestCandidate(t *testing.T) {
	presenter.TestCandidateHandleFunc(t, Candidate{})
}
