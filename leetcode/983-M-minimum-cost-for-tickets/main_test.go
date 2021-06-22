package main

import "testing"

func TestFirst(t *testing.T) {
	days := []int{1,4,6,7,8,20}
	costs := []int{2,7,15}
	if 11 != mincostTickets(days, costs) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	days := []int{1,2,3,4,5,6,7,8,9,10,30,31}
	costs := []int{2,7,15}
	if 17 != mincostTickets(days, costs) {
		t.Error("=============== ERROR 2===============")
	}
}


func TestThird(t *testing.T) {
	days := []int{1,2,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,20,21,24,25,27,28,29,30,31,34,37,38,39,41,43,44,45,47,48,49,54,57,60,62,63,66,69,70,72,74,76,78,80,81,82,83,84,85,88,89,91,93,94,97,99}
	costs := []int{9,38,134}
	if 423 != mincostTickets(days, costs) {
		t.Error("=============== ERROR 3===============")
	}
}



func TestFourth(t *testing.T) {
	days := []int{1,4,6,7,8,20}
	costs := []int{7,2,15}
	if 6 != mincostTickets(days, costs) {
		t.Error("=============== ERROR 4===============")
	}
}