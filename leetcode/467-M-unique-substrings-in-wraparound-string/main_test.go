package main

import "testing"

func TestFirst(t *testing.T) {
	if 1 != findSubstringInWraproundString("a") {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if 2 != findSubstringInWraproundString("cac") {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if 6 != findSubstringInWraproundString("zab") {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	if 0 != findSubstringInWraproundString("") {
		t.Error("=============== ERROR 4===============")
	}
}