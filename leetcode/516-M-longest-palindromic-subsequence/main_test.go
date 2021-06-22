package main

import "testing"

func TestFirst(t *testing.T) {
	if 4 != longestPalindromeSubseq("bbbab") {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if 2 != longestPalindromeSubseq("cbbd") {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if 5 != longestPalindromeSubseq("abcba") {
		t.Error("=============== ERROR 3===============")
	}
}