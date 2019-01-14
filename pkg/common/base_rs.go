package common

type BaseRS struct {
	AuditData *AuditData
	Errors    []*AdviseMessage
	Warnings  []*AdviseMessage
}
