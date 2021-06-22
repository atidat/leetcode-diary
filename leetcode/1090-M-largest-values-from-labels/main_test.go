package main

import "testing"

func Test1(t *testing.T) {
	values, labels := []int{5,4,3,2,1}, []int{1,1,2,2,3}
	num_wanted, use_limit := 3, 1
	if 9 != largestValsFromLabels(values, labels, num_wanted, use_limit) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	values, labels := []int{5,4,3,2,1}, []int{1,3,3,3,2}
	num_wanted, use_limit := 3, 2
	if 12 != largestValsFromLabels(values, labels, num_wanted, use_limit) {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	values, labels := []int{9,8,8,7,6}, []int{0,0,0,1,1}
	num_wanted, use_limit := 3, 1
	if 16 != largestValsFromLabels(values, labels, num_wanted, use_limit) {
		t.Error("=============== ERROR 3===============")
	}
}

func Test4(t *testing.T) {
	values, labels := []int{9,8,8,7,6}, []int{0,0,0,1,1}
	num_wanted, use_limit := 3, 2
	if 24 != largestValsFromLabels(values, labels, num_wanted, use_limit) {
		t.Error("=============== ERROR 4===============")
	}
}

func Test5(t *testing.T) {
	values, labels := []int{5,54,99,76,64,35,70,99,58,47,35,69,58,92,59,97,14,31,30,12,26,13,7,16,3,75}, []int{1,2,2,4,1,0,4,4,0,2,4,1,3,3,5,3,5,1,0,2,4,4,1,3,4,2}
	num_wanted, use_limit := 13, 4
	if 970 != largestValsFromLabels(values, labels, num_wanted, use_limit) {
		t.Error("=============== ERROR 5===============")
	}
}

func Test6(t *testing.T) {
	values, labels := []int{3,0,3,0,6}, []int{0,2,1,1,0}
	num_wanted, use_limit := 4, 1
	if 9 != largestValsFromLabels(values, labels, num_wanted, use_limit) {
		t.Error("=============== ERROR 6===============")
	}
}