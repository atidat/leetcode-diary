package main

import "testing"

func TestFirst(t *testing.T) {
	nums := []int{2,1,6,4}
	if 1 != waysToMakeFair(nums) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	nums := []int{1,1,1}
	if 3 != waysToMakeFair(nums) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	nums := []int{1,2,3}
	if 0 != waysToMakeFair(nums) {
		t.Error("=============== ERROR 3===============")
	}
}


func TestFourth(t *testing.T) {
	nums := []int{1,2}
	if 0 != waysToMakeFair(nums) {
		t.Error("=============== ERROR 4===============")
	}
}

func TestFifth(t *testing.T) {
	nums := []int{1}
	if 1 != waysToMakeFair(nums) {
		t.Error("=============== ERROR 5===============")
	}
}
