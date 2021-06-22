package main

import "testing"

func TestFirst(t *testing.T) {
	stoneValue := []int{1,2,3,7}
	if "Bob" != stoneGameIII(stoneValue) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	stoneValue := []int{1,2,3,-9}
	if "Alice" != stoneGameIII(stoneValue) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	stoneValue := []int{1,2,3,6}
	if "Tie" != stoneGameIII(stoneValue) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	stoneValue := []int{1,2,3,-1,-2,-3,7}
	if "Alice" != stoneGameIII(stoneValue) {
		t.Error("=============== ERROR 4===============")
	}
}

func TestFifth(t *testing.T) {
	stoneValue := []int{-1,-2,-3}
	if "Tie" != stoneGameIII(stoneValue) {
		t.Error("=============== ERROR 5===============")
	}
}