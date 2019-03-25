package search

import (
	"github.com/travelgateX/presenters-benchmark/pkg/common"
	"github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"
)

type HotelSearchRS struct {
	common.BaseRS

	Options []*domainHotelCommon.Option

	// rq is the request that this rs comes from
	RQ *HotelSearchRQ
}

func (HotelSearchRS) IsResponse() {}
