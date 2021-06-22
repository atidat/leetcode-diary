package main

import "testing"

func TestFirst(t *testing.T) {
	if 1 != minDeletionSize([]string{"cba","daf","ghi"}) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if 0 != minDeletionSize([]string{"a","b"}) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if 3 != minDeletionSize([]string{"zyx","wvu","tsr"}) {
		t.Error("=============== ERROR 3===============")
	}
}