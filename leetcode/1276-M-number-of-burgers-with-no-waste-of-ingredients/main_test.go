package main

import (
	"fmt"
	"testing"
)

func cmp(a, b []int) bool {
	fmt.Println(a, b)
	if len(a) == 0 && len(b) == 0 {
		return false
	}
	if len(a) == 0 || len(b) == 0 {
		return true
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return true
		}
	}
	return false
}

func TestFirst(t *testing.T) {
	var a = []int{1, 6}
	if cmp(a, numOfBurgers(16, 7)) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	var a = []int{}
	if cmp(a, numOfBurgers(17, 4)) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	var a = []int{}
	if cmp(a, numOfBurgers(4, 17)) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	var a = []int{0, 0}
	if cmp(a, numOfBurgers(0, 0)) {
		t.Error("=============== ERROR 4===============")
	}
}

func TestFifth(t *testing.T) {
	var a = []int{0, 1}
	if cmp(a, numOfBurgers(2,1)) {
		t.Error("=============== ERROR 5===============")
	}
}