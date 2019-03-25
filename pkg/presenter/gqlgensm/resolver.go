package gqlgensm

import (
	"context"
	"github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"
)

type Resolver struct {
	Options []*domainHotelCommon.Option
}

func (r *Resolver) HotelOptionSearch() HotelOptionSearchResolver {
	return &hotelOptionSearchResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type hotelOptionSearchResolver struct{ *Resolver }

func (r *hotelOptionSearchResolver) AddOns(ctx context.Context, obj *domainHotelCommon.Option) (*AddOns, error) {
	return nil, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) HotelX(ctx context.Context) (*HotelXQuery, error) {
	return &HotelXQuery{
		Search: &HotelSearch{
			Options: r.Options,
		},
	}, nil
}
