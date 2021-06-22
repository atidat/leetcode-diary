package main

import "testing"

func TestFirst(t *testing.T) {
	if 1 != findContentChildren([]int{1,2,3}, []int{1,1}) {
		t.Error("=============== ERROR 1 ===============")
	}
}

func TestTwo(t *testing.T) {
	if 2 != findContentChildren([]int{1,2}, []int{1,2,3}) {
		t.Error("=============== ERROR 2===============")
	}
}
/*
func TestFirst(t *testing.T) {
	if 1 != findContentChildren([]int{}, []int{}) {
		t.Error("=============== ERROR ===============")
	}
}*/