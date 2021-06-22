package main

import "testing"

func TestOne(t *testing.T) {
	if 5 != maxProfit([]int{7, 1, 5, 3, 6, 4}) {
		t.Logf("====> Test One Failed.")
	}
}

func TestTwo(t *testing.T) {
	if 0 != maxProfit([]int{7, 6, 4, 3, 1}) {
		t.Logf("====> Test Two Failed.")
	}
}

func TestThree(t *testing.T) {
	if 5 != maxProfit([]int{1, 2, 3, 4, 5, 6}) {
		t.Logf("====> Test Three Failed.")
	}
}

func TestFour(t *testing.T) {
	if 0 != maxProfit([]int{1, 1, 1, 1, 1}) {
		t.Logf("====> Test Three Failed.")
	}
}

func TestFive(t *testing.T) {
	if 6 != maxProfit([]int{1, 7, 3, 6, 5, 4}) {
		t.Logf("====> Test Three Failed.")
	}
}
