package main

import "testing"

func TestFirst(t *testing.T) {
	stones := []int{5, 3, 1, 4, 2}
	if 6 != stoneGameVII(stones) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	stones := []int{7, 90, 5, 1, 100, 10, 10, 2}
	if 122 != stoneGameVII(stones) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	stones := []int{4, 2}
	if 4 != stoneGameVII(stones) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	stones := []int{4}
	if 0 != stoneGameVII(stones) {
		t.Error("=============== ERROR 4===============")
	}
}