package main

import "testing"

func Test_numBusesToDestination(t *testing.T) {
	type args struct {
		routes [][]int
		source int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal case 1", args{[][]int{{1, 2, 7}, {3, 6, 7}}, 1, 1}, 0},
		{"normal case 2", args{[][]int{{1, 2, 7}, {3, 6, 7}}, 1, 2}, 1},
		{"normal case 3", args{[][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6}, 2},
		{"normal case 4", args{[][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}, 15, 12}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numBusesToDestination(tt.args.routes, tt.args.source, tt.args.target); got != tt.want {
				t.Errorf("numBusesToDestination() = %v, want %v", got, tt.want)
			}
		})
	}
}
