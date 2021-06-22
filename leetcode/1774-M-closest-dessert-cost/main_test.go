package main

import "testing"

func Test1(t *testing.T) {
	baseCosts, toppingCosts := []int{1,7},  []int{3,4}
	target := 10
	if 10 != closestCost(baseCosts, toppingCosts, target) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	baseCosts, toppingCosts := []int{3,10},  []int{2,5}
	target := 9
	if 8 != closestCost(baseCosts, toppingCosts, target) {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	baseCosts, toppingCosts := []int{2,3},  []int{4,5,100}
	target := 18
	if 17 != closestCost(baseCosts, toppingCosts, target) {
		t.Error("=============== ERROR 3===============")
	}
}

func Test4(t *testing.T) {
	baseCosts, toppingCosts := []int{10},  []int{1}
	target := 1
	if 10 != closestCost(baseCosts, toppingCosts, target) {
		t.Error("=============== ERROR 4===============")
	}
}