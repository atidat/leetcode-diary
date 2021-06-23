package main

import "testing"

func Test1(t *testing.T) {
	letterCombinations("23")
	if true {
		t.Error("=============== ERROR ===============")
	}
}


func Test2(t *testing.T) {
	letterCombinations("")
	if true {
		t.Error("=============== ERROR ===============")
	}
}



func Test3(t *testing.T) {
	letterCombinations("2")
	if true {
		t.Error("=============== ERROR ===============")
	}
}
