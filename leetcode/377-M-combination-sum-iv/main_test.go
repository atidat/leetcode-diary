package main

import "testing"

func TestFirst(t *testing.T) {
	nums := []int{1, 2, 3}
	target := 4
	if 7 != combinationSum4(nums, target) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	nums := []int{}
	target := 4
	if 0 != combinationSum4(nums, target) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	nums := []int{0}
	target := 0
	if 1 != combinationSum4(nums, target) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	nums := []int{2,1,3}
	target := 35
	if 1132436852 != combinationSum4(nums, target) {
		t.Error("=============== ERROR 4===============")
	}
}

func TestFifth(t *testing.T) {
	nums := []int{2,1,3}
	target := 8
	if 81 != combinationSum4(nums, target) {
		t.Error("=============== ERROR 5===============")
	}
}

func TestSixth(t *testing.T) {
	nums := []int{2,1,3}
	target := 70
	if 1132436852 != combinationSum4(nums, target) {
		t.Error("=============== ERROR 6===============")
	}
}