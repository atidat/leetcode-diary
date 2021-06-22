package main

import "testing"

func Test1(t *testing.T) {
	if 6 != subsetXORSum([]int{1,3}) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	if 28 != subsetXORSum([]int{5, 1, 6}) {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	if 480 != subsetXORSum([]int{3,4,5,6,7,8}) {
		t.Error("=============== ERROR 3===============")
	}
}