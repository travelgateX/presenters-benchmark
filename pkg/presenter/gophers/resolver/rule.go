package graphResolver

import (
	"github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"
)

type RuleResolver struct {
	rule *domainHotelCommon.Rule
}

func (r *RuleResolver) Id() string {
	return r.rule.Id
}

func (r *RuleResolver) Name() *string {
	return r.rule.Name
}

func (r *RuleResolver) Type() domainHotelCommon.MarkupRuleType {
	return r.rule.Type
}

func (r *RuleResolver) Value() float64 {
	return r.rule.Value
}
