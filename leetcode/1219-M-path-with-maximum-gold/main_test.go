package main

import "testing"

func TestFirst(t *testing.T) {
	area := [][]int{
		{0,6,0},
		{5,8,7},
		{0,9,0},
	}
	if 24 != getMaximumGold(area) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	area := [][]int{
		{1,0,7},
		{2,0,6},
		{3,4,5},
		{0,3,0},
		{9,0,20},
	}
	if 28 != getMaximumGold(area) {
		t.Error("=============== ERROR 2===============")
	}
}

/*func TestThird(t *testing.T) {
	if !=  {
		t.Error("=============== ERROR 3===============")
	}
}*/