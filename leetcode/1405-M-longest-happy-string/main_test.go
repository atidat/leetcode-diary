package main

import "testing"

func TestFirst(t *testing.T) {
	if "ccaccbcc" != longestDiverseString(1,1,7) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if "aabbc" != longestDiverseString(2,2,1) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if "aabaa" != longestDiverseString(7,1,0) {
		t.Error("=============== ERROR 3===============")
	}
}