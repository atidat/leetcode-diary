package main

import "testing"

func cmp(a, b[]int) bool {
	for k := 0; k < len(a); k++ {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}

func TestFirst(t *testing.T) {
	var a = []int{0,1,1,1,1,0}
	if cmp(a, maxDepthAfterSplit("(()())")) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	var a = []int{0,0,0,1,1,0,1,1}
	if cmp(a, maxDepthAfterSplit("()(())()")) {
		t.Error("=============== ERROR 2===============")
	}
}
