package search

import (
	"rfc/presenters/pkg/common"
	"rfc/presenters/pkg/domainHotelCommon"
)

type HotelSearchRS struct {
	common.BaseRS

	Options []*domainHotelCommon.Option

	// rq is the request that this rs comes from
	RQ *HotelSearchRQ
}

func (HotelSearchRS) IsResponse() {}
