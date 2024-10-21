package main

import "testing"

func Test_smallestRangeII(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test-1", args{[]int{1}, 0}, 0},
		{"test-2", args{[]int{0, 10}, 2}, 6},
		{"test-3", args{[]int{1, 3, 6}, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestRangeII(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("smallestRangeII() = %v, want %v", got, tt.want)
			}
		})
	}
}
