package utilities

import (
	"testing"
)

func TestMac(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want string
	}{
		{"Should convert a byte array to a mac address containing hex bytes separated by colons", []byte{6, 13, 16, 220, 218, 168}, "6:D:10:DC:DA:A8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mac(tt.data); got != tt.want {
				t.Errorf("Mac() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIP(t *testing.T) {
	tests := []struct {
		name string
		ip []byte
		want string
	}{
		{"Should convert a byte array to an ip address containing decimals separatd by dots", []byte{192,168,0,14}, "192.168.0.14"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IP(tt.ip); got != tt.want {
				t.Errorf("IP() = %v, want %v", got, tt.want)
			}
		})
	}
}
