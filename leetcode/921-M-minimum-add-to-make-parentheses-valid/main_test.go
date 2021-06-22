package main

import "testing"

func Test1(t *testing.T) {
	S := "())"
	if 1 != minAddToMakeValid(S) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	S := "((("
	if 3 != minAddToMakeValid(S) {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	S :="()"
	if 0 != minAddToMakeValid(S) {
		t.Error("=============== ERROR 3===============")
	}
}

func Test4(t *testing.T) {
	S := "()))(("
	if 4 != minAddToMakeValid(S) {
		t.Error("=============== ERROR 4===============")
	}
}

func Test5(t *testing.T) {
	S := "))()())("
	if 4 != minAddToMakeValid(S) {
		t.Error("=============== ERROR 5===============")
	}
}

func Test6(t *testing.T) {
	S := "((()))"
	if 0 != minAddToMakeValid(S) {
		t.Error("=============== ERROR 6===============")
	}
}