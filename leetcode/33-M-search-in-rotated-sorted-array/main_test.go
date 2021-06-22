package main

import "testing"

func Test1(t *testing.T) {
	if 4 != search([]int{4,5,6,7,0,1,2} ,0) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	if -1 != search([]int{4,5,6,7,0,1,2} ,3) {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	if -1 != search([]int{1} ,0) {
		t.Error("=============== ERROR 3===============")
	}
}

func Test4(t *testing.T) {
	if 0 != search([]int{1} ,1) {
		t.Error("=============== ERROR 4===============")
	}
}

func Test5(t *testing.T) {
	if 0 != search([]int{2,1} ,2) {
		t.Error("=============== ERROR 5===============")
	}
}

func Test6(t *testing.T) {
	if -1 != search([]int{2,1} ,3) {
		t.Error("=============== ERROR 6===============")
	}
}

func Test7(t *testing.T) {
	if 3 != search([]int{4,5,6,7,0,1,2} ,7) {
		t.Error("=============== ERROR 7===============")
	}
}

func Test8(t *testing.T) {
	if 1 != search([]int{1,3} ,3) {
		t.Error("=============== ERROR 8===============")
	}
}

func Test9(t *testing.T) {
	if 4 != search([]int{4,5,6,7,8,1,2,3} ,8) {
		t.Error("=============== ERROR 9===============")
	}
}

func Test10(t *testing.T) {
	if -1 != search([]int{1,3,5} ,8) {
		t.Error("=============== ERROR 10===============")
	}
}

func Test11(t *testing.T) {
	if 0 != search([]int{1,3,5} ,1) {
		t.Error("=============== ERROR 11===============")
	}
}
func Test12(t *testing.T) {
	if -1 != search([]int{1,3,5} ,2) {
		t.Error("=============== ERROR 12===============")
	}
}