package main

import "testing"

func TestFirst(t *testing.T) {
	if 3 != minSteps(3) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if 179 != minSteps(179) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if 21 != minSteps(1000) {
		t.Error("=============== ERROR 3===============")
	}
}