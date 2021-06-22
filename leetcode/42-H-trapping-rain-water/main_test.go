package main

import "testing"

func Test1(t *testing.T) {
	height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	if 6 != trap(height) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	height := []int{4,2,0,3,2,5}
	if 9 != trap(height) {
		t.Error("=============== ERROR 2===============")
	}
}
/*
func Test3(t *testing.T) {
	if !=  {
		t.Error("=============== ERROR 3===============")
	}
}*/