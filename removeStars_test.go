package main

import "testing"

func Test_removeStars(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case-1", args{"leet**cod*e"}, "lecoe"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeStars(tt.args.s); got != tt.want {
				t.Errorf("removeStars() = %v, want %v", got, tt.want)
			}
		})
	}
}
