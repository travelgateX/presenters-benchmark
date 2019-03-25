package graphResolver

import (
	"github.com/travelgateX/presenters-benchmark/pkg/common"
)

type WarningResolver struct {
	warning *common.AdviseMessage
}

func (e *WarningResolver) Type() string {
	return e.warning.Type
}

func (e *WarningResolver) Code() string {
	return e.warning.Code
}

func (e *WarningResolver) Description() string {
	return e.warning.Description
}
