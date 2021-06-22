package main

import "testing"

func Test1(t *testing.T) {
	if 2 != countArrangement(2) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	if 1 != countArrangement(1) {
		t.Error("=============== ERROR 2===============")
	}
}
/*
func Test3(t *testing.T) {
	if !=  {
		t.Error("=============== ERROR 3===============")
	}
}*/