package graphResolver

import "rfc/presenters/pkg/domainHotelCommon"

type ResortResolver struct {
	Resort *domainHotelCommon.Resort
}

func (r *ResortResolver) Code() string {
	return (*r.Resort).Code
}

func (r *ResortResolver) Name() *string {
	return (*r.Resort).Name
}

func (r *ResortResolver) Description() *string {
	return (*r.Resort).Description
}
