package main

import (
	"testing"
)

func Test_example1(t *testing.T) {
	if solve([]int{1, 2}, []int{1, 0}) {
		t.FailNow()
	}
}

func Test_example2(t *testing.T) {
	if !solve([]int{1, 2, 4, 7}, []int{6, 7, 5, 3}) {
		t.FailNow()
	}
}

func Test_example3(t *testing.T) {
	if !solve([]int{1, 2, 4, 8}, []int{8, 4, 2, 1}) {
		t.FailNow()
	}
}

func Test_example4(t *testing.T) {
	if solve([]int{1, 2, 3, 4}, []int{1, 2, 4, 5}) {
		t.FailNow()
	}
}

func Test_example5(t *testing.T) {
	if solve([]int{1, 2, 0, 3}, []int{3, 3, 0, 3}) {
		t.FailNow()
	}
}

func Test_example6(t *testing.T) {
	if !solve([]int{3, 5, 6, 9, 10, 12}, []int{6, 5, 3, 12, 15, 9}) {
		t.FailNow()
	}
}
