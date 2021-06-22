package main

import "testing"

func TestFirst(t *testing.T) {
	ratings := []int{1,0,2}
	if 5 != candy(ratings) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	ratings := []int{1,2,2}
	if 4 != candy(ratings) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	ratings := []int{1,0,2,3,2,1,0}
	if 15 != candy(ratings) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	ratings := []int{1,3,2,2,1}
	if 7 != candy(ratings) {
		t.Error("=============== ERROR 4===============")
	}
}