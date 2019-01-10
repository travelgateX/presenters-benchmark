package common

import "hub-aggregator/common/stats"

type BaseRS struct {
	AuditData *AuditData
	Stats     *stats.Stats
	Errors    []*AdviseMessage
	Warnings  []*AdviseMessage
}
