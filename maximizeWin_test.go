package main

import "testing"

func Test_maximizeWin(t *testing.T) {
	type args struct {
		prizePositions []int
		k              int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case-1", args{[]int{1, 1, 2, 2, 3, 3, 5}, 2}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximizeWin(tt.args.prizePositions, tt.args.k); got != tt.want {
				t.Errorf("maximizeWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
