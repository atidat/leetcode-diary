package main

import "testing"

func Test1(t *testing.T) {
	grid := [][]int{
		{0,0,1,0,0,0,0,1,0,0,0,0,0},
		{0,0,0,0,0,0,0,1,1,1,0,0,0},
		{0,1,1,0,1,0,0,0,0,0,0,0,0},
		{0,1,0,0,1,1,0,0,1,0,1,0,0},
		{0,1,0,0,1,1,0,0,1,1,1,0,0},
		{0,0,0,0,0,0,0,0,0,0,1,0,0},
		{0,0,0,0,0,0,0,1,1,1,0,0,0},
		{0,0,0,0,0,0,0,1,1,0,0,0,0},
	}
	if 6 != maxAreaOfIsland(grid) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	grid := [][]int{
		{0,0,0,0,0,0,0,0},
	}
	if 0 != maxAreaOfIsland(grid) {
		t.Error("=============== ERROR 2===============")
	}
}

/*func Test3(t *testing.T) {
	if !=  {
		t.Error("=============== ERROR 3===============")
	}
}*/