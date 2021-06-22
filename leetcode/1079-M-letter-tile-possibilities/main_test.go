package main

import "testing"

func TestFirst(t *testing.T) {
	if 8 != numTilePossibilities("AAB") {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if 188 != numTilePossibilities("AAABBC") {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if 1 != numTilePossibilities("V") {
		t.Error("=============== ERROR 3===============")
	}
}