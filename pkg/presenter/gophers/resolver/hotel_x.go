package graphResolver

import (
	"context"

	"github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"

	"github.com/travelgateX/presenters-benchmark/pkg/search"
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
