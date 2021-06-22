package main

import "testing"

func TestFirst(t *testing.T) {
	cost := []int{10, 15, 20}
	if 15 != minCostClimbingStairs(cost) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	if 6 != minCostClimbingStairs(cost) {
		t.Error("=============== ERROR 2===============")
	}
}

/*func TestThird(t *testing.T) {
	if !=  {
		t.Error("=============== ERROR 3===============")
	}
}*/