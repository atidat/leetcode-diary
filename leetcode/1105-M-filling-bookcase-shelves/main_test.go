package main

import "testing"

func TestFirst(t *testing.T) {
	books := [][]int{{1,1},{2,3},{2,3},{1,1},{1,1},{1,1},{1,2}}
	shelf_width := 4
	if 6 != minHeightShelves(books, shelf_width) {
		t.Error("=============== ERROR 1===============")
	}
}

/*func TestSecond(t *testing.T) {
	if !=  {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	if !=  {
		t.Error("=============== ERROR 3===============")
	}
}*/