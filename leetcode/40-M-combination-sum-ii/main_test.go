package main

import "testing"

func Test1(t *testing.T) {
	candidates := []int{10,1,2,7,6,1,5}
	target := 8
	combinationSum2(candidates, target)
	if true {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	candidates := []int{2,5,2,1,2}
	target := 5
	combinationSum2(candidates, target)
	if true {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	candidates := []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}
	target := 27
	combinationSum2(candidates, target)
	if true {
		t.Error("=============== ERROR 3===============")
	}
}

func Test4(t *testing.T) {
	candidates := []int{3,1,3,5,1,1}
	target := 8
	combinationSum2(candidates, target)
	if true {
		t.Error("=============== ERROR 4===============")
	}
}