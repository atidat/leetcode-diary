package main

import "testing"

func TestFirst(t *testing.T) {
	N := 5
	mines := [][]int{{4, 2}}
	if 2 != orderOfLargestPlusSign(N, mines) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	N := 2
	mines := [][]int{{}}
	if 1 != orderOfLargestPlusSign(N, mines) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	N := 1
	mines := [][]int{{0,0}}
	if 0 != orderOfLargestPlusSign(N, mines) {
		t.Error("=============== ERROR 3===============")
	}
}