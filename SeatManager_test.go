package main

import (
	"testing"
)

func TestSeatManager_Reserve(t *testing.T) {
	type fields struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"test-1", fields{5}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructorer(tt.fields.n)
			if got := this.Reserve(); got != tt.want {
				t.Errorf("SeatManager.Reserve() = %v, want %v", got, tt.want)
			}
		})
	}
}
