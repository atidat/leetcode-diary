package main

import "testing"

func TestFirst(t *testing.T) {
	if 3 != numTrees(5) {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}