package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test-1", args{[]int{0, 1, 2, 3, 6, 7, 8}, 4}, []int{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.args.arr, tt.args.target)
			fmt.Println(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
