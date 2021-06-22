package main

import "testing"

func Test1(t *testing.T) {
	graph := [][]int{{1,2},{3},{3},{}}
	allPathsSourceTarget(graph)
	if true {
		t.Error("[[0,1,3],[0,2,3]]")
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	graph := [][]int{{4,3,1},{3,2,4},{3},{4},{}}
	allPathsSourceTarget(graph)
	if true {
		t.Error("[[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]")
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	graph := [][]int{{1},{}}
	allPathsSourceTarget(graph)
	if true {
		t.Error("[[0,1]]")
		t.Error("=============== ERROR 3===============")
	}
}

func Test4(t *testing.T) {
	graph := [][]int{{1,2,3},{2},{3},{}}
	allPathsSourceTarget(graph)
	if true {
		t.Error("[[0,1,2,3],[0,2,3],[0,3]]")
		t.Error("=============== ERROR 4===============")
	}
}

func Test5(t *testing.T) {
	graph := [][]int{{1,3},{2}, {3},{}}
	allPathsSourceTarget(graph)
	if true {
		t.Error("[[0,1,2,3],[0,3]]")
		t.Error("=============== ERROR 5===============")
	}
}

func Test6(t *testing.T) {
	graph := [][]int{{2},{3}, {1},{}}
	allPathsSourceTarget(graph)
	if true {
		t.Error("[[0,2,1,3]]")
		t.Error("=============== ERROR 6===============")
	}
}