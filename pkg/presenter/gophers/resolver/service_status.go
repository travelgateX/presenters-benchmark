package graphResolver

import (
	"github.com/travelgateX/presenters-benchmark/pkg/common"
)

type ServiceStatusResolver struct {
	am *common.AdviseMessage
}

func (e *ServiceStatusResolver) Type() *string {
	return &e.am.Type
}

func (e *ServiceStatusResolver) Code() *string {
	return &e.am.Code
}

func (e *ServiceStatusResolver) Description() *string {
	return &e.am.Description
}
