package graphResolver

import (
	"github.com/travelgateX/presenters-benchmark/pkg/common"
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

func (r *TransactionsResolver) TimeStamp() DateTime {
	return DateTime(r.Transactions.TimeStamp)
}
