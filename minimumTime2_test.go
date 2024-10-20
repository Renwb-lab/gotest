package main

import "testing"

func Test_minimumTime2(t *testing.T) {
	type args struct {
		time       []int
		totalTrips int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// {"test-1", args{[]int{2}, 1}, 2},
		{"test-2", args{[]int{1, 2, 3}, 5}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTime2(tt.args.time, tt.args.totalTrips); got != tt.want {
				t.Errorf("minimumTime2() = %v, want %v", got, tt.want)
			}
		})
	}
}
