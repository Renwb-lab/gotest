package main

import (
	"reflect"
	"testing"
)

func Test_singleNumberIII(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test-1", args{[]int{1, 2, 1, 3, 2, 5}}, []int{3, 5}},
		{"test-2", args{[]int{-1, 0}}, []int{-1, 0}},
		{"test-1", args{[]int{0, 1}}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumberIII(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleNumberIII() = %v, want %v", got, tt.want)
			}
		})
	}
}
