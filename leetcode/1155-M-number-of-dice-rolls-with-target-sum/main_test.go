package main

import "testing"

func TestFirst(t *testing.T) {
	if 1 != numRollsToTarget(1,6,3) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if 6 != numRollsToTarget(2,6,7) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if 1 != numRollsToTarget(2,5,10) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFifth(t *testing.T) {
	if 0 != numRollsToTarget(1,2,3) {
		t.Error("=============== ERROR 5===============")
	}
}

func TestFourth(t *testing.T) {
	if 222616187 != numRollsToTarget(30,30,500) {
		t.Error("=============== ERROR 4===============")
	}
}

