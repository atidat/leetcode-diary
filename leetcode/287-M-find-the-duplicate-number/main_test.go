package main

import "testing"

func TestFirst(t *testing.T) {
	if 2 != findDuplicate([]int{1,3,4,2,2}) {
		t.Error("=============== ERROR First ===============")
	}
}

func TestSecond(t *testing.T) {
	if 3 != findDuplicate([]int{3, 1,3,4,2}) {
		t.Error("=============== ERROR Second ===============")
	}
}

func TestThird(t *testing.T) {
	if 3 != findDuplicate([]int{3,3,3,3,4}) {
		t.Error("=============== ERROR Third ===============")
	}
}

func TestFour(t *testing.T) {
	if 1 != findDuplicate([]int{1,1,2}) {
		t.Error("=============== ERROR Four ===============")
	}
}

func TestFive(t *testing.T) {
	if 5 != findDuplicate([]int{1,2,3,4,5,5,5,6,7,8,10}) {
		t.Error("=============== ERROR Five ===============")
	}
}

func TestSix(t *testing.T) {
	if 1 != findDuplicate([]int{1,1}) {
		t.Error("=============== ERROR Six ===============")
	}
}

func TestSeven(t *testing.T) {
	if 1 != findDuplicate([]int{1,3,4,2,1}) {
		t.Error("=============== ERROR Seven ===============")
	}
}

func TestEight(t *testing.T) {
	if 2 != findDuplicate([]int{2,2,2,2,2,2,2,2,2,2}) {
		t.Error("=============== ERROR Eight ===============")
	}
}