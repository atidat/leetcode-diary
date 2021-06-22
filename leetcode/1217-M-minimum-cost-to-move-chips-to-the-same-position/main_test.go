package main

import "testing"

func TestFirst(t *testing.T) {
	if 1 != minCostToMoveChips([]int{1,2,3}) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if 2 != minCostToMoveChips([]int{2,2,2, 3,3}) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if 1 != minCostToMoveChips([]int{1,100000000}) {
		t.Error("=============== ERROR 3===============")
	}
}