package main

import "testing"

func TestFirst(t *testing.T) {
	A := []int{3,6,9,12}
	if 4 != longestArithSeqLength(A) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	A := []int{9,4,7,2,10}
	if 3 != longestArithSeqLength(A) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	A := []int{20,1,15,3,10,5,8}
	if 4 != longestArithSeqLength(A) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	A := []int{3,6}
	if 2 != longestArithSeqLength(A) {
		t.Error("=============== ERROR 4===============")
	}
}

func TestFifth(t *testing.T) {
	A := []int{3,6, 8}
	if 2 != longestArithSeqLength(A) {
		t.Error("=============== ERROR 5===============")
	}
}