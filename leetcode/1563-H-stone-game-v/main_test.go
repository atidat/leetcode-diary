package main

import "testing"

func TestFirst(t *testing.T) {
	stoneValue := []int{6,2,3,4,5,5}
	if 18 != stoneGameV(stoneValue) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	stoneValue := []int{7,7,7,7,7,7,7}
	if 28 != stoneGameV(stoneValue) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	stoneValue := []int{4}
	if 0 != stoneGameV(stoneValue) {
		t.Error("=============== ERROR 3===============")
	}
}