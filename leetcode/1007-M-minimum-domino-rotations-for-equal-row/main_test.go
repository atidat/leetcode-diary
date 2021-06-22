package main

import "testing"

func TestFirst(t *testing.T) {
	if 2 != minDominoRotations([]int{2,1,2,4,2,2}, []int{5,2,6,2,3,2}){
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if -1 != minDominoRotations([]int{3,5,1,2,3}, []int{3,6,3,3,4}){
		t.Error("=============== ERROR 2===============")
	}
}


func TestThird(t *testing.T) {
	if 0 != minDominoRotations([]int{1,1,1,1,1,1,1,1},
	[]int{1,1,1,1,1,1,1,1}){
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	if 1 != minDominoRotations([]int{1,2,1,1,1,2,2,2},
		[]int{2,1,2,2,2,2,2,2}){
		t.Error("=============== ERROR 4===============")
	}
}