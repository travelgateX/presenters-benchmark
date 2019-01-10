package graphResolver

import (
	"hub-aggregator/common/graphql"
	"hub-aggregator/common/stats"
	"rfc/presenters/pkg/common"
)

const auditTimeFormat = "2006-01-02T15:04:05.9999999Z"

type AuditDataResolver struct {
	AuditData *common.AuditData
	times     stats.Times
}

func (r *AuditDataResolver) Transactions() []*TransactionsResolver {
	if len(r.AuditData.Transactions) == 0 {
		return nil
	}

	resolvers := make([]*TransactionsResolver, len(r.AuditData.Transactions))
	for i := range r.AuditData.Transactions {
		resolvers[i] = &TransactionsResolver{Transactions: &r.AuditData.Transactions[i]}
	}
	return resolvers
}

func (r *AuditDataResolver) TimeStamp() graphql.DateTime {
	if r.times.StartAtUtc == nil {
		return ""
	}
	st := r.times.StartAtUtc
	return graphql.DateTime(st.Format(auditTimeFormat))
}

func (r *AuditDataResolver) ProcessTime() float64 {
	return r.times.ElapsedTimeMsFloat64()
}
