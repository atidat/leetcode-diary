package main

import "testing"

func Test1(t *testing.T) {
	root := buildTree1()
	if 2 != pseudoPalindromicPaths(root) {
		t.Error("=============== ERROR 1===============")
	}
}

/*func Test2(t *testing.T) {
	if !=  {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	if !=  {
		t.Error("=============== ERROR 3===============")
	}
}*/