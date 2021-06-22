package main

import "testing"

func TestFirst(t *testing.T) {
	if false != canIWin(10, 11) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if true != canIWin(10, 0) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if true != canIWin(10, 1) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	if true != canIWin(4, 6) {
		t.Error("=============== ERROR 4===============")
	}
}

func TestFifth(t *testing.T) {
	if false != canIWin(10, 40) {
		t.Error("=============== ERROR 5===============")
	}
}

func TestSixth(t *testing.T) {
	if true != canIWin(7, 16) {
		t.Error("=============== ERROR 6===============")
	}
}