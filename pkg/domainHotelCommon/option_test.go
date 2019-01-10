package domainHotelCommon

import (
	"testing"
)

func TestOptionsSortByNetPrice(t *testing.T) {
	// ordered floats
	nets := []float64{-1, 0, 10, 10.2, 20}
	// options with unordered (done arbitrary) nets
	options := Options{
		&Option{Price: Price{Net: nets[3]}},
		&Option{Price: Price{Net: nets[1]}},
		&Option{Price: Price{Net: nets[4]}},
		&Option{Price: Price{Net: nets[0]}},
		&Option{Price: Price{Net: nets[2]}},
	}
	// sort and check that nets are ordered
	options.SortByNetPrice()
	for i, o := range options {
		if o.Price.Net != nets[i] {
			t.Fatalf("wrong sorted options, for index '%v' found options with net '%v', expecting '%v'", i, o.Price.Net, nets[i])
		}
	}
}

func TestOptionsGroupBy(t *testing.T) {
	options := Options{
		&Option{HotelCode: "0"},
		&Option{HotelCode: "1"},
		&Option{HotelCode: "2"},
		&Option{HotelCode: "1"},
		&Option{HotelCode: "0"},
		&Option{HotelCode: "1"},
		&Option{HotelCode: "1"},
		&Option{HotelCode: "0"},
	}

	// group by hotel code appending every option
	og := BasicOptionGrouper{
		IndexFunc: OptionIndexHotelCode,
		GroupFunc: OptionAppend,
	}

	idxOpts := options.GroupBy(og)
	if len(idxOpts) != 3 {
		t.Fatalf("idxOpts len should be 3, found: %v", len(idxOpts))
	}
	if len(idxOpts["0"]) != 3 {
		t.Errorf("options for hotel '0' len should be 3, found: %v", len(idxOpts["0"]))
	}
	if len(idxOpts["1"]) != 4 {
		t.Errorf("options for hotel '1' len should be 4, found: %v", len(idxOpts["1"]))
	}
	if len(idxOpts["2"]) != 1 {
		t.Errorf("options for hotel '2' len should be 1, found: %v", len(idxOpts["2"]))
	}
}
