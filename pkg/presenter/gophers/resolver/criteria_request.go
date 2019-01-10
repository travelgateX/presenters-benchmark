package graphResolver

import (
	"hub-aggregator/common/graphql"
	"rfc/presenters/pkg/search"
)

type CriteriaRequestResolver struct {
	Criteria *search.CriteriaSearch
}

func (r *CriteriaRequestResolver) CheckIn() graphql.Date {
	return graphql.Date(r.Criteria.CheckIn)
}

func (r *CriteriaRequestResolver) CheckOut() graphql.Date {
	return graphql.Date(r.Criteria.CheckOut)
}

func (r *CriteriaRequestResolver) Nationality() *graphql.Country {
	if r.Criteria.Nationality == nil {
		return nil
	}
	tmp := graphql.Country(*r.Criteria.Nationality)
	return &tmp
}

func (r *CriteriaRequestResolver) Currency() *graphql.Currency {
	if r.Criteria.Currency == nil {
		return nil
	}
	tmp := graphql.Currency(*r.Criteria.Currency)
	return &tmp
}

func (r *CriteriaRequestResolver) Language() *graphql.Language {
	if r.Criteria.Language == nil {
		return nil
	}
	tmp := graphql.Language(*r.Criteria.Language)
	return &tmp
}

func (r *CriteriaRequestResolver) Market() string {
	return *r.Criteria.Market
}

func (r *CriteriaRequestResolver) Hotels() []string {
	if r.Criteria.Hotels == nil {
		return []string{}
	}
	return *r.Criteria.Hotels
}

func (r *CriteriaRequestResolver) Occupancies() []*OccupancyResolver {
	occupancies := make([]*OccupancyResolver, 0, len(r.Criteria.Occupancies))
	for _, occupancy := range r.Criteria.Occupancies {
		occupancy_aux := occupancy
		occupancies = append(occupancies, &OccupancyResolver{Occupancy: &occupancy_aux})
	}
	return occupancies
}
