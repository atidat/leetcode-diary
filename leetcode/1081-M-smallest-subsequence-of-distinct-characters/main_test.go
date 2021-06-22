package main

import "testing"

func TestFirst(t *testing.T) {
	if "adbc" != smallestSubsequence("cdadabcc") {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if "abcd" != smallestSubsequence("abcd") {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if "eacb" != smallestSubsequence("ecbacba") {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	if "letcod" != smallestSubsequence("leetcode") {
		t.Error("=============== ERROR 4===============")
	}
}

func TestFifth(t *testing.T) {
	if "bca" != smallestSubsequence("bcbcbcababa") {
		t.Error("=============== ERROR 5===============")
	}
}