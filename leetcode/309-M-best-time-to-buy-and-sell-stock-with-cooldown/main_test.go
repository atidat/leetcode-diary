package main

import (
	"testing"
)

func TestZero(t *testing.T) {
	if 2 != maxProfit([]int{1,2,3,0}) {
		t.Errorf("=============== Error 0==================")
	}
}
func TestOne(t *testing.T) {
	if 3 != maxProfit([]int{1,2,3,0,2}) {
		t.Errorf("=============== Error 1==================")
	}
}
func TestTwo(t *testing.T) {
	if 6 != maxProfit([]int{1,2,3,0,2,1,4,2,5,2}) {
		t.Errorf("=============== Error 2==================")
	}
}
func TestThree(t *testing.T) {
	if 6 != maxProfit([]int{6,1,3,2,4,7}) {
		t.Errorf("=============== Error 3==================")
	}
}