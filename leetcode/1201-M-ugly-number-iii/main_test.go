package main

import "testing"

func TestFirst(t *testing.T) {
	if 4 != nthUglyNumber(3,2,3,5) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if 6 != nthUglyNumber(4,2,3,4){
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if 10 != nthUglyNumber(5,2,11,13) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	if 1999999984 != nthUglyNumber(1000000000, 2,217983653,336916467) {
		t.Error("=============== ERROR 4===============")
	}
}


