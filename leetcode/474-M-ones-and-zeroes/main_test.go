package main

import "testing"

func TestFirst(t *testing.T) {
	strs := []string{"10", "0001", "111001", "1", "0"}
	if 4 != findMaxForm(strs, 5,3) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	strs := []string{"10", "0", "1"}
	if 2 != findMaxForm(strs, 1, 1) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	strs := []string{"00","000"}
	if 0 != findMaxForm(strs, 1, 10) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	strs := []string{"0001","0101","1000","1000"}
	if 3 != findMaxForm(strs, 9, 3) {
		t.Error("=============== ERROR 4===============")
	}
}