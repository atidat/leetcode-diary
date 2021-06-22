package main

import "testing"

func TestFirst(t *testing.T) {
	if 110 != twoCitySchedCost([][]int{
		{10, 20},{30, 200},{400, 50},{30, 20},
	}) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	if 1859 != twoCitySchedCost([][]int{
		{259, 770},{448, 54},{926, 667},{184, 139},{840, 118},{577, 469},
	}) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if 3086 != twoCitySchedCost([][]int{
		{515, 563},{451, 713},{537, 709},{343, 819},
		{855, 779},{457, 60},{650, 359},{631, 42},
	}) {
		t.Error("=============== ERROR 3===============")
	}
}