package main

import "testing"

func TestFirst(t *testing.T) {
	if 9 != monotoneIncreasingDigits(10) {
		t.Error("=============== ERROR 1===============")
	}
}


func TestSecond(t *testing.T) {
	if 1234 != monotoneIncreasingDigits(1234) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if 299 != monotoneIncreasingDigits(332) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	if 1 != monotoneIncreasingDigits(1) {
		t.Error("=============== ERROR 4===============")
	}
}

func TestFive(t *testing.T) {
	if 3999 != monotoneIncreasingDigits(4431) {
		t.Error("=============== ERROR 5===============")
	}
}

func TestSix(t *testing.T) {
	if 3999 != monotoneIncreasingDigits(4031) {
		t.Error("=============== ERROR 6===============")
	}
}

func TestSeven(t *testing.T) {
	if 99 != monotoneIncreasingDigits(100) {
		t.Error("=============== ERROR 7===============")
	}
}