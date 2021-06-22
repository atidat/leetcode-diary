package main

import "testing"

func TestFirst(t *testing.T) {
	if 1 != minimumSwap("xx", "yy") {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if 2 != minimumSwap("xy", "yx") {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if -1 != minimumSwap("xx", "xy") {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	if 4 != minimumSwap("xxyyxyxyxx", "xyyxyxxxyx") {
		t.Error("=============== ERROR 4===============")
	}
}