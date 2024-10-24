package main

import "testing"

func Test_maximumTotalCost(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"test-1", args{[]int{1, -2, 3, 4}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumTotalCost(tt.args.nums); got != tt.want {
				t.Errorf("maximumTotalCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumTotalCostV1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"test-1", args{[]int{1, -2, 3, 4}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumTotalCostV1(tt.args.nums); got != tt.want {
				t.Errorf("maximumTotalCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
