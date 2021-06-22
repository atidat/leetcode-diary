package main

import "testing"

func TestFirst(t *testing.T) {
	mat := [][]int{
		{1,0,1},
		{1,1,0},
		{1,1,0},
	}
	if 13 != numSubmat(mat) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	mat := [][]int{
		{0,1,1,0},
		{0,1,1,1},
		{1,1,1,0},
	}
	if 24 != numSubmat(mat) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	mat := [][]int{
		{1,1,1,1,1,1},
	}
	if 21 != numSubmat(mat) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	mat := [][]int{
		{1,0,1},
		{0,1,0},
		{1,0,1},
	}
	if 5 != numSubmat(mat) {
		t.Error("=============== ERROR 4==============")
	}
}

func TestFifth(t *testing.T) {
	mat := [][]int{
		{1,1,1,1,0},
		{1,0,0,1,0},
		{0,0,1,0,1},
		{0,1,0,0,0},
	}
	if 17 != numSubmat(mat) {
		t.Error("=============== ERROR 5==============")
	}
}