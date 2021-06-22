package main

import "testing"

func Test1(t *testing.T) {
	if "111011" != maximumBinaryString("000110") {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	if "01" != maximumBinaryString("01") {
		t.Error("=============== ERROR 2===============")
	}
}

/*func Test3(t *testing.T) {
	if !=  {
		t.Error("=============== ERROR 3===============")
	}
}*/