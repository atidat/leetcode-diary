package main

import "testing"

func TestFirst(t *testing.T) {
	if 0.0625 != knightProbability(3, 2, 0, 0) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if 0.029099689547439307 != knightProbability(25, 88, 3, 7) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if 1 != knightProbability(25, 0, 3, 7) {
		t.Error("=============== ERROR 3===============")
	}
}
