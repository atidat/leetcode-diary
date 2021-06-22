package main

import "testing"

func TestFirst(t *testing.T) {
	if true != stoneGame([]int{5, 3, 4, 5}) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if true != stoneGame([]int{5, 3, 4, 5, 3, 5}) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if true != stoneGame([]int{1,2}) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	if true != stoneGame([]int{9,9,10,1,7,3}) {
		t.Error("=============== ERROR 4===============")
	}
}

func TestFifth(t *testing.T) {
	if true != stoneGame([]int{3,7,2,3}) {
		t.Error("=============== ERROR 5===============")
	}
}