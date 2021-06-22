package main

import "testing"

func Test1(t *testing.T) {
	if 1 != peakIndexInMountainArray([]int{0,1,0}) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	if 1 != peakIndexInMountainArray([]int{0,2,1,0}) {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	if 1 != peakIndexInMountainArray([]int{0,10,5,2}) {
		t.Error("=============== ERROR 3===============")
	}
}

func Test4(t *testing.T) {
	if 2 != peakIndexInMountainArray([]int{3,4,5,1}) {
		t.Error("=============== ERROR 4===============")
	}
}

func Test5(t *testing.T) {
	if 2 != peakIndexInMountainArray([]int{24,69,100,99,79,78,67,36,26,19}) {
		t.Error("=============== ERROR 5===============")
	}
}