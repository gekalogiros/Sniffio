package utilities

import (
	"reflect"
	"testing"
)

func TestLookupVendor(t *testing.T) {
	macLookup := NewMacLookup()
	macLookup.Load()
	tests := []struct {
		name string
		mac string
		want string
	}{
		{"Should return vendor for mac address", "00:22:72:DC:DA:A8", "American Micro-Fuel Device Corp."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := macLookup.LookupVendor(tt.mac); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestLookupVendor() = %v, want %v", got, tt.want)
			}
		})
	}
}
