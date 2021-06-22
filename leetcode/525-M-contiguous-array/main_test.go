package main

import "testing"

func Test1(t *testing.T) {
	nums := []int{0,1}
	if 2 != findMaxLength(nums) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	nums := []int{0,1,0}
	if 2 != findMaxLength(nums) {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	nums := []int{0,1,0,1,1,1,1,1,1,0,0,0,0,1,0,0,0,0,1,1}
	if 20 != findMaxLength(nums) {
		t.Error("=============== ERROR 3===============")
	}
}


func Test4(t *testing.T) {
	nums := []int{0,1,0,0,1,0,1,0,1,0,0,1,1,0,1,0,1}
	if 16 != findMaxLength(nums) {
		t.Error("=============== ERROR 4===============")
	}
}

func Test5(t *testing.T) {
	nums := []int{0,1,1,0,1,1,1,0}
	if 4 != findMaxLength(nums) {
		t.Error("=============== ERROR 5===============")
	}
}