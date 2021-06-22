package main

import "testing"

func TestFirst(t *testing.T) {
	if 10 != knightDialer(1) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if 20 != knightDialer(2) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if 46 != knightDialer(3) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	if 540641702 != knightDialer(100) {
		t.Error("=============== ERROR 4===============")
	}
}