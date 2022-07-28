// 单元测试 文件名必须以_test.go结尾
package main

import (
	"fmt"
	"testing"
)

// TestPrint 单元测试
// 参数必须是testing.T
func TestPrint(t *testing.T) {
	fmt.Println("test print")
}

// 测试函数中没有assert的用法？
func TestAdd(t *testing.T) {
	var x, y = 1, 2
	var rel = 3
	if rel != add(x, y) {
		t.Errorf("test failed")
	}
}

// 组测试
func TestAddGroup(t *testing.T) {
	type data struct {
		want     int
		inX, inY int
	}

	var group = []data{
		{3, 1, 2},
		{4, 2, 2},
		{6, 2, 4},
	}

	for _, g := range group {
		if g.want != add(g.inX, g.inY) {
			t.Errorf("test failed")
		}
	}
}

// 子测试
func TestAddChild(t *testing.T) {
	type data struct {
		want     int
		inX, inY int
	}

	var child = map[string]data{
		"testcase1": {3, 1, 2},
		"testcase2": {4, 2, 2},
		"testcase3": {6, 2, 4},
	}

	for name, c := range child {
		t.Run(name, func(t *testing.T) {
			if c.want != add(c.inX, c.inY) {
				t.Errorf("test failed")
			}
		})
	}
}

// pprof test go 支持
// go tool pprof cpu.prof  打开交互式命令行
