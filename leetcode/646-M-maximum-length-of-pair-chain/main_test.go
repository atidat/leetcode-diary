package main

import "testing"

func TestFirst(t *testing.T) {
	pairs := [][]int{
		{1,2},
		{2,3},
		{3,4},
	}
	if 2 != findLongestChain(pairs) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	pairs := [][]int{
		{2,3},
		{1,2},
		{2, 101},
		{3,4},
	}
	if 2 != findLongestChain(pairs) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	pairs := [][]int{
		{1,2},
		{3,4},
		{5,6},
		{7,108},
		{101,102},
		{103,104},
	}
	if 5 != findLongestChain(pairs) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	pairs := [][]int{
		{1,2},
		{3,6},
		{4,101},
		{7,8},
	}
	if 3 != findLongestChain(pairs) {
		t.Error("=============== ERROR 4===============")
	}
}

func TestFifth(t *testing.T) {
	pairs := [][]int{
		{5,6},
		{7,99},
		{102,103},
		{104,105},
		{100,105},
		{108,109},
	}
	if 5 != findLongestChain(pairs) {
		t.Error("=============== ERROR 5===============")
	}
}

func TestSixth(t *testing.T) {
	pairs := [][]int{
		{5,6},
		{7,8},
	}
	if 2 != findLongestChain(pairs) {
		t.Error("=============== ERROR 6===============")
	}
}

func TestSeventh(t *testing.T) {
	pairs := [][]int{
		{5,6},
	}
	if 1 != findLongestChain(pairs) {
		t.Error("=============== ERROR 7===============")
	}
}


func TestEighth(t *testing.T) {
	pairs := [][]int{
		{-1,1},
		{-2,7},
		{-5,8},
		{-3,8},
		{1,3},
		{-2,9},
		{-5,2},
	}
	if 1 != findLongestChain(pairs) {
		t.Error("=============== ERROR 8===============")
	}
}