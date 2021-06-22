package main

import "testing"

func TestFirst(t *testing.T) {
	if 2 != maxNumberOfFamilies(2, [][]int{{2,1},{1,8},{2,6}}) {
		t.Error("=============== ERROR 1 ===============")
	}
}

func TestSecond(t *testing.T) {
	if 4 != maxNumberOfFamilies(3, [][]int{{1,2},{1,3},{1,8},{2,6},{3,1}}) {
		t.Error("=============== ERROR 2 ===============")
	}
}

func TestThird(t *testing.T) {
	if 4 != maxNumberOfFamilies(4, [][]int{{4,3},{1,4},{4,6},{1,7}}) {
		t.Error("=============== ERROR 3 ===============")
	}
}

func TestFourth(t *testing.T) {
	if 1 != maxNumberOfFamilies(1, [][]int{{1,3}}) {
		t.Error("=============== ERROR 4 ===============")
	}
}