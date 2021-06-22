package main

import "testing"

func TestFirst(t *testing.T) {
	in := []int{3, 4, 5, 1, 2}
	if findMin(in) == 1 {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestSecond(t *testing.T) {
	in := []int{4, 5, 6, 7, 0, 1, 2}
	if findMin(in) == 0 {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestThird(t *testing.T) {
	in := []int{2, 2, 2, 0, 1}
	if findMin(in) == 0 {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestFoutrh(t *testing.T) {
	in := []int{1, 3, 3}
	if findMin(in) == 1 {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}
