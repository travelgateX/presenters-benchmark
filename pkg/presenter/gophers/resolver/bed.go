package graphResolver

import (
	"github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"
)

type BedResolver struct {
	Bed *domainHotelCommon.Bed
}

func (r *BedResolver) Type() *string {
	return (*r.Bed).Type
}

func (r *BedResolver) Description() *string {
	return (*r.Bed).Description
}

func (r *BedResolver) Count() *int32 {
	if r.Bed.Count == nil {
		return nil
	}
	b := int32(*r.Bed.Count)
	return &b
}

func (r *BedResolver) Shared() *bool {
	return (*r.Bed).Shared
}
