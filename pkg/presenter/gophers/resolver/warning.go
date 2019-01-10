package graphResolver

import (
	"rfc/presenters/pkg/common"
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
