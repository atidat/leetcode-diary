package main

import "testing"

func Test1(t *testing.T) {
	cost := []int{4,3,2,5,6,7,2,5,5}
	target := 9
	if "7772" != largestNumber(cost, target) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	cost := []int{7,6,5,5,5,6,8,7,8}
	target := 12
	if "85" != largestNumber(cost, target) {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	cost := []int{2,4,6,2,4,6,4,4,4}
	target := 5
	if "0" != largestNumber(cost, target) {
		t.Error("=============== ERROR 3===============")
	}
}

func Test4(t *testing.T) {
	cost := []int{6,10,15,40,40,40,40,40,40}
	target := 47
	if "32211" != largestNumber(cost, target) {
		t.Error("=============== ERROR 4===============")
	}
}

func Test5(t *testing.T) {
	cost := []int{1,1,1,1,1,1,1,3,2}
	target := 10
	if "7777777777" != largestNumber(cost, target) {
		t.Error("=============== ERROR 5===============")
	}
}