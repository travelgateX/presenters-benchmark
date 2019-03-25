package graphResolver

import (
	"github.com/travelgateX/presenters-benchmark/pkg/common"
)

type ErrorResolver struct {
	error *common.AdviseMessage
}

func (e *ErrorResolver) Type() string {
	return e.error.Type
}

func (e *ErrorResolver) Code() string {
	return e.error.Code
}

func (e *ErrorResolver) Description() string {
	return e.error.Description
}
