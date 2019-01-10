package gophers

import (
	"bytes"
	"encoding/json"
	"github.com/graph-gophers/graphql-go"
	hagraphql "hub-aggregator/common/graphql"
	"hub-aggregator/common/kit/routing"
	"hub-aggregator/common/stats"
	"net/http"
	"rfc/presenters/pkg/domainHotelCommon"
	"rfc/presenters/pkg/presenter"
	"rfc/presenters/pkg/presenter/gophers/resolver"
	"strconv"
)

type Candidate struct{}

var _ presenter.CandidateServer = (*Candidate)(nil)
var _ presenter.CandidateHandlerFunc = (*Candidate)(nil)

func (Candidate) NewServer(addr, pattern string, options []*presenter.Option, results chan<- presenter.OperationResult) (*routing.Server, error) {
	soptions := make([]*domainHotelCommon.Option, len(options))
	for i, o := range options {
		opt := (domainHotelCommon.Option)(*o)
		soptions[i] = &opt
	}

	schema, err := graphql.ParseSchema(
		schema,
		&graphResolver.QueryResolver{soptions},
		&graphResolver.MutationResolver{},
		graphql.Logger(hagraphql.Logger{}),
	)

	if err != nil {
		return nil, err
	}

	return presenter.NewGzipCandidateServer(
		addr,
		pattern,
		HandlerFunc(schema, results),
	), nil
}

func (Candidate) HandlerFunc(options []*presenter.Option) (http.HandlerFunc, error) {
	soptions := make([]*domainHotelCommon.Option, len(options))
	for i, o := range options {
		opt := (domainHotelCommon.Option)(*o)
		soptions[i] = &opt
	}

	schema, err := graphql.ParseSchema(
		schema,
		&graphResolver.QueryResolver{soptions},
		&graphResolver.MutationResolver{},
		graphql.Logger(hagraphql.Logger{}),
	)

	if err != nil {
		return nil, err
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var params struct {
			Query         string                 `json:"query"`
			OperationName string                 `json:"operationName"`
			Variables     map[string]interface{} `json:"variables"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := schema.Exec(r.Context(), params.Query, params.OperationName, params.Variables)
		// schema exec errors
		if len(response.Errors) > 0 {
			buf := bytes.Buffer{}
			for i, qErr := range response.Errors {
				if i > 0 {
					buf.WriteString("\n")
				}
				buf.WriteString("[" + strconv.Itoa(i) + "]: " + qErr.Error())
			}
			errStr := "schema exec errors:\n " + buf.String()
			http.Error(w, errStr, http.StatusBadRequest)
			return
		}

		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}, nil
}

func HandlerFunc(schema *graphql.Schema, results chan<- presenter.OperationResult) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		times := presenter.OperationResult{}
		totalInit := stats.UtcNow()
		var params struct {
			Query         string                 `json:"query"`
			OperationName string                 `json:"operationName"`
			Variables     map[string]interface{} `json:"variables"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := schema.Exec(r.Context(), params.Query, params.OperationName, params.Variables)
		// schema exec errors
		if len(response.Errors) > 0 {
			buf := bytes.Buffer{}
			for i, qErr := range response.Errors {
				if i > 0 {
					buf.WriteString("\n")
				}
				buf.WriteString("[" + strconv.Itoa(i) + "]: " + qErr.Error())
			}
			errStr := "schema exec errors:\n " + buf.String()
			http.Error(w, errStr, http.StatusBadRequest)
			return
		}

		SerializeInit := stats.UtcNow()
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		times.SerializeTime = stats.NewTimes(SerializeInit)
		times.TotalTime = stats.NewTimes(totalInit)
		results <- times
		return
	}
}

var schema = `schema {
  query: Query
}

type Query{
  hotelX: HotelXQuery
}

type HotelXQuery{
  # Available options of an hotel for a given date and itinerary. It does not filter different classes, times or
  # fares. It will always retrieve all results returned by the suppliers. The availability request is very straight
  # forward. It only requires the criteria of search (destination, travel dates and the number of pax in each room).
  # But you must preload the other fields in our system by complete the fields absents.
  search: HotelSearch

  # Returns status of the search service.
  searchStatusService: ServiceStatus!
}
# Business rules type
enum BusinessRulesType {
  # The cheapest options is returned without exceeding the optionsQuota limit.
  CHEAPER_AMOUNT
  
  # Groups the option by room type without exceeding the optionsQuota limit.
  ROOM_TYPE
}

# Options type
enum CancelPenaltyType {
  # Indicates the number of nights to be penalized.
  NIGHTS
  
  # Indicates the percentage to pay based on the option price.
  PERCENT
  
  # Indicates the exact amount payable.
  IMPORT
}

# Charge Type
enum ChargeType {
  # The charge is included.
  INCLUDE
  
  # The charge is excluded.
  EXCLUDE
}

# Duration Type
enum DurationType {
  # Date range is set.
  RANGE
  
  # Not restricted by date.
  OPEN
}

# Indicates what type of value is the markup, by percentage or is an import.
enum MarkupRuleType {
  # Indicates the percentage applied by a rule.
  PERCENT
  
  # Indicates the exact amount applied by a rule.
  IMPORT
}

# Plugin Step Type. https://docs.travelgatex.com/hotelx/plugins/overview/
enum PluginStepType {
  # Plugins executed after Buyer requests message to HotelX
  REQUEST
  
  # Plugins executed before sending request to Supplier using Access and after Accesses have been calculated
  REQUEST_ACCESS
  
  # Plugins executed after Supplier responds message. For every option returned
  RESPONSE_OPTION
  
  # Plugins executed after all Access options has been responded
  RESPONSE_ACCESS
  
  # Plugins executed before HotelX responds message to to Buyer
  RESPONSE
}

# Plugin Type. https://docs.travelgatex.com/hotelx/plugins/overview/
enum PluginType {
  # PRE_STEP is the first plugin that a step will execute, allows a full range of operations:
  # split arrays, join arrays, modify object values, add or remove object instances
  PRE_STEP
  
  # HOTEL_MAP allows to match Seller and Buyer hotel codes based on contexts
  HOTEL_MAP
  
  # BOARD_MAP allows to match Seller and Buyer board codes based on contexts
  BOARD_MAP
  
  # ROOM_MAP allows to match Seller and Buyer room codes based on contexts
  ROOM_MAP
  
  # CURRENCY_CONVERSION allows to match Seller and Buyer hotel codes based on contexts
  CURRENCY_CONVERSION
  
  # MARKUP allows to apply markup over price
  MARKUP
  
  # AGGREGATION allows to aggregate multiple supplier options
  AGGREGATION
  
  # POST_STEP is the last plugin that a step will execute, allows a full range of operations:
  # split arrays, join arrays, modify object values, add or remove object instances
  POST_STEP
}

# Price Type
enum PriceType {
  # Price without deductions.
  GROSS
  
  # Price after deducting all discounts and rebates.
  NET
  
  # Final quantity. Sum of multiple quantities.
  AMOUNT
}

# Service Type
enum ServiceType {
  # A ticket or pass authorizing the holder to ski in a certain place or resort. Gross.
  SKI_PASS
}

# Indicartes options status
enum StatusType {
  # The status of the avail is available
  OK
  
  # The status of the avail is On request
  RQ
}

# Supplement Type
enum SupplementType {
  # A ticket or pass authorizing the holder to ski in a certain place or resort.
  SKI_PASS
  
  # Lessons of any type that the costumer can take.
  LESSONS
  
  # Supplement of a determined meal plan.
  MEALS
  
  # Extra equipment for a specific purpose.
  EQUIPMENT
  
  # Admission to some service.
  TICKET
  
  # Transfers used by the costumer.
  TRANSFERS
  
  # Gala: A festive occasion, celebration or special entertainment.
  GALA
  
  # Activities that the costumer can do.
  ACTIVITY
}

# Unit Time Type
enum UnitTimeType {
  # Day
  DAY
  
  # Hour
  HOUR
}

# Options payment type
enum PaymentType {
  # The payment is managed by the supplier.
  MERCHANT
  
  # The payment is made straight to the actual payee, without sending it through an intermediary or a third party.
  DIRECT
  
  # The payment is managed by the supplier. The payment is effectuated at the time of booking.
  CARD_BOOKING
  
  # The payment is managed by the supplier. The payment is effectuated at check in in the hotel.
  CARD_CHECK_IN
}

# Rate Rules
enum RateRulesType {
  # The product can't be sold separately from another product attached to it, such as a flight.
  PACKAGE
  
  # Options that can only be sold to people who are 55 and older.
  OLDER55
  
  # Options that can only be sold to people who are 60 and older.
  OLDER60
  
  # Options that can only be sold to people who are 65 and older.
  OLDER65
  
  # The rate CanaryResident is applicable to Canary Islands residents only.
  CANARY_RESIDENT
  
  # The rate BalearicResident is applicable to Balearic Islands residents only.
  BALEARIC_RESIDENT
  
  # The rate largeFamily is applied to large families and is determined by each supplier
  LARGE_FAMILY
  
  # The rate honeymoon is applied to those who just got married and is determined by each supplier.
  HONEYMOON
  
  # The rate publicServant is applicable to public servants only.
  PUBLIC_SERVANT
  
  # The rate unemployed is applied to those without work.
  UNEMPLOYED

  #The rate normal refers to options without RateRule
  NORMAL

  #The rate non refundable is applied to non refundable options
  NON_REFUNDABLE
}

# Include *OR* exclude accesses in this specific search query. If not specified, default accesses will be used.
# Only one list (includes or excludes) *MUST* be used.
input AccessFilterInput {
  # These Access IDs will overwrite the default configuration. Only the IDs on this list will be used in the search query.
  includes: [ID!]
  
  # These Access IDs will overwrite the default configuration. The IDs on this list will be excluded from the search query.
  excludes: [ID!]
}

# List of business rules to use as filter on the options.
input BusinessRulesInput {
  # Options quota per search. Maximum numbers of options to be returned by the search query.
  optionsQuota: Int
  
  # Different business rules to filter the returned options.
  businessRulesType: BusinessRulesType
}

# The information and credentials required to access the supplier’s system.
input ConfigurationInput {
  # User name for the connection.
  username: String
  
  # Password for the connection
  password: String
  
  # URL or endpoint for the connection.
  urls:           UrlsInput!
  
  # List of parameters with additional required information.
  parameters: [ParameterInput!]
  
  # Source Markets allowed for the Access
  markets: [String!]
  
  # RateRules allowed for the access.
  rateRules: [RateRulesType!]
}

input HotelXFilterSearchInput{
  # Filter that selects the filter criteria which will be used in this availability. Currently you can only choose the accesses.
  # You must choose one of them, include or exclude, or the other alternative isn't specified anything.
  # If input both, you will receive a validation error that indicates this error.

  # You can specify one of the filters or any of them. In this latter case, all the configurated accesses will be executed.
  access: AccessFilterInput
}
# Filter that selects the filter criteria which will be used in this availability. Currently you can only choose the accesses.
# You must choose one of them, include or exclude, or the other alternative isn't specified anything.
# If input both, you will receive a validation error that indicates this error.
#@deprecated(reason: "deprecated from 2018-08-20. Please, use filterSearch")
input FilterInput {
  # You can specify one of the filters or any of them. In this latter case, all the configurated accesses will be executed.
  access: AccessFilterInput

}

# Search criteria contains destination, travel dates and the number of pax in each room.
# You must preload the other fields in our system by complete the fields absents.
input HotelCriteriaSearchInput {
  # Check-in date for booking
  # Format: YYYY-MM-DD
  checkIn: Date!
  
  # Check-out, booking date
  # Format: YYYY-MM-DD
  checkOut: Date!
  
  # Hotel Codes.
  hotels: [String!]

  # Destination codes.
  destinations: [String!]
  
  # For multi-room bookings, this array will contain multiple elements (rooms).
  # For each room you have to specify its own occupancy.
  occupancies: [RoomInput!]!
  
  # Language to be used in request
  language : Language
  
  # Currency requested if supported by supplier
  currency : Currency
  
  # Nationality of the guest (use ISO3166_1_alfa_2)
  nationality : Country
  
  # Targeted zone, country or point-ofsale-to be used in request.
  market : String
}

# Settings that you can edit for this avail. Values are loaded by default in our Back Office.
input HotelSettingsInput {
  # Indicates the context of the I/O codes (hotel, board, room and rates)
  context: String
  
  #Indicates if you want use context, or not, by default is true.
  #@deprecated(reason: "deprecated from 2017-12-12. Redundant.")
  useContext: Boolean
  
  # This field is occurs only if the authorization header is of the type JWT.. It is used to change the user that has been set by default in the preload.
  #@deprecated(reason: "deprecated from 2018-03-19. Redundant.")
  connectUser: String
  
  # Client name, this field is occurs only if the authorization header is of the type JWT.. It is used to change the user that has been set by default in the preload.
  client: ID
  
  # Group whose resources want to be used
  group: ID
  
  # Milliseconds before the connection is closed.
  timeout: Int
  
  # Returns all the transactions exchanged with the supplier.
  auditTransactions: Boolean

  # This flag allows only the accesses checked as test. By default is production.
  testMode: Boolean

  # Used to identify the origin of the request, this is only used in plugins. 
  clientTokens: [String!] 
}

# AccessInput overwrites an existent access in our Back Office or creates a new
# one to be used in this search query only. An access object contains its own code, configuration and settings.
input HotelXAccessInput {
  # The accessID used to identify the existing access in our Back Office in order to
  # overwrite it. Acts as an identifier in this search. It can either exist or not.
  accessId: ID!
  
  # Information required to access the supplier's system.
  configuration: ConfigurationInput
  
  # You can configure an special settings for any access. This level overwrites the search and supplier settings levels.
  settings: SettingsBaseInput
}

# Supplier object. Contains its own settings, code and access.
input HotelXSupplierInput {
  # You can configure an special settings for any supplier. This level overwrites the avail settings level but not the
  # access settings level.
  settings: SettingsBaseInput
  
  # Code that represents a supplier in our system.
  # This information is mandatory.
  code: String!
  
  # Array of accesses that can overwrite an existing access information or include a new access for this avail.
  accesses: [HotelXAccessInput!]
}

# Pax object that contains the pax age.
input PaxInput {
  # Pax age.
  age: Int!
}

# If requested, only options with the specified rateRules will be returned
input RateRulesFilterInput {
  # if includes not nil: only options without rate rules and options with rate rules found in includes will be returned
  includes: [RateRulesType!]
  
  # if excludes not nil: only options without rate rules and options with rate rules that haven't been sent in excludes will be returned
  excludes: [RateRulesType!]
}

# Occupancy for a room. It contains a list of pax ages.
input RoomInput {
  # Array of pax ages. The number of items in the array will indicate the pax occupancy.
  paxes: [PaxInput!]!
}

# Plugin to execute.
input PluginsInput {
  # type of the plugins to execute
  type: PluginType!
  
  # name of plugin to execute
  name: String!
  
  # Plugin's parameters
  parameters: [ParameterInput!]
}

# Plugin to execute.
input PluginStepInput {
  # Indicates where the plugin will be executed.
  step: PluginStepType!
  
  # Indicates the plugin that will be executed.
  pluginsType: [PluginsInput!]
}

# Contains the time out and business rules of a supplier or an access.
input SettingsBaseInput {
  # Milliseconds before the connection is closed.
  timeout: Int
  
  # Specifies if transactions exchanged with the supplier have to be logged or not.
  auditTransactions: Boolean

  # The currency
  currency: Currency
}

input HotelXPluginFilterInput{
  # Plugins to include (only these plugins will be executed)
  includes: [HotelXFilterPluginTypeInput!]
  # Plugins to exclude
  excludes: [HotelXFilterPluginTypeInput!]
}
input HotelXFilterPluginTypeInput{
  # The Step of the plugin to filter
  step: PluginStepType!
  # The Type of the plugin to filter
  type: String!
  # The Name of the plugin to filter
  name: String!
}
# Parameters Input.
input ParameterInput {  
  # Contains the keyword/Id to identify a parameter.
  # This information is mandatory.
  key: String!
  # Contains the parameter values.
  # This information is mandatory.
  value: String!
}

# URLs Input
input UrlsInput {
  # Specific URL for Availability method.
  search:         URI
  # Specific URL for Reservation method.
  quote:          URI
  # Specific URL for Valuation method.
  book:           URI
  # Supplier URL used for multiple methods.
  generic:        URI
}

interface BookableOptionSearch {
  # Supplier that offers this option.
  supplierCode: String!
  
  # Access code of this option.
  accessCode: String!
  
  # Indicates the id to be used on Quote as key
  id: String!
}

interface Priceable {
  # Specifies the currency.
  currency: Currency!
  
  # Is binding.
  binding: Boolean!
  
  # Specifies the import net.
  net: Float!
  
  # Specifies the import gross.
  gross: Float
  
  # Specifies the exchange.
  exchange: Exchange!
}

interface Response {
  # Application stats
  stats(token: String!): StatsRequest
  
  # Data sent and received in the supplier’s original format.
  auditData: AuditData
  
  # Errors that lead the service to stop
  errors: [Error!]
  
  # Potentially harmful situations or errors that do not stop the service
  warnings: [Warning!]
}

# Additional information about the option
type AddOns {
  # Extra information from the distribution layer
  distribute: JSON @deprecated(reason: "deprecated from 2018-05-21. You can find it in distribution AddOn")
  
  # Extra information from the distribution layer
  distribution: [AddOn!]
}

# Additional information about the option
type AddOn {
  # Contains keyword/ID to identify the AddOn.
  key: String!
  
  # Contains AddOn values.
  value: JSON!
}

# Data sent and received in the supplier’s native format.
type AuditData {
  # List of transactions data
  transactions:    [Transactions!]!
  
  # TimeStamp
  timeStamp:       DateTime!
  
  # Process time in milliseconds (ms)
  processTime:     Float!
}

# Contains information about a bed.
type Bed {
  # Specifies the bed type
  type: String
  
  # Description about the bed
  description: String
  
  # Indicates number of beds in a room
  count: Int
  
  # Specifies if the bed is shared or not
  shared: Boolean
}

# Contains information for cancellation penalities..
type CancelPenalty {
  # Cancellation fees applicable X number of hours before the check-in date
  hoursBefore: Int!
  
  # Type of penalty; this can be Nights, Percent or Import
  penaltyType: CancelPenaltyType!
  
  # Currency used in the cancellation policy
  currency: Currency!
  
  # Value of the cancellation policy
  value: Float!
}

# Information about a policy cancellation.
type CancelPolicy {
  # Indicates if the option is refundable or non-refundable
  refundable: Boolean!
  
  # List of cancellation penalties
  cancelPenalties: [CancelPenalty!]
}

# Search criteria contains destination, travel dates and the number of pax in each room.
type CriteriaSearch {
  # Check-in date for booking
  # Format: YYYY-MM-DD
  checkIn: Date!
  
  # Check-out, booking date
  # Format: YYYY-MM-DD
  checkOut: Date!
  
  # Contains the list of hotels's ID
  hotels: [String!]!
  
  # For multi-room bookings, this array will contain multiple elements (rooms).
  # For each room you have to specify its own occupancy.
  occupancies: [RoomCriteria!]!
  
  # Language to be used in request
  language : Language
  
  # Currency requested if supported by supplier
  currency : Currency
  
  # Nationality of the guest (use ISO3166_1_alfa_2)
  nationality : Country
  
  # Targeted zone, country or point-ofsale-to be used in request.
  market : String!
}

# Provides information about the currency of original, and its rate applied over the results returned by the Supplier.
type Exchange {
  # Provide information about the currency of origin
  currency: Currency!
  
  # Provides information about the rate applied over results
  rate: Float!
}

# An option includes hotel information, meal plan, total price, conditions and room description
type HotelOptionSearch implements BookableOptionSearch {
  # Supplier that offers this option.
  supplierCode: String!
  
  # Access code of this option.
  accessCode: String!
  
  # Market of this option.
  market: String!
  
  # Code of the hotel in the context selected.
  hotelCode: String!
  
  # Supplier's hotel code.
  hotelCodeSupplier: String!
  
  # Name of the hotel.
  hotelName: String
  
  # Code of the board in the context selected.
  boardCode: String!
  
  # Supplier's board code.
  boardCodeSupplier: String!
  
  # Indicates the payment type of the option returned. Possible options: Merchant, Direct, Card Booking, Card check in and Mixed.
  paymentType: PaymentType!
  
  # The possible values in status in response are Available (OK) or On Request (RQ).
  status: StatusType!
  
  # List of occupancies for the request
  occupancies: [Occupancy!]!
  
  # List of rooms of the option returned.
  rooms: [Room!]!
  
  # Specifies the prices (Gross, Net and Amount) of the option returned.
  price: Price!
  
  # List of supplements of the option returned.
  supplements: [Supplement!]
  
  # List of surcharges of the option returned.
  surcharges: [Surcharge!]
  
  # Specifies rate rules of the option returned.
  rateRules: [RateRulesType!]
  
  # Specifies cancel policies of the option returned.
  cancelPolicy: CancelPolicy
  
  # Additional information about the option.
  remarks: String
  
  # Additional information about the option
  addOns: AddOns
  
  # Token for Deep Link
  token: String!
  
  # Indicates the quote key
  id: String!
}

# Results from Avail Hotel; contains all the available options for a given date and itinerary
type HotelSearch implements Response {
  # Indicates the context of the response.
  context: String
  
  # Application stats in string format
  stats(token: String!): StatsRequest
  
  # Data sent and received in the supplier's native format.
  auditData: AuditData
  
  # Request Criteria
  requestCriteria: CriteriaSearch
  
  # List of options returned according to the request.
  options: [HotelOptionSearch!]
  
  # Errors that abort services
  errors: [Error!]
  
  # Potentially harmful situations or errors that won't force the service to abort
  warnings: [Warning!]
}

# Informs markup applied over supplier price.
type Markup implements Priceable {
  #channel of markup application.
  channel: String
  
  # Currency code indicating which currency should be paid.
  # This information is mandatory.
  currency: Currency!
  
  # It indicates if the price indicated in the gross must be respected.
  # That is, the customer can not sell the room / option at a price lower than that established by the supplier.
  # This information is mandatory.
  binding: Boolean!
  
  # Indicates the net price that the customer must pay to the supplier plus the markup.
  # This information is mandatory.
  net: Float!
  
  # Indicates the retail price that the supplier sells to the customer plus the markup.
  gross: Float
  
  # Informs about the currency of origin, and the rate applied over result.
  # This information is mandatory.
  exchange: Exchange!
  
  # Breakdown of the applied rules for a markup
  rules: [Rule!]!
}

# Information about occupancy.
type Occupancy {
  # Unique ID room in this option.
  id: Int!
  
  # List of pax of this occupancy.
  paxes: [Pax!]!
}

# Specifies the age pax. The range of what is considered an adult, infant or baby is particular to each supplier.
type Pax {
  # Specifies the age pax.
  age: Int!
}

# Price indicates the value of the room/option.
# Supplements and/or surcharges can be included into the price, and will be verified with nodes Supplements/Surcharges.
type Price implements Priceable {
  # Currency code indicating which currency should be paid.
  # This information is mandatory.
  currency: Currency!
  
  # It indicates if the price indicated in the gross must be respected.
  # That is, the customer can not sell the room / option at a price lower than that established by the supplier.
  # This information is mandatory.
  binding: Boolean!
  
  # Indicates the net price that the customer must pay to the supplier.
  # This information is mandatory.
  net: Float!
  
  # Indicates the retail price that the supplier sells to the customer.
  gross: Float
  
  # Provides information about the currency of original, and its rate applied over the results returned by the Supplier.
  # This information is mandatory.
  exchange: Exchange!
  
  # Informs markup applied over supplier price.
  markups: [Markup!]
}

# Contains internal information.
type StatsRequest {
  # Total transaction time
  total:          Stat!
  
  # Request validation time
  validation:     Stat!
  
  # Process time. Contains communication time, parse time and plugin time.
  process:        Stat!
  
  # Build access time
  configuration:	Stat!
  
  # Request time
  request:				Stat!
  
  # Response time
  response:       Stat!
  
  # Plugin execution time
  requestPlugin:  StatPlugin
  
  # Plugin execution time
  responsePlugin: StatPlugin
  
  # Number of hotels
  hotels:         Int!
  
  # Number of zones
  zones:          Int!
  
  # Number of cities
  cities:         Int!
  
  # Docker Id
  dockerID:       String!
  
  # Detail access time
  Accesses:       [StatAccess!]!
}

# Information about daily price.
type PriceBreakdown {
  # Start date in which the price becomes effective.
  effectiveDate: Date!
  
  # Expire date of price.
  expireDate: Date!
  
  # Specifies the daily price.
  price: Price!
}

# Information about room promotions(offers).
type Promotion {
  # Specifies the promotion code.
  code: String!
  
  # Specifies the promotion name.
  name: String
  
  # Promotion effective date.
  effectiveDate: Date
  
  # Promotion expire date.
  expireDate: Date
}

# Information about the rate of the option returned.
type RatePlan {
  # Specifies the rate code.
  code: String!
  
  # Specifies the rate name.
  name: String
  
  # Start date in which the rate becomes effective.
  effectiveDate: Date
  
  # Expire date of the rate.
  expireDate: Date
}

# Contains information about the Resort.
type Resort {
  # Specifies the resort code.
  code: String!
  
  # Specifies the resort name.
  name: String
  
  # Specifies the resort description.
  description: String
}

# Contains the room information of the option returned.
type Room {
  # ID reference to the occupancy
  occupancyRefId: Int!
  
  # Indicates the room code
  code: String!
  
  # Description about the room
  description: String
  
  # Identifies if the room is refundable or not.
  refundable: Boolean
  
  # Number of rooms available with the same type.
  units: Int
  
  # Specifies the room price.
  roomPrice: RoomPrice!
  
  # List of beds.
  beds: [Bed!]
  
  # Daily break downs rate plan.
  ratePlans: [RatePlan!]
  
  # Daily break downs promotions.
  promotions: [Promotion!]
}

# Occupancy for a room. It contains a list of pax ages.
type RoomCriteria {
  # Array of pax ages. The number of items in the array will indicate the pax occupancy.
  paxes: [Pax!]!
}

# Specifies the room price.
type RoomPrice {
  # Total price for all days.
  price: Price!
  
  # Daily break downs price.
  breakdown: [PriceBreakdown!]
}

type Rule {
  # rule identifier
  id: String!
  
  # rule name
  name: String
  
  # type of the value
  type: MarkupRuleType!
  
  # value applied by this rule
  value: Float!
}

# Indicates the status of the service
type ServiceStatus{
  # Status code
  code: String
  
  # Status type
  type: String
  
  # Status description
  description : String
}

type Stat {
  # Start UTC
  start:      DateTime!
  
  # End UTC
  end:        DateTime!
  
  # Difference between start and end in miliseconds
  duration:   Float
}

type StatAccess {
  # Access name
  name:                   String!
  
  # Total access time
  total:                  Stat!
  
  # Static configuration time
  staticConfiguration:    Stat
  
  # Number of hotels
  hotels:                 Int!
  
  # Number of zones
  zones:                  Int!
  
  # Number of cities
  cities:                 Int!
  
  # Access request time
  requestAccess:          StatPlugin
  
  # Access response time
  responseAccess:         StatPlugin
  
  # Detail transaction time
  transactions:           [StatTransaction!]!
  
  # Plugin execution time
  plugins:                [StatPlugin!]
}

type StatTransaction {
  # Extra information about transaction.
  reference:              String!
  
  # Total transaction time
  total:                  Stat!
  
  # Build request time
  buildRequest:           Stat!
  
  # Worker connection time
  workerCommunication:    Stat!
  
  # Parse response time
  parseResponse:          Stat!
}

type StatPlugin{
  # Plugin name
  name:     String!
  
  # total plugin time
  total:    Stat!
}

# Supplement that it can be or its already added to the option returned. Contains all the information about the supplement.
type Supplement {
  # Specifies the supplement code.
  code: String!
  
  # Specifies the supplement name.
  name: String
  
  # Specifies the supplement description.
  description: String
  
  # Indicates the supplement type. Possible types: Fee, Ski_pass, Lessons, Meals, Equipment, Ticket, Transfers, Gla, Activity or Null.
  supplementType: SupplementType!
  
  # Indicates the charge types. We need to know whether the supplements have to be paid when the consumer gets to the hotel or beforehand.
  # Possible charge types: Include or Exclude.
  # when include: this supplement is mandatory and included in the option's price
  # when exclude: this supplement is not included in the option's price
  chargeType: ChargeType!
  
  # Indicates if the supplement is mandatory or not. If mandatory, this supplement will be applied to this option
  # if the chargeType is excluded the customer will have to pay it directly at the hotel
  mandatory: Boolean!
  
  # Specifies the duration type. Possible duration types: Range (specified dates) or Open. This field is mandatory for PDI.
  durationType: DurationType
  
  # Indicates the quantity of field in the element "unit".
  quantity: Int
  
  # Indicates the unit type. Possible unit types: Day or Hour.
  unit: UnitTimeType
  
  # Indicates the effective date of the supplement.
  effectiveDate: Date
  
  # Indicates the expire date of the supplement.
  expireDate: Date
  
  # Contains information about the resort
  resort: Resort
  
  # Indicates the supplement price.
  price: Price
}

# Surcharge that it can be or it is already added to the option returned. Contains all the information about the surcharge.
type Surcharge {
  # Indicates the charge types. We need to know whether the supplements have to be paid when the consumer gets to the hotel or beforehand.
  # Possible charge types: Include or Exclude.
  # when include: this surcharge is mandatory and included in the option's price
  # when exclude: this surcharge is not included in the option's price
  chargeType: ChargeType!
  
  # Indicates if the surcharge is mandatory or not. If mandatory, this surcharge will be applied to this option
  # if the chargeType is excluded the customer will have to pay it directly at the hotel
  mandatory: Boolean!
  
  # Indicates the surcharge price.
  price: Price!
  
  # Specifies the surcharge description.
  description: String
}

# Supplier transaction
type Transactions {
  # Transaction Request.
  request:        String!
  
  # Transaction Response.
  response:       String!
  
  # Time when the request has been processed.
  timeStamp:      DateTime!
}

# Application errors
type Error {
  # Error code
  code: String!

  # Error type
  type: String!

  # Error description
  description : String!
}

# Application warnings
type Warning {
  # Warning code
  code: String!

  # Warning type
  type: String!

  # Warning description
  description : String!
}

# The Country type represents Country values. A good example might be a Passenger Nationality.
# In queries or mutations, Country fields have to be specified in ISO 3166-1 alpha-2 format with enclosing double quotes "ES".
scalar Country

# The Currenty type represents Currency values. A good example might be a Rate Price Currency.
# In queries or mutations, Currency fields have to be specified in ISO 4217 format with enclosing double quotes "EUR".
scalar Currency

# The Date type represents Date values. A good example might be a Hotel CheckIn Date.
# In queries or mutations, DateTime fields have to be specified in ISO 8601 format with enclosing double quotes: "2017-10-22".
scalar Date

# The DateTime type represents DateTime values. A good example might be a transaction TimeSpan.
# In queries or mutations, DateTime fields have to be specified in ISO 8601 format with enclosing double quotes: "2017-10-22T13:57:31.123Z".
scalar DateTime

# The JSON type makes sure that it is actually valid JSON and returns the value as a parsed JSON object/array instead of a string.
# In queries or mutations, JSON fields have to be specified with enclosing double quotes. Special characters have to be escaped: "{\"int\": 1, \"string\": \"value\"}".
scalar JSON

# The Language type represents Language values. A good example might be a Hotel Description Language.
# In queries or mutations, Language fields have to be specified in ISO 3166-1 alpha-2 format with enclosing double quotes "es".
scalar Language

# The URI type represents a URI values. A good example mith be an Hotel Image URL.
# In queries or mutations, URI fields have to be specified in RFC 3986, RFC 3987, and RFC 6570 (level 4) compliant URI string format with enclosing double quotes: "http:\\www.travelgatex.com".
scalar URI

`
