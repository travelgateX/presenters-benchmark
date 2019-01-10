package graphResolver

import (
	"context"

	"rfc/presenters/pkg/domainHotelCommon"

	"rfc/presenters/pkg/search"
)

type HotelXQueryResolver struct {
	Options []*domainHotelCommon.Option
}

func (r *HotelXQueryResolver) Search(ctx context.Context) *HotelSearchResolver {
	hsr := &HotelSearchResolver{rs: &search.HotelSearchRS{
		Options: r.Options,
	}}
	return hsr
}

func (r *HotelXQueryResolver) SearchStatusService() *ServiceStatusResolver {
	panic("not implemented")
}
