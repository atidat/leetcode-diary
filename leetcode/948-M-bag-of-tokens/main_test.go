package main

import "testing"

func TestFirst(t *testing.T) {
	if 0 != bagOfTokensScore([]int{100}, 50) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if 1 != bagOfTokensScore([]int{100, 200}, 150) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if 2 != bagOfTokensScore([]int{100, 200, 300, 400}, 200) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	if 2 != bagOfTokensScore([]int{100, 200, 300, 400}, 500) {
		t.Error("=============== ERROR 4===============")
	}
}