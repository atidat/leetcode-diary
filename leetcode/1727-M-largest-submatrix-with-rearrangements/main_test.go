package main

import "testing"

func Test1(t *testing.T) {
	matrix := [][]int{{0,0,1},{1,1,1},{1,0,1}}
	if 4 != largestSubmatrix(matrix) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	matrix := [][]int{{1,0,1,0,1}}
	if 3 != largestSubmatrix(matrix) {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	matrix := [][]int{{1,1,0},{1,0,1}}
	if 2 != largestSubmatrix(matrix) {
		t.Error("=============== ERROR 3===============")
	}
}

func Test4(t *testing.T) {
	matrix := [][]int{{0,0},{0,0}}
	if 0 != largestSubmatrix(matrix) {
		t.Error("=============== ERROR 4===============")
	}
}