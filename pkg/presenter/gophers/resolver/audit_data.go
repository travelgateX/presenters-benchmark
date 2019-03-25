package graphResolver

import (
	"fmt"
	"github.com/travelgateX/presenters-benchmark/pkg/common"
	"strconv"
)

const auditTimeFormat = "2006-01-02T15:04:05.9999999Z"

type AuditDataResolver struct {
	AuditData *common.AuditData
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

func (r *AuditDataResolver) TimeStamp() DateTime {
	return ""
}

func (r *AuditDataResolver) ProcessTime() float64 {
	return 0.0
}

type DateTime string

func (_ DateTime) ImplementsGraphQLType(name string) bool {
	return name == "DateTime"
}

func (datetime *DateTime) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		*datetime = DateTime(input)
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

func (datetime DateTime) MarshalJSON() ([]byte, error) {
	return strconv.AppendQuote(nil, string(datetime)), nil
}
