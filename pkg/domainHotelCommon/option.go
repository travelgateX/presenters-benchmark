package domainHotelCommon

import (
	"presenters-benchmark/pkg/access"
	"sort"
)

type Option struct {
	Supplier          string        `json:"supplierCode" gqlgen:"supplierCode"`
	Access            string        `json:"accessCode" gqlgen:"accessCode"`
	Market            string        `json:"market"`
	HotelCode         string        `json:"hotelCode"`
	HotelName         *string       `json:"hotelName"`
	BoardCode         *string       `json:"boardCode"`
	BoardCodeOriginal string        `gqlgen:"boardCodeSupplier"`
	PaymentType       PaymentType   `json:"-"`
	Status            StatusType    `json:"status"`
	Occupancies       []Occupancy   `json:"occupancies"`
	Rooms             []Room        `json:"rooms"`
	Price             Price         `json:"price"`
	Supplements       []*Supplement `json:"supplements"`

	Surcharges   []Surcharge            `json:"surcharges"`
	RateRules    []access.RateRulesType `json:"rateRules"`
	CancelPolicy *CancelPolicy          `json:"cancelPolicy"`
	Remarks      *string                `json:"remarks"`
	OptionID     string                 `json:"id" gqlgen:"id"`
	Id           OptionID               `json:"-"`
	Token        string                 `json:"token"`
	Context      *string                `json:"-"`
	Group        *string                `json:"-"`
	RqDeepLink   *string                `json:"-"`
	Criteria     string                 `json:"-"`
}

func (*Option) IsBookableOptionSearch() {}

func (o *Option) HotelCodeSupplier() string {
	return o.Id.HotelCode
}

func (o *Option) Prices() []*Price {
	prices := make([]*Price, 0, 10)

	// Price
	prices = append(prices, &o.Price)

	// Room Price
	for ir := range o.Rooms {
		prices = append(prices, &o.Rooms[ir].RoomPrice.Price)
		if len(o.Rooms[ir].RoomPrice.Breakdown) > 0 {
			for ibd := range o.Rooms[ir].RoomPrice.Breakdown {
				// Breakdown Price
				prices = append(prices, &o.Rooms[ir].RoomPrice.Breakdown[ibd].Price)
			}
		}
	}

	// Supplements
	if len(o.Supplements) > 0 {
		for isup := range o.Supplements {
			if o.Supplements[isup].Price != nil {
				prices = append(prices, o.Supplements[isup].Price)
			}
		}
	}

	return prices
}

func (o *Option) PriceOption() *Price {
	return &o.Price
}

// Options is a set of options with helper methods
type Options []*Option

// SortBy sorts the provided slice given the provided less function.
func (opts Options) SortBy(less func(i, j int) bool) {
	sort.Slice(opts, less)
}

// SortByNetPrice will sort options from the lowest net price to the highest
func (opts Options) SortByNetPrice() {
	sort.Slice(opts, func(i, j int) bool {
		return opts[i].Price.Net < opts[j].Price.Net
	})
}

// IndexedOptions is group of indexed sets of options
type IndexedOptions map[string]Options

func (r IndexedOptions) ToOptions() Options {
	ret := make(Options, 0, len(r))
	for _, options := range r {
		for _, option := range options {
			ret = append(ret, option)
		}
	}
	return ret
}

// OptionGrouper permits to make option groups given different patterns
type OptionGrouper interface {
	// Index builds an option's index, two options with the same Index will be grouped
	Index(*Option) string
	// Group takes a set of options already grouped, a new option that belongs to this group, and a group result
	Group(string, Options, *Option) Options
	// GroupLen is the len used when instancing new groups
	GroupLen() int
}

// GroupBy takes 'Options' and indexes them by an 'OptionGrouper'
func (opts Options) GroupBy(og OptionGrouper) IndexedOptions {
	idxOpts := make(IndexedOptions)
	for _, opt := range opts {
		idx := og.Index(opt)
		if group, ok := idxOpts[idx]; ok {
			idxOpts[idx] = og.Group(idx, group, opt)
		} else {
			group = make([]*Option, 0, og.GroupLen())
			idxOpts[idx] = og.Group(idx, group, opt)
		}
	}
	return idxOpts
}

// BasicOptionGrouper is a straightforward implementation of OptionGrouper that can be instanced with
// index and group functions, this package provides some of them
type BasicOptionGrouper struct {
	// IndexFunc builds an index from an Option
	IndexFunc OptionIndexFunc
	// OptionGroupFunc takes a set of options and an option that results in another set of options
	GroupFunc     OptionGroupFunc
	GroupLenValue int
}

func (og BasicOptionGrouper) Index(opt *Option) string { return og.IndexFunc(opt) }
func (og BasicOptionGrouper) Group(idx string, opts Options, opt *Option) Options {
	return og.GroupFunc(idx, opts, opt)
}
func (og BasicOptionGrouper) GroupLen() int { return og.GroupLenValue }

// OptionIndexFunc builds an index from an Option
type OptionIndexFunc func(*Option) string

// OptionGroupFunc takes a set of options and an option that results in another set of options
type OptionGroupFunc func(string, Options, *Option) Options

// OptionAppend appends the option to the set
func OptionAppend(opts Options, opt *Option) Options {
	return append(opts, opt)
}

// OptionLimitedAppend appends the option to the set
func OptionLimitedAppend(groupLimit uint) OptionGroupFunc {
	gl := int(groupLimit)
	return func(idx string, opts Options, opt *Option) Options {
		if len(opts) >= gl {
			return opts
		}
		return append(opts, opt)
	}
}
