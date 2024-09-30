package main

import "testing"

func Test_edgeScore(t *testing.T) {
	type args struct {
		edges []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"good-case-1", args{edges: []int{1, 0, 0, 0, 0, 7, 7, 5}}, 7},
		{"good-case-2", args{edges: []int{2, 0, 0, 2}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := edgeScore(tt.args.edges); got != tt.want {
				t.Errorf("edgeScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
