package cmd

import (
	"strings"
	"bytes"
	"reflect"
	"testing"
)

func TestNewIfacesCommand(t *testing.T) {
	tests := []struct {
		name    string
		wantUse string
		wantShort string
	}{
		{"Test Descriptions", "ifaces", "Lists all network interfaces that exist in this machine"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			got := NewIfacesCommand(out)
			if !reflect.DeepEqual(got.Use, tt.wantUse) {
				t.Errorf("NewIfacesCommand() = %v, want %v", got.Use, tt.wantUse)
			}
			if !reflect.DeepEqual(got.Short, tt.wantShort) {
				t.Errorf("NewIfacesCommand() = %v, want %v", got.Short, tt.wantShort)
			}
		})
	}
}

func Test_findNetworkInterfaces(t *testing.T) {
	tests := []struct {
		name    string
		wantOut string
	}{
		{"List of Interfaces contains lo0", "lo0\n"},
		{"List of Interfaces contains en0", "en0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			findNetworkInterfaces(out)
			gotOut := out.String()
			if !strings.Contains(gotOut, tt.wantOut) {
				t.Errorf("findNetworkInterfaces() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
