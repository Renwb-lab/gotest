package main

import "testing"

func Test_mincostTickets(t *testing.T) {
	type args struct {
		days  []int
		costs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"test-1", args{[]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}}, 11},
		{"test-2", args{[]int{1, 4, 6, 7, 8, 20}, []int{7, 2, 15}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mincostTickets(tt.args.days, tt.args.costs); got != tt.want {
				t.Errorf("mincostTickets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mincostTickets2(t *testing.T) {
	type args struct {
		days  []int
		costs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test-1", args{[]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}}, 11},
		{"test-2", args{[]int{1, 4, 6, 7, 8, 20}, []int{7, 2, 15}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mincostTickets2(tt.args.days, tt.args.costs); got != tt.want {
				t.Errorf("mincostTickets() = %v, want %v", got, tt.want)
			}
		})
	}
}
