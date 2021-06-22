package main

import "testing"

func TestOne(t *testing.T) {
	if 6 != nthUglyNumber(6) {
		t.Error("=============== ERROR 1 ===============")
	}
}

func TestTwo(t *testing.T) {
	if 8 != nthUglyNumber(7) {
		t.Error("=============== ERROR 2 ===============")
	}
}

func TestThree(t *testing.T) {
	if 12 != nthUglyNumber(10) {
		t.Error("=============== ERROR 3 ===============")
	}
}

func TestFour(t *testing.T) {
	if 2123366400 != nthUglyNumber(1690) {
		t.Error("=============== ERROR 4 ===============")
	}
}
