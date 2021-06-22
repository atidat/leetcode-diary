package main

import "testing"

func TestFirst(t *testing.T) {
	grid := [][]int{{1,1,1},{1,0,1},{1,1,1}}
	if 9 != largest1BorderedSquare(grid) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	grid := [][]int{{1,1,0,0}}
	if 1 != largest1BorderedSquare(grid) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	grid := [][]int{{1,1,1,1},{1,0,1,1},{1,1,1,1},{1,1,1,1}}
	if 16 != largest1BorderedSquare(grid) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	grid := [][]int{{1,1,1},{1,1,0},{1,1,1},{0,1,1},{1,1,1}}
	if 4 != largest1BorderedSquare(grid) {
		t.Error("=============== ERROR 4===============")
	}
}
