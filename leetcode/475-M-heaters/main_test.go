package main

import "testing"

func TestFirst(t *testing.T) {
	houses := []int{1,2,3}
	heaters := []int{2}
	if 1 != findRadius(houses, heaters) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	houses := []int{1,2,3,4}
	heaters := []int{1,4}
	if 1 != findRadius(houses, heaters) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	houses := []int{1,5}
	heaters := []int{2}
	if 3 != findRadius(houses, heaters) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	houses := []int{1}
	heaters := []int{2, 3}
	if 1 != findRadius(houses, heaters) {
		t.Error("=============== ERROR 4===============")
	}
}