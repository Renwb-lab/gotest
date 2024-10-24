package main

import "testing"

func Test_destCity(t *testing.T) {
	type args struct {
		paths [][]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test-1",
			args{
				paths: [][]string{
					{"London", "New York"},
					{"New York", "Lima"},
					{"Lima", "Sao Paulo"},
				},
			},
			"Sao Paulo",
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := destCity(tt.args.paths); got != tt.want {
				t.Errorf("destCity() = %v, want %v", got, tt.want)
			}
		})
	}
}
