package main

import (
	"fmt"
	"testing"
)

func compareSlice(a, b []int) bool {
	fmt.Println("expect", b)
	lena := len(a)
	if lena != len(b) {
		return false
	}
	for i := 0; i < lena; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}


func TestFirst(t *testing.T) {
	if !compareSlice(advantageCount([]int{2, 7, 11, 15}, []int{1, 10, 4, 11}),
		[]int{2, 11, 7, 15}) {
		t.Error("=============== ERROR 1 ===============")
	}
}

func TestTwo(t *testing.T) {
	if !compareSlice(advantageCount([]int{12, 24, 8, 32}, []int{13, 25, 32, 11}),
		[]int{24, 32, 8, 12}) {
		t.Error("=============== ERROR 2 ===============")
	}
}

func TestThree(t *testing.T) {
	if !compareSlice(advantageCount([]int{2,0,4,1,2}, []int{1,3,0,0,2}),
		[]int{2, 0, 2, 1, 4}) {
		t.Error("=============== ERROR 3 ===============")
	}
}

/*
0, 1, 2, 2, 4
[{0, 2}, {0, 3}, {1,0}, {2, 4}, {3, 1}]




=> [{1,0}, {3, 1}, {0, 2}, {0, 3}, {2, 4}]
*/