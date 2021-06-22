package main

import "testing"

func TestFirst(t *testing.T) {
	ta := [][]int{
		{1,1,1,1},
		{1,2,2,1},
		{1,2,2,1},
		{1,1,1,1},
	}
	if !isPrintable(ta) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	ta := [][]int{
		{1,1,1,1},
		{1,1,3,3},
		{1,1,3,4},
		{5,5,1,4},
	}
	if !isPrintable(ta) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	ta := [][]int{
		{1,2,1},
		{2,1,2},
		{1,2,1},
	}
	if isPrintable(ta) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	ta := [][]int{
		{1,1,1},
		{3,1,3},
	}
	if isPrintable(ta) {
		t.Error("=============== ERROR 4===============")
	}
}