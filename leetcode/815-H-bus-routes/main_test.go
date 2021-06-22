package main

import "testing"

func Test1(t *testing.T) {
	routes := [][]int{{1,2,7},{3,6,7}}
	source, target := 1, 6
	if 2 != numBusesToDestination(routes, source, target) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	routes := [][]int{{7,12},{4,5,15},{6},{15,19},{9,12,13}}
	source, target := 15, 12
	if -1 != numBusesToDestination(routes, source, target) {
		t.Error("=============== ERROR 2===============")
	}
}

/*func Test3(t *testing.T) {
	if !=  {
		t.Error("=============== ERROR 3===============")
	}
}*/