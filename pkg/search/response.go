package search

import (
	"presenters-benchmark/pkg/common"
	"presenters-benchmark/pkg/domainHotelCommon"
)

type HotelSearchRS struct {
	common.BaseRS

	Options []*domainHotelCommon.Option

	// rq is the request that this rs comes from
	RQ *HotelSearchRQ
}

func (HotelSearchRS) IsResponse() {}
