package main

import (
	"testing"
)

func TestFirst(t *testing.T) {
	forbidden := []int{14,4,18,1,15}
	a,b, x := 3, 15, 9
	if 3 != minimumJumps(forbidden, a, b, x) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	forbidden := []int{8,3,16,6,12,20}
	a,b, x := 15, 13, 11
	if -1 != minimumJumps(forbidden, a, b, x) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	forbidden := []int{1,6,2,14,5,17,4}
	a,b, x := 16, 9, 7
	if 2 != minimumJumps(forbidden, a, b, x) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	forbidden := []int{61,104,19,60,68,157,183,148,116,93,190,13,177,47,15,133,111}
	a,b, x := 75, 165, 150
	if 2 != minimumJumps(forbidden, a, b, x) {
		t.Error("=============== ERROR 4===============")
	}
}