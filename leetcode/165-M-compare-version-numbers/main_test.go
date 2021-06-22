package main

import "testing"

func TestFirst(t *testing.T) {
	if compareVersion("0.1", "1.1") == -1 {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestSecond(t *testing.T) {
	if compareVersion("1.0.1", "1") == 1 {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestThird(t *testing.T) {
	if compareVersion("7.5.2.4", "7.5.3") == -1 {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestFourth(t *testing.T) {
	if compareVersion("1.01", "1.001") == 0 {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestFifth(t *testing.T) {
	if compareVersion("1.0", "1.0.0") == 0 {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}
