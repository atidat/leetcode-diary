package main

import "testing"

func Test1(t *testing.T) {
	grid := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0},
	}
	if 2 != closedIsland(grid) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	grid := [][]int{
		{0,0,1,0,0},
		{0,1,0,1,0},
		{0,1,1,1,0},
	}
	if 1 != closedIsland(grid) {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	grid := [][]int{
		{1,1,1,1,1,1,1},
		{1,0,0,0,0,0,1},
		{1,0,1,1,1,0,1},
		{1,0,1,0,1,0,1},
		{1,0,1,1,1,0,1},
		{1,0,0,0,0,0,1},
		{1,1,1,1,1,1,1},
	}
	if 2 != closedIsland(grid) {
		t.Error("=============== ERROR 3===============")
	}
}