package main

import "testing"

func TestFirst(t *testing.T) {
	if 1 != countNumbersWithUniqueDigits(0) {
		t.Error("=============== ERROR 0 ===============")
	}
}

func TestSecond(t *testing.T) {
	if 10 != countNumbersWithUniqueDigits(1) {
		t.Error("=============== ERROR 1 ===============")
	}
}
func TestThird(t *testing.T) {
	if 91 != countNumbersWithUniqueDigits(2) {
		t.Error("=============== ERROR 2 ===============")
	}
}
func TestFour(t *testing.T) {
	if 739 != countNumbersWithUniqueDigits(3) {
		t.Error("=============== ERROR 3 ===============")
	}
}