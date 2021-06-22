package main

import "testing"

func Test1(t *testing.T) {
	if "aay" != getSmallestString(3, 27) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	if "aaszz" != getSmallestString(5, 73) {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	if "aaszz" != getSmallestString(1000, 10000) {
		t.Error("=============== ERROR 3===============")
	}
}