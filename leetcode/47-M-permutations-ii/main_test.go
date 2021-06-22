package main

import "testing"

func Test1(t *testing.T) {
	nums := []int{1,1,2}
	permuteUnique(nums)
	if false {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	nums := []int{1,2,3}
	permuteUnique(nums)
	if false {
		t.Error("=============== ERROR 2===============")
	}
}
/*
func Test3(t *testing.T) {
	if !=  {
		t.Error("=============== ERROR 3===============")
	}
}*/