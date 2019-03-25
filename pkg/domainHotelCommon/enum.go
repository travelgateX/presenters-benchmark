package domainHotelCommon

import (
	"fmt"
	"io"
	"github.com/travelgateX/presenters-benchmark/pkg/access"
	"strconv"
)

var UnknownType = fmt.Errorf("unknown input type")

type OptionType string

const (
	Hotel         OptionType = "HOTEL"
	HotelSkiPass  OptionType = "HOTEL_SKI_PASS"
	HotelEntrance OptionType = "HOTEL_ENTRANCE"
)

type StatusType string

const (
	StatusTypeOk StatusType = "OK"
	StatusTypeRq StatusType = "RQ"
)

func (e StatusType) IsValid() bool {
	switch e {
	case StatusTypeOk, StatusTypeRq:
		return true
	}
	return false
}

func (e StatusType) String() string {
	return string(e)
}

func (e *StatusType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StatusType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StatusType", str)
	}
	return nil
}

func (e StatusType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

func ToStatusType(s string) (StatusType, error) {
	var ret StatusType
	switch s {
	case "OK":
		ret = StatusTypeOk
	case "RQ":
		ret = StatusTypeRq
	default:
		return "", UnknownType
	}
	return ret, nil
}

func ToRateRulesType(s string) (access.RateRulesType, error) {
	var ret access.RateRulesType
	switch s {
	case "NonRefundable":
		ret = access.RateRulesTypeNonRefundable
	case "Package":
		ret = access.RateRulesTypePackage
	case "Older55":
		ret = access.RateRulesTypeOlder55
	case "Older60":
		ret = access.RateRulesTypeOlder60
	case "Older65":
		ret = access.RateRulesTypeOlder65
	case "CanaryResident":
		ret = access.RateRulesTypeCanaryResident
	case "BalearicResident":
		ret = access.RateRulesTypeBalearicResident
	case "largeFamily":
		ret = access.RateRulesTypeLargeFamily
	case "honeymoon":
		ret = access.RateRulesTypeHoneymoon
	default:
		return "", UnknownType
	}
	return ret, nil
}

type CancelPenaltyType string

const (
	CancelPenaltyTypeNights  CancelPenaltyType = "NIGHTS"
	CancelPenaltyTypePercent CancelPenaltyType = "PERCENT"
	CancelPenaltyTypeImport  CancelPenaltyType = "IMPORT"
)

func (e CancelPenaltyType) IsValid() bool {
	switch e {
	case CancelPenaltyTypeNights, CancelPenaltyTypePercent, CancelPenaltyTypeImport:
		return true
	}
	return false
}

func (e CancelPenaltyType) String() string {
	return string(e)
}

func (e *CancelPenaltyType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CancelPenaltyType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CancelPenaltyType", str)
	}
	return nil
}

func (e CancelPenaltyType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

func ToCancelPenaltyType(s string) (CancelPenaltyType, error) {
	var ret CancelPenaltyType
	switch s {
	case "Noches":
		ret = CancelPenaltyTypeNights
	case "Porcentaje":
		ret = CancelPenaltyTypePercent
	case "Importe":
		ret = CancelPenaltyTypeImport
	default:
		return "", UnknownType
	}
	return ret, nil
}

type PriceType string

const (
	Gross  PriceType = "GROSS"
	Net    PriceType = "NET"
	Amount PriceType = "Amount"
)

type DurationType string

const (
	DurationTypeRange DurationType = "RANGE"
	DurationTypeOpen  DurationType = "OPEN"
)

func (e DurationType) IsValid() bool {
	switch e {
	case DurationTypeRange, DurationTypeOpen:
		return true
	}
	return false
}

func (e DurationType) String() string {
	return string(e)
}

func (e *DurationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DurationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DurationType", str)
	}
	return nil
}

func (e DurationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

func ToDurationType(s string) (DurationType, error) {
	var ret DurationType
	switch s {
	case "Range":
		ret = DurationTypeRange
	case "open":
		ret = DurationTypeOpen
	default:
		return "", UnknownType
	}
	return ret, nil
}

type SupplementType string

const (
	SupplementTypeSkiPass   SupplementType = "SKI_PASS"
	SupplementTypeLessons   SupplementType = "LESSONS"
	SupplementTypeMeals     SupplementType = "MEALS"
	SupplementTypeEquipment SupplementType = "EQUIPMENT"
	SupplementTypeTicket    SupplementType = "TICKET"
	SupplementTypeTransfers SupplementType = "TRANSFERS"
	SupplementTypeGala      SupplementType = "GALA"
	SupplementTypeActivity  SupplementType = "ACTIVITY"
)

func (e SupplementType) IsValid() bool {
	switch e {
	case SupplementTypeSkiPass, SupplementTypeLessons, SupplementTypeMeals, SupplementTypeEquipment, SupplementTypeTicket, SupplementTypeTransfers, SupplementTypeGala, SupplementTypeActivity:
		return true
	}
	return false
}

func (e SupplementType) String() string {
	return string(e)
}

func (e *SupplementType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SupplementType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SupplementType", str)
	}
	return nil
}

func (e SupplementType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

func ToSupplementType(s string) (SupplementType, error) {
	var ret SupplementType
	switch s {
	case "SkiPass":
		ret = SupplementTypeSkiPass
	case "Lessons":
		ret = SupplementTypeLessons
	case "Meals":
		ret = SupplementTypeMeals
	case "Equipment":
		ret = SupplementTypeEquipment
	case "Ticket":
		ret = SupplementTypeTicket
	case "Transfers":
		ret = SupplementTypeTransfers
	case "Gala":
		ret = SupplementTypeGala
	case "Activity":
		ret = SupplementTypeActivity
	default:
		return "", UnknownType
	}
	return ret, nil
}

type ChargeType string

const (
	ChargeTypeInclude ChargeType = "INCLUDE"
	ChargeTypeExclude ChargeType = "EXCLUDE"
)

func (e ChargeType) IsValid() bool {
	switch e {
	case ChargeTypeInclude, ChargeTypeExclude:
		return true
	}
	return false
}

func (e ChargeType) String() string {
	return string(e)
}

func (e *ChargeType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ChargeType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ChargeType", str)
	}
	return nil
}

func (e ChargeType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UnitTimeType string

const (
	UnitTimeTypeDay  UnitTimeType = "DAY"
	UnitTimeTypeHour UnitTimeType = "HOUR"
)

func (e UnitTimeType) IsValid() bool {
	switch e {
	case UnitTimeTypeDay, UnitTimeTypeHour:
		return true
	}
	return false
}

func (e UnitTimeType) String() string {
	return string(e)
}

func (e *UnitTimeType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UnitTimeType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UnitTimeType", str)
	}
	return nil
}

func (e UnitTimeType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

func ToUnitType(s string) (UnitTimeType, error) {
	switch s {
	case "Day":
		return UnitTimeTypeDay, nil
	case "Hour":
		return UnitTimeTypeHour, nil
	}
	return "", UnknownType
}

type PaymentCardType string

const (
	VI PaymentCardType = ""
	AX                 = ""
	BC                 = ""
	CA                 = ""
	CB                 = ""
	CU                 = ""
	DS                 = ""
	DC                 = ""
	T                  = ""
	R                  = ""
	N                  = ""
	L                  = ""
	E                  = ""
	JC                 = ""
	TO                 = ""
	S                  = ""
	EC                 = ""
	EU                 = ""
	TP                 = ""
	OP                 = ""
	ER                 = ""
	XS                 = ""
	O                  = ""
)

type BookingStatusType string

const (
	BookKO BookingStatusType = "KO"
	BookOK BookingStatusType = "OK"
	BookRQ BookingStatusType = "ON_REQUEST"
	BookCN BookingStatusType = "CANCELLED"
	BookUN BookingStatusType = "UNKNOWN"
)

func ToBookingStatusType(s string) (BookingStatusType, error) {
	var ret BookingStatusType
	switch s {
	case "OK":
		ret = BookOK
	case "RQ":
		ret = BookRQ
	case "CN":
		ret = BookCN
	case "UN":
		ret = BookUN
	default:
		return "", UnknownType
	}
	return ret, nil
}

// Enums are typed as integers to ease the build of optionId

type BusinessRulesType string

const (
	BusinessRulesTypeCheaperAmount BusinessRulesType = "CHEAPER_AMOUNT"
	BusinessRulesTypeRoomType      BusinessRulesType = "ROOM_TYPE"
)

func (e BusinessRulesType) Code() (int, error) {
	switch e {
	case BusinessRulesTypeCheaperAmount:
		return 0, nil
	case BusinessRulesTypeRoomType:
		return 1, nil
	}
	return 1, UnknownType
}

func (e BusinessRulesType) IsValid() bool {
	switch e {
	case BusinessRulesTypeCheaperAmount, BusinessRulesTypeRoomType:
		return true
	}
	return false
}

func (e BusinessRulesType) String() string {
	return string(e)
}

func (e *BusinessRulesType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BusinessRulesType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid BusinessRulesType", str)
	}
	return nil
}

func (e BusinessRulesType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

func CheckBusinessRuleType(s BusinessRulesType) (BusinessRulesType, error) {
	var ret BusinessRulesType
	switch s {
	case BusinessRulesTypeCheaperAmount:
		return BusinessRulesTypeCheaperAmount, nil
	case BusinessRulesTypeRoomType:
		return BusinessRulesTypeRoomType, nil
	}
	return ret, UnknownType
}

func (brt BusinessRulesType) HotelApiDescription() string {
	switch brt {
	case BusinessRulesTypeCheaperAmount:
		return "CheaperAmount"
	case BusinessRulesTypeRoomType:
		return "RoomType"
	}
	return ""
}

type PaymentType string

const (
	PaymentTypeMerchant    PaymentType = "MERCHANT"
	PaymentTypeDirect      PaymentType = "DIRECT"
	PaymentTypeCardBooking PaymentType = "CARD_BOOKING"
	PaymentTypeCardCheckIn PaymentType = "CARD_CHECK_IN"
)

func (e PaymentType) IsValid() bool {
	switch e {
	case PaymentTypeMerchant, PaymentTypeDirect, PaymentTypeCardBooking, PaymentTypeCardCheckIn:
		return true
	}
	return false
}

func (e PaymentType) String() string {
	return string(e)
}

func (e *PaymentType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PaymentType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PaymentType", str)
	}
	return nil
}

func (e PaymentType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

func (e PaymentType) Code() int {
	switch e {
	case PaymentTypeMerchant:
		return 0
	case PaymentTypeDirect:
		return 1
	case PaymentTypeCardBooking:
		return 2
	case PaymentTypeCardCheckIn:
		return 3
	default:
		return -1
	}
}

func (pt PaymentType) Description() string {
	return pt.String()
}

func (e PaymentType) HotelApiDescription() string {
	switch e {
	case PaymentTypeMerchant:
		return "MerchantPay"
	case PaymentTypeDirect:
		return "LaterPay"
	case PaymentTypeCardBooking:
		return "CardBookingPay"
	case PaymentTypeCardCheckIn:
		return "CardCheckInPay"
	default:
		return "-1"
	}
}

func ToPaymentType(s string) (PaymentType, error) {
	var ret PaymentType
	switch s {
	case "pagoDirecto", "LaterPay":
		ret = PaymentTypeDirect
	case "pagoMinorista", "MerchantPay":
		ret = PaymentTypeMerchant
	case "pagoTarjetaFechaReserva", "CardBookingPay":
		ret = PaymentTypeCardBooking
	case "pagoTarjetaFechaEntrada", "CardCheckInPay":
		ret = PaymentTypeCardCheckIn
	}
	return ret, UnknownType
}

type DateType int

const (
	ArrivalDate DateType = iota
	BookingCreationDate
)

func (dt DateType) HotelApiDescription() string {
	switch dt {
	case ArrivalDate:
		return "A"
	default:
		return "B"
	}
}

type BookingSearchType int

const (
	Dates BookingSearchType = iota
	References
)

func (dt BookingSearchType) HotelApiDescription() string {
	switch dt {
	case Dates:
		return "Dates"
	default:
		return "References"
	}
}

type Rate int

const (
	B2b Rate = iota
	B2c
)

type MarkupRuleType string

const (
	MarkupRuleTypePercent MarkupRuleType = "PERCENT"
	MarkupRuleTypeImport  MarkupRuleType = "IMPORT"
)

func (e MarkupRuleType) IsValid() bool {
	switch e {
	case MarkupRuleTypePercent, MarkupRuleTypeImport:
		return true
	}
	return false
}

func (e MarkupRuleType) String() string {
	return string(e)
}

func (e *MarkupRuleType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MarkupRuleType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MarkupRuleType", str)
	}
	return nil
}

func (e MarkupRuleType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
