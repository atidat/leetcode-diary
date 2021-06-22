package main

import "testing"

func TestFirst(t *testing.T) {
	if true == isIsomorphic("egg", "add") {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestSecond(t *testing.T) {
	if false == isIsomorphic("foo", "bar") {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestThird(t *testing.T) {
	if true == isIsomorphic("paper", "title") {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestFourth(t *testing.T) {
	if false == isIsomorphic("ab", "aa") {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}
