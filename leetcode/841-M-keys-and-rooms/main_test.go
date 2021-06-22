package main

import "testing"

func Test1(t *testing.T) {
	rooms := [][]int{{1}, {2}, {3}, {}}
	if true != canVisitAllRooms(rooms) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	rooms := [][]int{{1,3},{3,0,1},{2},{0}}
	if false != canVisitAllRooms(rooms) {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	rooms := [][]int{{2}, {}, {1}}
	if true != canVisitAllRooms(rooms) {
		t.Error("=============== ERROR 3===============")
	}
}