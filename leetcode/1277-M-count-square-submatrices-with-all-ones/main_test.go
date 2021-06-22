package main

import "testing"

func TestFirst(t *testing.T) {
	matrix := [][]int{
		{0,1,1,1},
		{1,1,1,1},
		{0,1,1,1},
	}
	if 15 != countSquares(matrix) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	matrix := [][]int{
		{1,0,1},
		{1,1,0},
		{1,1,0},
	}
	if 7 != countSquares(matrix) {
		t.Error("=============== ERROR 2===============")
	}
}

/*func TestThird(t *testing.T) {
	if !=  {
		t.Error("=============== ERROR 3===============")
	}
}*/