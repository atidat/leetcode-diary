package main

import "testing"

func TestFirst(t *testing.T) {
	if 25 != robotSim([]int{4, -1, 3}, [][]int{}){
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if 65 != robotSim([]int{4, -1,4, -2,4}, [][]int{{2, 4}}) {
		t.Error("=============== ERROR 2===============")
	}
}

/*func TestThird(t *testing.T) {
	if != {
		t.Error("=============== ERROR 3===============")
	}
}*/