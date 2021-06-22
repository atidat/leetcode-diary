package main

import "testing"

func TestFirst(t *testing.T) {
	A := []int{1,2,3,2,1}
	B := []int{3,2,1,4,7}
	if findLength(A,B) != 3 {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	A := []int{1,2,3,2,4}
	B := []int{3,2,1,4,7}
	if findLength(A,B) != 2 {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	A := []int{1,2,3,4,2}
	B := []int{3,2,1,4,7}
	if findLength(A,B) != 1 {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	A := []int{1,2,3,2,1}
	B := []int{4,3,2,1,7}
	if findLength(A,B) != 3 {
		t.Error("=============== ERROR 4===============")
	}
}

func TestFifth(t *testing.T) {
	A := []int{0,0,0,0,0}
	B := []int{0,0,0,0,0}
	if findLength(A,B) != 5 {
		t.Error("=============== ERROR 5===============")
	}
}