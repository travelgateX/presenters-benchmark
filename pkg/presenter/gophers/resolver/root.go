package graphResolver

import "rfc/presenters/pkg/domainHotelCommon"

type QueryResolver struct {
	Options []*domainHotelCommon.Option
}
type MutationResolver struct{}

func (r *QueryResolver) HotelX() *HotelXQueryResolver {

	return &HotelXQueryResolver{r.Options}
}

func (r *QueryResolver) Search() *SearchResolver {
	panic("not impl")
}
