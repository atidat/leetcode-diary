package main

import "testing"

func TestFirst(t *testing.T) {
	A := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	if 12 != minFallingPathSum(A) {
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