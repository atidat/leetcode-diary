package main

import "testing"

func Test1(t *testing.T) {
	if true != isPowerOfFour(1) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	if false != isPowerOfFour(2) {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	if true != isPowerOfFour(64) {
		t.Error("=============== ERROR 3===============")
	}
}


func Test4(t *testing.T) {
	if false != isPowerOfFour(-1) {
		t.Error("=============== ERROR 4===============")
	}
}