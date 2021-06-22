package main

import "testing"

func TestFirst(t *testing.T) {
	if 39 != matrixScore([][]int{{0,0,1,1},{1,0,1,0},{1,1,0,0}}) {
		t.Error("=============== ERROR 1 ===============")
	}
}

func TestTwo(t *testing.T) {
	if 8 != matrixScore([][]int{{1,1},{1,1},{0,1}}) {
		t.Error("=============== ERROR 2 ===============")
	}
}
