package cmd

import (
	"bytes"
	"reflect"
	"testing"
)

func TestNewVersionCommand(t *testing.T) {
	tests := []struct {
		name      string
		wantUse   string
		wantShort string
	}{
		{"Test Command and description", "version", "Print the version number of Sniffio"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewVersionCommand(&bytes.Buffer{})
			if !reflect.DeepEqual(got.Use, tt.wantUse) {
				t.Errorf("NewVersionCommand() = %v, want %v", got.Use, tt.wantUse)
			}
			if !reflect.DeepEqual(got.Short, tt.wantShort) {
				t.Errorf("NewVersionCommand() = %v, want %v", got.Short, tt.wantShort)
			}
		})
	}
}
