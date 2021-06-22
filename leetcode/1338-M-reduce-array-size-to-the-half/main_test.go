package main

import "testing"

func TestFirst(t *testing.T) {
	if 2 != minSetSize([]int{3,3,3,3,5,5,5,2,2,7}) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if 1 != minSetSize([]int{7,7,7,7,7,7}) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if 1 != minSetSize([]int{1,9}) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	if 1 != minSetSize([]int{1000,1000,3,7}) {
		t.Error("=============== ERROR 4===============")
	}
}

func TestFifth(t *testing.T) {
	if 5 != minSetSize([]int{1,2,3,4,5,6,7,8,9,10}) {
		t.Error("=============== ERROR 5===============")
	}
}