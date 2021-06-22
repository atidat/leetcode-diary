package main

import "testing"

func TestFirst(t *testing.T) {
	snd3 := TreeNode{3, nil, nil}
	fst2 := TreeNode{2, &snd3, nil}
	root := TreeNode{1, nil, &fst2}
	if == {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestSencond(t *testing.T) {
	if == {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}