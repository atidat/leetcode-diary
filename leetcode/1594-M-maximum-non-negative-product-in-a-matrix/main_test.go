package main

import "testing"

func TestFirst(t *testing.T) {
	grid := [][]int{
		{-1, -2, -3},
		{-2, -3, -3},
		{-3, -3, -2},
	}
	if -1 != maxProductPath(grid) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	grid := [][]int{
		{1, -2, 1},
		{1, -2, 1},
		{3, -4, 1},
	}
	if 8 != maxProductPath(grid) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	grid := [][]int{
		{1, 3},
		{0, -4},
	}
	if 0 != maxProductPath(grid) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	grid := [][]int{
		{1, 4, 4, 0},
		{-2, 0, 0, 1},
		{1, -1, 1, 1},
	}
	if 2 != maxProductPath(grid) {
		t.Error("=============== ERROR 4===============")
	}
}

func TestFifth(t *testing.T) {
	grid := [][]int{
		{1},
		{-2},
		{1},
	}
	if -1 != maxProductPath(grid) {
		t.Error("=============== ERROR 5===============")
	}
}

func TestSixth(t *testing.T) {
	grid := [][]int{
		{1, 4, 4, 0},
	}
	if 0 != maxProductPath(grid) {
		t.Error("=============== ERROR 6===============")
	}
}

func TestSeven(t *testing.T) {
	grid := [][]int{
		{},
		{},
	}
	if 0 != maxProductPath(grid) {
		t.Error("=============== ERROR 7===============")
	}
}

func TestEigth(t *testing.T) {
	grid := [][]int{
		{2,1,3,0,-3,3,-4,4,0,-4},{-4,-3,2,2,3,-3,1,-1,1,-2},{-2,0,-4,2,4,-3,-4,-1,3,4},{-1,0,1,0,-3,3,-2,-3,1,0},{0,-1,-2,0,-3,-4,0,3,-2,-2},{-4,-2,0,-1,0,-3,0,4,0,-3},{-3,-4,2,1,0,-4,2,-4,-1,-3},{3,-2,0,-4,1,0,1,-3,-1,-1},{3,-4,0,2,0,-2,2,-4,-2,4},{0,4,0,-3,-4,3,3,-1,-2,-2},
	}
	if 19215865 != maxProductPath(grid) {
		t.Error("=============== ERROR 8===============")
	}
}