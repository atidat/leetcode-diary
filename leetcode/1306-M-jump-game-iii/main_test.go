package main

import "testing"

func Test1(t *testing.T) {
	arr := []int{4,2,3,0,3,1,2}
	start := 5
	if true != canReach(arr, start) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	arr := []int{4,2,3,0,3,1,2}
	start := 0
	if true != canReach(arr, start) {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	arr := []int{3,0,2,1,2}
	start := 2
	if false != canReach(arr, start) {
		t.Error("=============== ERROR 3===============")
	}
}

func Test4(t *testing.T) {
	arr := []int{1,1,1,1,0}
	start := 2
	if true != canReach(arr, start) {
		t.Error("=============== ERROR 4===============")
	}
}

func Test5(t *testing.T) {
	arr := []int{58,48,64,36,19,19,67,13,32,2,59,50,29,68,50,0,69,31,54,20,22,43,30,9,68,71,20,22,48,74,2,65,27,54,30,5,66,24,64,68,9,31,50,59,15,72,6,49,11,71,12,61,5,66,30,1,2,39,59,35,53,21,76,17,71,40,68,57,64,53,70,21,50,49,25,63,35}
	start := 46
	if true != canReach(arr, start) {
		t.Error("=============== ERROR 5===============")
	}
}