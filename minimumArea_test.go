package main

import "testing"

func Test_minimumArea(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test-1", args{[][]int{{0, 1, 0}, {1, 0, 1}}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumArea(tt.args.grid); got != tt.want {
				t.Errorf("minimumArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
