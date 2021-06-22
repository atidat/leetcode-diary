package main

import "testing"

func Test1(t *testing.T) {
	if true != robot("URR", [][]int{{}}, 3,2)  {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	if false != robot("URR", [][]int{{2,2}}, 3,2)  {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	if true != robot("URR", [][]int{{4, 2}}, 3,2)  {
		t.Error("=============== ERROR 3===============")
	}
}