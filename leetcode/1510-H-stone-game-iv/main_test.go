package main

import "testing"

func TestFirst(t *testing.T) {
	if true != winnerSquareGame(1) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if false != winnerSquareGame(2) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if true != winnerSquareGame(4) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	if false != winnerSquareGame(7) {
		t.Error("=============== ERROR 4===============")
	}
}

func TestFifth(t *testing.T) {
	if false != winnerSquareGame(17) {
		t.Error("=============== ERROR 5===============")
	}
}