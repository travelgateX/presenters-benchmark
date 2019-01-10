package graphResolver

import (
	"context"
	"hub-aggregator/common/graphservice/hotelx"
	"rfc/presenters/pkg/common"
	"rfc/presenters/pkg/search"
)

type HotelSearchResolver struct {
	rq *search.HotelSearchRQ
	rs *search.HotelSearchRS

	errors   []*common.AdviseMessage
	warnings []*common.AdviseMessage

	serializedCriteria string
}

// GraphQl Resolvers
func (r *HotelSearchResolver) Context() *string {
	if r.rq.Settings == nil || r.rq.Settings.Context == nil {
		return nil
	}
	return r.rq.Settings.Context
}

func (r *HotelSearchResolver) Stats(args struct{ Token string }) *hotelx.StatsResolver {
	if r.rs == nil || args.Token != "sucLH5TC8NPbDAH5gBgaJLQ0BE0IqnYnk8" {
		return nil
	}
	return &hotelx.StatsResolver{Stats: r.rs.Stats}
}

func (r *HotelSearchResolver) AuditData() *AuditDataResolver {
	if r.rs == nil || r.rs.AuditData == nil {
		return nil
	}

	return &AuditDataResolver{AuditData: r.rs.AuditData, times: r.rs.Stats.Total}
}

func (r *HotelSearchResolver) RequestCriteria() *CriteriaRequestResolver {
	if r.rq.Criteria == nil {
		return nil
	}
	return &CriteriaRequestResolver{Criteria: r.rq.Criteria}
}

func (r *HotelSearchResolver) Options(ctx context.Context) *[]*HotelOptionResolver {
	if r.rs == nil || r.rs.Options == nil {
		return nil
	}

	options := make([]*HotelOptionResolver, 0, len(r.rs.Options))
	hotelResponse := make(map[string]struct{})

	for i := range r.rs.Options {
		opt := r.rs.Options[i]
		//Check if exists option with same id
		optResolver := &HotelOptionResolver{
			Option:   r.rs.Options[i],
			criteria: r.serializedCriteria,
		}
		options = append(options, optResolver)
		hotelResponse[opt.HotelCode] = struct{}{}
	}

	return &options
}

func (r *HotelSearchResolver) Errors() *[]*ErrorResolver {
	if r.errors == nil || len(r.errors) == 0 {
		return nil
	}

	er := make([]*ErrorResolver, len(r.errors))
	for i, e := range r.errors {
		er[i] = &ErrorResolver{error: e}
	}
	return &er
}

func (r *HotelSearchResolver) Warnings() *[]*WarningResolver {
	if r.warnings == nil || len(r.warnings) == 0 {
		return nil
	}
	wr := make([]*WarningResolver, len(r.warnings))
	for i, w := range r.warnings {
		wr[i] = &WarningResolver{warning: w}
	}
	return &wr
}
