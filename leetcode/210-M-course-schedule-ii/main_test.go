package main

import "testing"

func Test1(t *testing.T) {
	numCourses := 2
	prerequisites := [][]int{{1,0}}
	findOrder(numCourses, prerequisites)
	if true {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	numCourses := 4
	prerequisites := [][]int{{1,0},{2,0},{3,1},{3,2}}
	findOrder(numCourses, prerequisites)
	if true {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	if !=  {
		t.Error("=============== ERROR 3===============")
	}
}