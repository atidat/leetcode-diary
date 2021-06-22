package main

import "testing"

func TestFirst(t *testing.T) {
	if true != makesquare([]int{1,1,2,2,2}) {
		t.Error("=============== ERROR 1 ===============")
	}
}

func TestSecond(t *testing.T) {
	if false != makesquare([]int{3,3,3,3,4}) {
		t.Error("=============== ERROR 2 ===============")
	}
}