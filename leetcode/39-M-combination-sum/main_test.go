package main

import "testing"

func TestFirst(t *testing.T) {
	if 18 != maxSumDivThree([]int{3, 6, 5, 1, 8}) {
		t.Error("=============== ERROR ===============")
	}
}

func TestSecond(t *testing.T) {
	if 0 != maxSumDivThree([]int{4}) {
		t.Error("=============== ERROR ===============")
	}
}


func TestThree(t *testing.T) {
	if 12 != maxSumDivThree([]int{1, 2, 3, 4, 4}) {
		t.Error("=============== ERROR ===============")
	}
}


func TestFour(t *testing.T) {
	if 0 != maxSumDivThree([]int{4}) {
		t.Error("=============== ERROR ===============")
	}
}

func TestFive(t *testing.T) {
	if 15 != maxSumDivThree([]int{2,6,2,2,7}) {
		t.Error("=============== ERROR ===============")
	}
}

func TestSix(t *testing.T) {
	if 12 != maxSumDivThree([]int{4,1,5,3,1}) {
		t.Error("=============== ERROR ===============")
	}
}

func TestSeven(t *testing.T) {
	if 117 != maxSumDivThree([]int{2,14,15,17,6,18,12,18,15,4}) {
		t.Error("=============== ERROR ===============")
	}
}