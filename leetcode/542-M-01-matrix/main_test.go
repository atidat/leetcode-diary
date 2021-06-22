package main

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	mat := [][]int{{0,0,0},{0,1,0},{0,0,0}}
	fmt.Println(updateMatrix(mat))
	if true {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	mat := [][]int{{0,0,0},{0,1,0},{1,1,1}}
	fmt.Println(updateMatrix(mat))
	if true {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	mat := [][]int{{0,0,0},{0,1,0},{1,1,0}}
	fmt.Println(updateMatrix(mat))
	if true {
		t.Error("=============== ERROR 3===============")
	}
}

func Test4(t *testing.T) {
	mat := [][]int{{0,0,1,0,1,1,1,0,1,1},{1,1,1,1,0,1,1,1,1,1},
		{1,1,1,1,1,0,0,0,1,1},{1,0,1,0,1,1,1,0,1,1},
		{0,0,1,1,1,0,1,1,1,1},{1,0,1,1,1,1,1,1,1,1},
		{1,1,1,1,0,1,0,1,0,1},{0,1,0,0,0,1,0,0,1,1},
		{1,1,1,0,1,1,0,1,0,1},{1,0,1,1,1,0,1,1,1,0}}
	fmt.Println(updateMatrix(mat))
	if true {
		t.Error("=============== ERROR 4===============")
	}
}