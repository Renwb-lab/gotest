package main

import "testing"

func Test_countSpecialNumbers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"20", args{20}, 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSpecialNumbers(tt.args.n); got != tt.want {
				t.Errorf("countSpecialNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
