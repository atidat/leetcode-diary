package main

import "testing"

func TestFirst(t *testing.T) {
	if 4 != longestStrChain([]string{"a","b","ba","bca","bda","bdca"})  {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if 5 != longestStrChain([]string{"xbc","pcxbcf","xb","cxbc","pcxbc"})  {
		t.Error("=============== ERROR 2===============")
	}
}

/*func TestThird(t *testing.T) {
	if !=  {
		t.Error("=============== ERROR 3===============")
	}
}*/