package main

import "testing"

func Test1(t *testing.T) {
	if 3 != hammingWeight(00000000000000000000000000001011) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	if 1 != hammingWeight(00000000000000000000000010000000) {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	if 31 != hammingWeight(11111111111111111111111111111101) {
		t.Error("=============== ERROR 3===============")
	}
}