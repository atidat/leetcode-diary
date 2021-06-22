package main

import "testing"

func TestFirst(t *testing.T) {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	if 8 != maxProfit(prices, fee) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	prices := []int{1,3,7,5,10,3}
	fee := 3
	if 6 != maxProfit(prices, fee) {
		t.Error("=============== ERROR 2===============")
	}
}

/*func TestThird(t *testing.T) {
	if !=  {
		t.Error("=============== ERROR 3===============")
	}
}*/