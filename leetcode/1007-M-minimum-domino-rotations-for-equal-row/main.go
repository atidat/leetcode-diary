package main

import "math"

func main() {

}

func minDominoRotations(a []int, b []int) int {
	min := math.MaxInt32
	for i := 1; i < 7; i++ {
		acnt, bcnt := 0, 0
		j := 0
		for ; j < len(a); j++ {
			if a[j] != i && b[j] != i {
				break
			}
			if a[j] == i && b[j] == i {
				continue
			}
			if a[j] != i {
				acnt++
			}
			if b[j] != i {
				bcnt++
			}
		}
		if j == len(a) {
			min = mmin(mmin(acnt, bcnt), min)
		}
	}
	if min != math.MaxInt32 {
		return min
	}
	return -1
}

func mmin(a, b int) int {
	 if a <  b {
	 	return a
	 }
	 return b
}