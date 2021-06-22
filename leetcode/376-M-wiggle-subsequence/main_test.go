package main

import "testing"

func TestFirst(t *testing.T) {
	if 6 != wiggleMaxLength([]int{1, 7,4 ,9,2,5}) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if 7 != wiggleMaxLength([]int{1, 17, 5, 10, 13, 15,10, 5, 16, 8}) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if 2 != wiggleMaxLength([]int{1,2,3,4,5,6,7,8,9}) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	if 1 != wiggleMaxLength([]int{1,1}) {
		t.Error("=============== ERROR 4===============")
	}
}