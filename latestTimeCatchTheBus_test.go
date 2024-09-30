package main

import "testing"

func Test_latestTimeCatchTheBus(t *testing.T) {
	type args struct {
		buses      []int
		passengers []int
		capacity   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case-1", args{[]int{10, 20}, []int{2, 17, 18, 19}, 2}, 16},
		{"case-2", args{[]int{20, 30, 10}, []int{19, 13, 26, 4, 25, 11, 21}, 2}, 20},
		{"case-3", args{[]int{3}, []int{2, 3}, 2}, 1},
		{"case-3", args{[]int{2}, []int{2}, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := latestTimeCatchTheBus(tt.args.buses, tt.args.passengers, tt.args.capacity); got != tt.want {
				t.Errorf("latestTimeCatchTheBus() = %v, want %v", got, tt.want)
			}
		})
	}
}
