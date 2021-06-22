package main

import "testing"

func TestFirst(t *testing.T) {
	if 4 != balancedStringSplit("RLRRLLRLRL") {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if 3 != balancedStringSplit("RLLLLRRRLR") {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if 1 != balancedStringSplit("LLLLRRRR") {
		t.Error("=============== ERROR 3===============")
	}
}