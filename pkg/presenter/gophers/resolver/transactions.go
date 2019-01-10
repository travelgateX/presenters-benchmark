package graphResolver

import (
	"hub-aggregator/common/graphql"
	"rfc/presenters/pkg/common"
)

type TransactionsResolver struct {
	Transactions *common.Transactions
}

func (r *TransactionsResolver) Request() string {
	return r.Transactions.Request
}

func (r *TransactionsResolver) Response() string {
	return r.Transactions.Response
}

func (r *TransactionsResolver) TimeStamp() graphql.DateTime {
	return graphql.DateTime(r.Transactions.TimeStamp)
}
