package main

import "testing"

func TestFirst(t *testing.T) {
	if "Radiant" != predictPartyVictory("RD") {
		t.Error("=============== ERROR 1===============")
	}
}

func TestTwo(t *testing.T) {
	if "Dire" != predictPartyVictory("RDD") {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThree(t *testing.T) {
	if "Radiant" != predictPartyVictory("RDDRDRRD") {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFour(t *testing.T) {
	if "Radiant" != predictPartyVictory("R") {
		t.Error("=============== ERROR 4===============")
	}
}

func TestFive(t *testing.T) {
	if "Dire" != predictPartyVictory("DDR") {
		t.Error("=============== ERROR 5===============")
	}
}

func TestSix(t *testing.T) {
	if "Radiant" != predictPartyVictory("RRD") {
		t.Error("=============== ERROR 6===============")
	}
}

func TestSeven(t *testing.T) {
	if "Radiant" != predictPartyVictory("DDRRRR") {
		t.Error("=============== ERROR 7===============")
	}
}