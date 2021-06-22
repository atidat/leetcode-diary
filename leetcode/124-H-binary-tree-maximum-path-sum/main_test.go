package main

import "testing"

func Test1(t *testing.T) {
	fst1 := &TreeNode{2, nil, nil}
	fst2 := &TreeNode{3, nil, nil}
	root := &TreeNode{1, fst1, fst2}
	if 6 != maxPathSum(root) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	snd3 := &TreeNode{15, nil, nil}
	snd4 := &TreeNode{7, nil, nil}
	fst1 := &TreeNode{9, nil, nil}
	fst2 := &TreeNode{20, snd3, snd4}
	root := &TreeNode{-10, fst1, fst2}
	if 42 != maxPathSum(root) {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	root := &TreeNode{-3, nil, nil}
	if -3 != maxPathSum(root) {
		t.Error("=============== ERROR 3===============")
	}
}

func Test4(t *testing.T) {
	fst1 := &TreeNode{-3, nil, nil}
	root := &TreeNode{-1, fst1, nil}
	if -1 != maxPathSum(root) {
		t.Error("=============== ERROR 4===============")
	}
}

func Test5(t *testing.T) {
	fst1 := &TreeNode{-1, nil, nil}
	root := &TreeNode{-3, fst1, nil}
	if -1 != maxPathSum(root) {
		t.Error("=============== ERROR 5===============")
	}
}

func Test6(t *testing.T) {
	fst1 := &TreeNode{-2, nil, nil}
	fst2 := &TreeNode{-3, nil, nil}
	root := &TreeNode{1, fst1, fst2}
	if 1 != maxPathSum(root) {
		t.Error("=============== ERROR 6===============")
	}
}

func Test7(t *testing.T) {
	fst1 := &TreeNode{2, nil, nil}
	fst2 := &TreeNode{-3, nil, nil}
	root := &TreeNode{1, fst1, fst2}
	if 3 != maxPathSum(root) {
		t.Error("=============== ERROR 7===============")
	}
}

func Test8(t *testing.T) {
	fst1 := &TreeNode{-2, nil, nil}
	fst2 := &TreeNode{3, nil, nil}
	root := &TreeNode{1, fst1, fst2}
	if 4 != maxPathSum(root) {
		t.Error("=============== ERROR 8===============")
	}
}

func Test9(t *testing.T) {
	fth1 := &TreeNode{-4, nil, nil}
	trd2 := &TreeNode{2, fth1, nil}
	snd1 := &TreeNode{4, nil, trd2}
	fst1 := &TreeNode{5, snd1, nil}
	root := &TreeNode{-1, fst1, nil}
	if 11 != maxPathSum(root) {
		t.Error("=============== ERROR 9===============")
	}
}

func Test10(t *testing.T) {
	fif2 := &TreeNode{2, nil, nil}
	for2 := &TreeNode{-9, nil, fif2}
	trd1 := &TreeNode{-3, nil, for2}
	snd2 := &TreeNode{0, trd1, nil}
	snd1 := &TreeNode{-9, nil, nil}
	fst2 := &TreeNode{2, snd2, nil}
	fst1 := &TreeNode{8, nil, snd1}
	root := &TreeNode{-1, fst1, fst2}
	if 9 != maxPathSum(root) {
		t.Error("=============== ERROR 10===============")
	}
}