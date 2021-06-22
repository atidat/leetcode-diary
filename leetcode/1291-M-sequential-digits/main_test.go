package main

import "testing"

func comp(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestFirst(t *testing.T) {
	var a = []int{123, 234}
	if !comp(a, sequentialDigits(100, 300)) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	var a = []int{1234, 2345, 3456, 4567, 5678, 6789, 12345}
	if !comp(a, sequentialDigits(1000, 13000)) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	var a = []int{67, 78,89, 123}
	if !comp(a, sequentialDigits(58, 155)) {
		t.Error("=============== ERROR 3===============")
	}
}