package main

import "testing"

func TestFirst(t *testing.T) {
	if 1 != numRescueBoats([]int{1, 2}, 3) {
		t.Error("=============== ERROR 1 ===============")
	}
}

func TestSec(t *testing.T) {
	if 3 != numRescueBoats([]int{3, 2, 2, 1}, 3) {
		t.Error("=============== ERROR 2 ===============")
	}
}

func TestTrd(t *testing.T) {
	if 4 != numRescueBoats([]int{3, 5, 3, 4}, 5) {
		t.Error("=============== ERROR 3 ===============")
	}
}