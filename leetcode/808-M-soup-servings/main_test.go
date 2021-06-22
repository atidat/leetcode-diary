package main

import "testing"

func TestFirst(t *testing.T) {
	if 0.625 != soupServings(50) {
		t.Error("=============== ERROR 1===============")
	}
}

/*func TestSecond(t *testing.T) {
	if !=  {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if !=  {
		t.Error("=============== ERROR 3===============")
	}
}*/