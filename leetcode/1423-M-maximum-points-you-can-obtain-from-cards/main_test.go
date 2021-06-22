package main

import "testing"

func TestFirst(t *testing.T) {
	cardPoints := []int{1,2,3,4,5,6,1}
	k := 3
	if 12 != maxScore(cardPoints, k) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	cardPoints := []int{2,2,2}
	k := 2
	if 4 != maxScore(cardPoints, k) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	cardPoints := []int{53,14,91,35,51,9,80,27,6,15,77,86,34,62,55,45,91,45,23,75,66,42,62,13,34,18,89,67,93,83,100,14,92,73,48,2,47,93,99,100,88,84,48}
	k := 43
	if 2429 != maxScore(cardPoints, k) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	cardPoints := []int{1,1000,1}
	k := 1
	if 1 != maxScore(cardPoints, k) {
		t.Error("=============== ERROR 4===============")
	}
}

func TestFifth(t *testing.T) {
	cardPoints := []int{1,79,80,1,1,1,200,1}
	k := 3
	if 202 != maxScore(cardPoints, k) {
		t.Error("=============== ERROR 5===============")
	}
}