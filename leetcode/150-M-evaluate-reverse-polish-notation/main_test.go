package main

import "testing"

func TestFirst(t *testing.T) {
	in := []string{"2", "1", "+", "3", "*"}
	if evalRPN(in) == 9 {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestSecond(t *testing.T) {
	in := []string{"4", "13", "5", "/", "+"}
	if evalRPN(in) == 6 {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}
