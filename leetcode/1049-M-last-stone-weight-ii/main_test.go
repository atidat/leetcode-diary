package main

import "testing"

func Test1(t *testing.T) {
	stones := []int{2,7,4,1,8,1}
	if 1 != lastStoneWeightII(stones) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	stones := []int{8}
	if 8 != lastStoneWeightII(stones) {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	stones := []int{3,8}
	if 5 != lastStoneWeightII(stones) {
		t.Error("=============== ERROR 3===============")
	}
}

func Test4(t *testing.T) {
	stones := []int{31,26,33,21,40}
	if 5 != lastStoneWeightII(stones) {
		t.Error("=============== ERROR 4===============")
	}
}

func Test5(t *testing.T) {
	stones := []int{2,1}
	if 1 != lastStoneWeightII(stones) {
		t.Error("=============== ERROR 5===============")
	}
}
