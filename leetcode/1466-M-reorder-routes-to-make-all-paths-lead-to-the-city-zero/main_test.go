package main

import "testing"

func Test1(t *testing.T) {
	n := 6
	connections := [][]int{{0,1},{1,3},{2,3},{4,0},{4,5}}
	if 3 != minReorder(n, connections) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	n := 5
	connections := [][]int{{1,0},{1,2},{3,2},{3,4}}
	if 2 != minReorder(n, connections) {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	n := 3
	connections := [][]int{{1, 0}, {2, 0}}
	if 0 != minReorder(n, connections) {
		t.Error("=============== ERROR 3===============")
	}
}