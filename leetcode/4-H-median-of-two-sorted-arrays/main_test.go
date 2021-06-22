package main

import (
	"testing"
)

func Test0(t *testing.T) {
	num1, num2 := []int{1,3}, []int{2,4,5,6}
	if 3.5 != findMedianSortedArrays(num1, num2) {
		t.Error("=============== ERROR 0===============")
	}
}

func Test1(t *testing.T) {
	num1, num2 := []int{1,3}, []int{2}
	if 2.0 != findMedianSortedArrays(num1, num2) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	num1, num2 := []int{1,2}, []int{3,4}
	if 2.50000 != findMedianSortedArrays(num1, num2) {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	num1, num2 := []int{0,0}, []int{0,0}
	if 0.0 != findMedianSortedArrays(num1, num2) {
		t.Error("=============== ERROR 3===============")
	}
}

func Test4(t *testing.T) {
	num1, num2 := []int{}, []int{1}
	if 1.0 != findMedianSortedArrays(num1, num2) {
		t.Error("=============== ERROR 4===============")
	}
}

func Test5(t *testing.T) {
	num1, num2 := []int{2}, []int{}
	if 2.0 != findMedianSortedArrays(num1, num2) {
		t.Error("=============== ERROR 5===============")
	}
}

func Test6(t *testing.T) {
	num1, num2 := []int{2}, []int{}
	if 2.0 != findMedianSortedArrays(num1, num2) {
		t.Error("=============== ERROR 6===============")
	}
}

func Test7(t *testing.T) {
	num1, num2 := []int{2}, []int{1,3,4}
	if 2.500000 != findMedianSortedArrays(num1, num2) {
		t.Error("=============== ERROR 7===============")
	}
}

func Test8(t *testing.T) {
	num1, num2 := []int{3}, []int{1,2,4}
	if 2.500000 != findMedianSortedArrays(num1, num2) {
		t.Error("=============== ERROR 8===============")
	}
}

func Test9(t *testing.T) {
	num1, num2 := []int{0,0,0,0,0}, []int{-1,0,0,0,0,0,1}
	if 0.0 != findMedianSortedArrays(num1, num2) {
		t.Error("=============== ERROR 9===============")
	}
}

func Test10(t *testing.T) {
	num1, num2 := []int{1,2,3}, []int{1,2,2}
	if 2.0 != findMedianSortedArrays(num1, num2) {
		t.Error("=============== ERROR 10===============")
	}
}
