package main

import (
	"fmt"
)

func main() {
	num1, num2 := []int{2}, []int{}
	fmt.Println(findMedianSortedArrays(num1, num2))
}

func findMedianSortedArrays(arr1 []int, arr2 []int) float64 {

	k := len(arr1) + len(arr2)
	tar := k/2
	cnt := 0
	for cnt < tar {
		del(&arr1, &arr2, &tar, &cnt)
	}
	fmt.Println(arr1, arr2)

	// [2], [1,3]
	if k & 1 == 0 { // 偶数
		return float64(arr1[0] + arr2[0])/2.0
	}
	if len(arr1) != 0 && (len(arr2) == 0 || arr1[0] < arr2[0]) {
		return float64(arr1[0])
	}

	return float64(arr2[0])
}


func del(arr1 []int, a1h int, arr2 []int, a2h int, k int) int {
	a1l := len(arr1)
	// a1l < a2l
	if a1l == a1h {
		return arr2[a2h+k-1]
	}
	if k == 1 {
		return mmin(arr1[a1h], arr2[a2h])
	}

	a1i := a1h + k/2
	a2i := a2h + k/2
	tmp := mmin(a1i, a1l)
	if arr1[a1i] < arr2[a2i] {
		return del(arr1, tmp, arr2, a2h, tmp)
	} else {
		return del(arr1, a1h, arr2, a2i, k/2 - 1)
	}

	// 常规场景：从arr1和arr2各取一部分
	ind1, ind2 := a1h, a2h
	for cnt := 0; cnt < k; cnt++ {
		if arr1[ind1] < arr2[ind2] {
			a1h++
		} else {
			a2h++
		}
	}
}

func mmin(a, b int) int {
	if a < b { return a }
	return b
}