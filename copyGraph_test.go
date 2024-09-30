package main

import (
	"context"
	"os"
	"reflect"
	"testing"

	"go.uber.org/goleak"
)

var testNode1, testNode2, testNode3, testNode4 *GNode

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)

	// 初始化测试数据
	node1 := &GNode{val: 1, nei: make([]*GNode, 0)}
	node2 := &GNode{val: 2, nei: make([]*GNode, 0)}
	node3 := &GNode{val: 3, nei: make([]*GNode, 0)}
	node4 := &GNode{val: 4, nei: make([]*GNode, 0)}

	node1.nei = append(node1.nei, node2, node4)
	node2.nei = append(node2.nei, node1, node3)
	node3.nei = append(node3.nei, node2, node4)
	node4.nei = append(node4.nei, node1, node3)

	// 将初始化的测试数据保存为全局变量，以便在测试用例中使用
	testNode1 = node1
	testNode2 = node2
	testNode3 = node3
	testNode4 = node4

	// 运行所有测试用例
	exitCode := m.Run()

	// 退出测试
	os.Exit(exitCode)
}

func Test_copyGraphDfs(t *testing.T) {
	type args struct {
		node *GNode
	}

	context.TODO()
	tests := []struct {
		name string
		args args
		want *GNode
	}{
		{"equal", args{testNode1}, testNode1},
		{"empty", args{nil}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyGraphDfs(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyGraphDfs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_copyGraphBfs(t *testing.T) {
	type args struct {
		node *GNode
	}
	tests := []struct {
		name string
		args args
		want *GNode
	}{
		{"equal", args{testNode1}, testNode1},
		{"empty", args{nil}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyGraphBfs(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyGraphBfs() = %v, want %v", got, tt.want)
			}
		})
	}
}
