package main

import "testing"

func TestFirst(t *testing.T) {
	if false != carPooling([][]int{{2,1,5}, {3,3,7}}, 4) {
		t.Error("=============== ERROR ===============")
	}
}

func TestSecond(t *testing.T) {
	if true != carPooling([][]int{{2,1,5}, {3,3,7}}, 5) {
		t.Error("=============== ERROR ===============")
	}
}

func TestThird(t *testing.T) {
	if true != carPooling([][]int{{2,1,5},{3,5,7}}, 3) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	if true != carPooling([][]int{{3,2,7},{3,7,9},{8,3,9}}, 11) {
		t.Error("=============== ERROR 4===============")
	}
}

func TestFive(t *testing.T) {
	if true != carPooling([][]int{{9,3,4},{9, 1, 7},{4,2,4}, {7,4,5}}, 23) {
		t.Error("=============== ERROR 5===============")
	}
}