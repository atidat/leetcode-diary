package main

import "testing"

func TestFirst(t *testing.T) {
	arr := []int{9,4,2,10,7,8,8,1,9}
	if 5 != maxTurbulenceSize(arr) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	arr := []int{4,8,12,16}
	if 2 != maxTurbulenceSize(arr) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	arr := []int{100}
	if 1 != maxTurbulenceSize(arr) {
		t.Error("=============== ERROR 3===============")
	}
}