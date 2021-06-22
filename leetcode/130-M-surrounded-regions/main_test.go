package main

import (
	"fmt"
	"testing"
)


func TestOne(t *testing.T) {
	if true != compareSlice([]int{-1, -1, 2, 6}, genSlice) {
		t.Error("=====================Error 1======================")
	}
}

func TestTwo(t *testing.T) {
	if true != compareSlice([]int{1, -1, -1, 3, 4}, genSlice) {
		t.Error("=====================Error 2======================")
	}
}