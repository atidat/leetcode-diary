package main

import "testing"

func TestFirst(t *testing.T) {
	nums := []int{23,2,4,6,7}
	if true != checkSubarraySum(nums, 6) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	nums := []int{23,2,6,4,7}
	if true != checkSubarraySum(nums, 6) {
		t.Error("=============== ERROR 2===============")
	}
}

/*func TestThird(t *testing.T) {
	if true != checkSubarraySum([]int{0, 0}, 0) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	if false != checkSubarraySum([]int{0}, 0) {
		t.Error("=============== ERROR 4===============")
	}
}

func TestFifth(t *testing.T) {
	nums := []int{23,2,6,4,7}
	if false != checkSubarraySum(nums, 0) {
		t.Error("=============== ERROR 5===============")
	}
}*/

func TestSixth(t *testing.T) {
	nums := []int{1,0}
	if false != checkSubarraySum(nums, 2) {
		t.Error("=============== ERROR 6===============")
	}
}

func TestSeventh(t *testing.T) {
	nums := []int{0,0}
	if true != checkSubarraySum(nums, -1) {
		t.Error("=============== ERROR 7===============")
	}
}