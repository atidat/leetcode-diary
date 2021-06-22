package main

import "testing"

func TestFirst(t *testing.T) {
	if convertToTitle(1) == "A" {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestSecond(t *testing.T) {
	if convertToTitle(28) == "AB" {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestThird(t *testing.T) {
	if convertToTitle(703) == "AAA" {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestFourth(t *testing.T) {
	if convertToTitle(52) == "AZ" {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestFifth(t *testing.T) {
	if convertToTitle(702) == "ZZ" {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}
