package main

import "testing"

func TestFirst(t *testing.T) {
	in := []int{2, 3, -2, 4}
	if maxProduct(in) == 6 {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestSecond(t *testing.T) {
	in := []int{-2, 0, -1}
	if maxProduct(in) == 0 {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestThird(t *testing.T) {
	in := []int{-2}
	if maxProduct(in) == -2 {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestFour(t *testing.T) {
	in := []int{2, -5, -2, -4, 3}
	if maxProduct(in) == 24 {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestFive(t *testing.T) {
	in := []int{0, 2}
	if maxProduct(in) == 2 {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}
