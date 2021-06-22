package main

import "testing"

func TestFirst(t *testing.T) {
	if 36 != integerBreak(10) {
		t.Error("=============== ERROR 1 ===============")	}
}

func TestSecond(t *testing.T) {
	if 774840978 != integerBreak(56) {
		t.Error("=============== ERROR 2 ===============")
	}
}