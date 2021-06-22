package main

import "testing"

func TestFirst(t *testing.T) {
	nums := []int{3,4,2}
	if 6 != deleteAndEarn(nums) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	nums := []int{2,2,3,3,3,4}
	if 9 != deleteAndEarn(nums) {
		t.Error("=============== ERROR 2===============")
	}
}

/*func TestThird(t *testing.T) {
	if !=  {
		t.Error("=============== ERROR 3===============")
	}
}*/