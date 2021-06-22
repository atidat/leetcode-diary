package main

import "testing"

func TestFirst(t *testing.T) {
	arr := []int{1,15,7,9,2,5,10}
	if 84 != maxSumAfterPartitioning(arr, 3) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	arr := []int{1,4,1,5,7,3,6,1,9,9,3}
	if 83 != maxSumAfterPartitioning(arr, 4) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	arr := []int{1}
	if 1 != maxSumAfterPartitioning(arr, 1) {
		t.Error("=============== ERROR 3===============")
	}
}