package main

import "testing"

func TestOne(t *testing.T) {
	if getMoneyAmount(10) != 21 {
		t.Error("=============== ERROR 1 ===============")
	}
}