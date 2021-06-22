package main

import "testing"

func TestFirst(t *testing.T) {
	price := []int{2, 5}
	special := [][]int{
		{3, 0, 5},
		{1, 2, 10},
	}
	need := []int{3, 2}
	if 14 != shoppingOffers(price, special, need) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	price := []int{2, 3, 4}
	special := [][]int{
		{1, 1, 0, 4},
		{2, 2, 1, 9},
	}
	need := []int{1, 2, 1}
	if 11 != shoppingOffers(price, special, need) {
		t.Error("=============== ERROR 2===============")
	}
}

/*func TestThird(t *testing.T) {
	if !=  {
		t.Error("=============== ERROR 3===============")
	}
}*/
