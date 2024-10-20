package main

import (
	"testing"
)

func Test_checkIfCanBreak(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test-1", args{"abe", "acd"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfCanBreak(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkIfCanBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkIfCanBreak2(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test-1", args{"abe", "acd"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfCanBreak2(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkIfCanBreak2() = %v, want %v", got, tt.want)
			}
		})
	}
}
