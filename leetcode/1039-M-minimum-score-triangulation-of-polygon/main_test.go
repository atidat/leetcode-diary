package main

import "testing"

func TestFirst(t *testing.T) {
	values := []int{1,2,3}
	if 6 != minScoreTriangulation(values) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	values := []int{3,7,4,5}
	if 144 != minScoreTriangulation(values) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	values := []int{1,3,1,4,1,5}
	if 13 != minScoreTriangulation(values) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	values := []int{1,2,3,55,23,6,12,74,100,32,76,12,74,1,2,3,55,23,6,12,74,100,32,76,12,74}
	if 34736 != minScoreTriangulation(values) {
		t.Error("=============== ERROR 4===============")
	}
}

func TestFifth(t *testing.T) {
	values := []int{5,3,5,5,1,6,2,3}
	if 88 != minScoreTriangulation(values) {
		t.Error("=============== ERROR 5===============")
	}
}