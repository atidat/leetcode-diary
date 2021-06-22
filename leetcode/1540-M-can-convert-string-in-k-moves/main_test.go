package main

import "testing"

func TestFirst(t *testing.T) {
	if true != canConvertString("input", "ouput", 9) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if false != canConvertString("abc", "bcd", 10) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if true != canConvertString("aab", "bbb", 27) {
		t.Error("=============== ERROR 3===============")
	}
}