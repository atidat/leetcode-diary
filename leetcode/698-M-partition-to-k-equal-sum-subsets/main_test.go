package main

import "testing"

func TestFirst(t *testing.T) {
	nums := []int{4, 3, 2, 3, 5, 2, 1}
	k := 4
	if true != canPartitionKSubsets(nums, k) {
		t.Error("=============== ERROR 1===============")
	}
}

/*func TestSecond(t *testing.T) {
	if !=  {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if !=  {
		t.Error("=============== ERROR 3===============")
	}
}*/