package main

import "testing"

func comp(a, b [][]int) bool {
	for i := range a {
		for j := range a[i] {
			if a[i][j] != b[i][i] {
				return false
			}
		}
	}
	return true
}

func TestFirst(t *testing.T) {
	rowSum := []int{3, 8}
	colSum := []int{4, 7}
	tar := [][]int{
		{3, 0},
		{1, 7},
	}
	if !comp(tar, restoreMatrix(rowSum, colSum)) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	rowSum := []int{5, 7, 10}
	colSum := []int{8, 6, 8}
	tar := [][]int{
		{0, 5, 0},
		{6, 1, 0},
		{2, 0, 8},
	}
	if !comp(tar, restoreMatrix(rowSum, colSum)) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	rowSum := []int{14, 9}
	colSum := []int{6, 8, 6}
	tar := [][]int{
		{0, 9, 5},
		{6, 0, 3},
	}
	if !comp(tar, restoreMatrix(rowSum, colSum)) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	rowSum := []int{1, 0}
	colSum := []int{1}
	tar := [][]int{
		{1},
		{0},
	}
	if !comp(tar, restoreMatrix(rowSum, colSum)) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFifth(t *testing.T) {
	rowSum := []int{0}
	colSum := []int{0}
	tar := [][]int{
		{0},
	}
	if !comp(tar, restoreMatrix(rowSum, colSum)) {
		t.Error("=============== ERROR 3===============")
	}
}