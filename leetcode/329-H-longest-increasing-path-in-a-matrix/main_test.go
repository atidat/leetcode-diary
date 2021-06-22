package main

import "testing"

func Test1(t *testing.T) {
	matrix := [][]int{{9,9,4},{6,6,8},{2,1,1}}
	if 4 != longestIncreasingPath(matrix) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	matrix := [][]int{{3,4, 5},{3,2,6},{2,2,1}}
	if 4 != longestIncreasingPath(matrix) {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	matrix := [][]int{{1}}
	if 1 != longestIncreasingPath(matrix) {
		t.Error("=============== ERROR 3===============")
	}
}

func Test4(t *testing.T) {
	matrix := [][]int{{7,8,9},{9,7,6},{7,2,3}}
	if 6 != longestIncreasingPath(matrix) {
		t.Error("=============== ERROR 4===============")
	}
}

func Test5(t *testing.T) {
	matrix := [][]int{{17,4,6,11,4,0,17,12,19,12,12,0},
		{0,6,4,4,5,9,15,1,11,13,18,0},
		{4,2,13,1,2,7,15,5,14,6,8,6}}
	if 6 != longestIncreasingPath(matrix) {
		t.Error("=============== ERROR 5===============")
	}
}