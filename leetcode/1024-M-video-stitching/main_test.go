package main

import "testing"

func TestFirst(t *testing.T) {
	a := [][]int{
		{0,2},
		{4,6},
		{8,10},
		{1,9},
		{1,5},
		{5,9},
	}
	if 3 != videoStitching(a,10) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	a := [][]int{
		{0, 1},
		{1, 2},
	}
	if -1 != videoStitching(a, 5) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	a := [][]int{
		{0,1},{6,8},{0,2},
		{5,6},{0,4},{0,3},
		{6,7},{1,3},{4,7},
		{1,4},{2,5},{2,6},
		{3,4},{4,5},{5,7},{6,9},
	}
	if 3 != videoStitching(a, 9) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	a := [][]int{
		{0, 4},
		{2, 8},
	}
	if 2 != videoStitching(a, 5) {
		t.Error("=============== ERROR 4===============")
	}
}

func TestFifth(t *testing.T) {
	a := [][]int{
		{2, 4},
	}
	if 0 != videoStitching(a, 0) {
		t.Error("=============== ERROR 5===============")
	}
}

func TestSixth(t *testing.T) {
	a := [][]int{
		{5,7},{1,8},{0,0},{2,3},{4,5},{0,6},{5,10},{7,10},
	}
	if 1 != videoStitching(a, 5) {
		t.Error("=============== ERROR 6===============")
	}
}

func TestSeventh(t *testing.T) {
	a := [][]int{
		{8,10},{17,39},{18,19},{8,16},{13,35},{33,39},{11,19},{18,35},
	}
	if -1 != videoStitching(a, 20) {
		t.Error("=============== ERROR 7===============")
	}
}