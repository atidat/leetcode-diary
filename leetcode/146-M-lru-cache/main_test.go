package main

import (
	"fmt"
	"testing"
)

func compareSlice(s1, s2 []int) bool {
	fmt.Printf("expect: %v, real: %v\n", s1, s2)
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func TestOne(t *testing.T) {
	genSlice := make([]int, 0)
	obj := Constructor(2)
	genSlice = append(genSlice, obj.Get(2))
	obj.Put(2, 6)
	genSlice = append(genSlice, obj.Get(1))
	obj.Put(1, 5)
	obj.Put(1, 2)
	genSlice = append(genSlice, obj.Get(1))
	genSlice = append(genSlice, obj.Get(2))

	if true != compareSlice([]int{-1, -1, 2, 6}, genSlice) {
		t.Error("=====================Error 1======================")
	}
}

func TestTwo(t *testing.T) {
	genSlice := make([]int, 0)
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	genSlice = append(genSlice, obj.Get(1))
	obj.Put(3, 3)
	genSlice = append(genSlice, obj.Get(2))
	obj.Put(4, 4)
	genSlice = append(genSlice, obj.Get(1))
	genSlice = append(genSlice, obj.Get(3))
	genSlice = append(genSlice, obj.Get(4))

	if true != compareSlice([]int{1, -1, -1, 3, 4}, genSlice) {
		t.Error("=====================Error 2======================")
	}
}