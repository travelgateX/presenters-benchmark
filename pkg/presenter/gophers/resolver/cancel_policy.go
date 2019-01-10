package graphResolver

import "rfc/presenters/pkg/domainHotelCommon"

type CancelPolicyResolver struct {
	CancelPolicy *domainHotelCommon.CancelPolicy
}

func (r *CancelPolicyResolver) Refundable() bool {
	return (*r.CancelPolicy).Refundable
}

func (r *CancelPolicyResolver) CancelPenalties() *[]*CancelPenaltyResolver {
	if r.CancelPolicy == nil || len(r.CancelPolicy.CancelPenalties) == 0 {
		return nil
	}

	cancelPenalties := make([]*CancelPenaltyResolver, 0, len(r.CancelPolicy.CancelPenalties))
	for _, cancelPolicy := range r.CancelPolicy.CancelPenalties {
		cancelPolicy_aux := cancelPolicy
		cancelPenalties = append(cancelPenalties, &CancelPenaltyResolver{CancelPenalty: &cancelPolicy_aux})
	}

	return &cancelPenalties
}
