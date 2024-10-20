package main

import "testing"

func Test_twoEggDrop(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test-1", args{100}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoEggDrop(tt.args.n); got != tt.want {
				t.Errorf("twoEggDrop() = %v, want %v", got, tt.want)
			}
		})
	}
}
