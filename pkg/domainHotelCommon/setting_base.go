package domainHotelCommon

type SettingsBase struct {
	Timeout           *int
	AuditTransactions *bool
	BusinessRules     *BusinessRules
	Currency          *string
}
