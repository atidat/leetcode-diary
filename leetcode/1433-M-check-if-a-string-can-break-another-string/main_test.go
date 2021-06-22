package main

import "testing"

func TestFirst(t *testing.T) {
	if !checkIfCanBreak("abc", "xya") {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if checkIfCanBreak("abe", "acd") {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if !checkIfCanBreak("leetcodee", "interview") {
		t.Error("=============== ERROR 3===============")
	}
}