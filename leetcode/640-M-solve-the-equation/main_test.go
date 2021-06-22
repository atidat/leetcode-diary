package main

import "testing"

func TestOne(t *testing.T) {
	if "x=2" != solveEquation("x+5-3+x=6+x-2") {
		t.Error("testOne Failed")
	}
}
