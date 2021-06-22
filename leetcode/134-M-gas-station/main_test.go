package main

import "testing"

func TestFirst(t *testing.T) {
	if 3 != canCompleteCircuit([]int{1,2,3,4,5}, []int{3,4,5,1,2}) {
		t.Error("=============== ERROR 1===============")
	}
}


func TestSecond(t *testing.T) {
	if -1 != canCompleteCircuit([]int{2,3,4}, []int{3,4,3}) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if 0 != canCompleteCircuit([]int{3,1,1}, []int{1,2,2}) {
		t.Error("=============== ERROR 3===============")
	}
}