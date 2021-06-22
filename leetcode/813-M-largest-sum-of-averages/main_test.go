package main

import "testing"

func TestFirst(t *testing.T) {
	a := []int{9,1,2,3,9}
	if 20 != largestSumOfAverages(a, 3) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	a := []int{1,2,3,4,5,6,7}
	if 20.500000 != largestSumOfAverages(a, 4) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	a := []int{4,1,7,5,6,2,3}
	if 18.166667 != largestSumOfAverages(a, 4) {
		t.Error("=============== ERROR 3===============")
	}
}