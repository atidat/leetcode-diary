package main

import "testing"

func TestOne(t *testing.T) {
	if 4 != rob([]int{1, 2, 3, 1}) {
		t.Errorf("********Failure**************\n")
	}
}

func TestTwo(t *testing.T) {
	if 12 != rob([]int{2, 7, 9, 3, 1}) {
		t.Errorf("********Failure**************\n")
	}
}

func TestThree(t *testing.T) {
	if 0 != rob([]int{}) {
		t.Errorf("********Failure**************\n")
	}
}

func TestFour(t *testing.T) {
	if 1 != rob([]int{1}) {
		t.Errorf("********Failure**************\n")
	}
}

func TestFive(t *testing.T) {
	if 2 != rob([]int{1, 2}) {
		t.Errorf("********Failure**************\n")
	}
}

func TestSix(t *testing.T) {
	if 2 != rob([]int{2, 1}) {
		t.Errorf("********Failure**************\n")
	}
}

func TestSeven(t *testing.T) {
	if 4 != rob([]int{2, 1, 1, 2}) {
		t.Errorf("********Failure**************\n")
	}
}
