package graphResolver

import (
	"context"
	"github.com/travelgateX/presenters-benchmark/pkg/common"
	"github.com/travelgateX/presenters-benchmark/pkg/search"
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

func (r *HotelSearchResolver) Stats(args struct{ Token string }) *StatsResolver {
	return nil
}

func (r *HotelSearchResolver) AuditData() *AuditDataResolver {
	return nil
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

type StatsResolver struct {
}

func (r *StatsResolver) Total() *StatResolver {
	statResolver := StatResolver{}
	return &statResolver
}

func (r *StatsResolver) Validation() *StatResolver {
	statResolver := StatResolver{}
	return &statResolver
}

func (r *StatsResolver) Process() *StatResolver {
	statResolver := StatResolver{}
	return &statResolver
}
func (r *StatsResolver) Configuration() *StatResolver {
	statResolver := StatResolver{}
	return &statResolver
}

// deprecated
func (r *StatsResolver) Request() *StatResolver {
	return &StatResolver{}
}
func (r *StatsResolver) Response() *StatResolver {
	return &StatResolver{}
}
func (r *StatsResolver) RequestPlugin() *StatPluginResolver {
	statResolver := StatPluginResolver{}
	return &statResolver
}
func (r *StatsResolver) ResponsePlugin() *StatPluginResolver {
	statResolver := StatPluginResolver{}
	return &statResolver
}

func (r *StatsResolver) Accesses() []*StatAccessResolver {
	return nil
}

func (r *StatsResolver) Hotels() int32 {
	return int32(0)
}
func (r *StatsResolver) Zones() int32 {
	return int32(0)
}
func (r *StatsResolver) Cities() int32 {
	return int32(0)
}
func (r *StatsResolver) DockerID() string {
	return ""
}

type StatResolver struct {
}

func (r *StatResolver) Start() DateTime {
	return ""
}

func (r *StatResolver) End() DateTime {
	return ""
}

func (r *StatResolver) Duration() *float64 {
	return nil
}

type StatAccessResolver struct {
}

func (r *StatAccessResolver) Name() string {
	return ""
}

func (r *StatAccessResolver) Total() *StatResolver {
	statResolver := StatResolver{}
	return &statResolver
}

func (r *StatAccessResolver) StaticConfiguration() *StatResolver {
	statResolver := StatResolver{}
	return &statResolver
}

func (r *StatAccessResolver) Transactions() []*StatTransactionResolver {
	return nil
}

func (r *StatAccessResolver) Plugins() *[]*StatPluginResolver {
	return nil
}

func (r *StatAccessResolver) Hotels() int32 {
	return int32(0)
}
func (r *StatAccessResolver) Zones() int32 {
	return int32(0)
}
func (r *StatAccessResolver) Cities() int32 {
	return int32(0)
}

func (r *StatAccessResolver) RequestAccess() *StatPluginResolver {
	statPluginResolver := StatPluginResolver{}
	return &statPluginResolver
}
func (r *StatAccessResolver) ResponseAccess() *StatPluginResolver {
	statPluginResolver := StatPluginResolver{}
	return &statPluginResolver
}

type StatPluginResolver struct{}

func (r *StatPluginResolver) Name() string {
	return ""
}

func (r *StatPluginResolver) Total() *StatResolver {
	statResolver := StatResolver{}
	return &statResolver
}

type StatTransactionResolver struct {
}

func (r *StatTransactionResolver) Reference() string {
	return ""
}

func (r *StatTransactionResolver) Total() *StatResolver {
	statResolver := StatResolver{}
	return &statResolver
}

func (r *StatTransactionResolver) BuildRequest() *StatResolver {
	statResolver := StatResolver{}
	return &statResolver
}

func (r *StatTransactionResolver) WorkerCommunication() *StatResolver {
	statResolver := StatResolver{}
	return &statResolver
}

func (r *StatTransactionResolver) ParseResponse() *StatResolver {
	statResolver := StatResolver{}
	return &statResolver
}
