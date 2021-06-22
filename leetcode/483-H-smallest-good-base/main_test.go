package main

import "testing"

func Test1(t *testing.T) {
	if "3" != smallestGoodBase("13") {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	if "8" != smallestGoodBase("4681") {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	if "999999999999999999" != smallestGoodBase("1000000000000000000") {
		t.Error("=============== ERROR 3===============")
	}
}

func Test4(t *testing.T) {
	if "2" != smallestGoodBase("2251799813685247") {
		t.Error("=============== ERROR 4===============")
	}
}