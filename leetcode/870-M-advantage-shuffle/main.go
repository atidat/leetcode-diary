package main

import
(
	"fmt"
)

func main() {
	advantageCount([]int{2, 7, 11, 15}, []int{1, 10, 4, 11})
	fmt.Println("=====>", []int{2, 11, 7, 15})
	fmt.Println()
	advantageCount([]int{12, 24, 8, 32}, []int{13, 25, 32, 11})
	fmt.Println("=====>", []int{24, 32, 8, 12})
}

/*
Given two arrays A and B of equal size, the advantage of A with respect to B is the number of indices i for which A[i] > B[i].

Return any permutation of A that maximizes its advantage with respect to B.


Example 1:

Input: A = [2,7,11,15], B = [1,10,4,11]
Output: [2,11,7,15]
Example 2:

Input: A = [12,24,8,32], B = [13,25,32,11]
Output: [24,32,8,12]


Note:

1 <= A.length = B.length <= 10000
0 <= A[i] <= 10^9
0 <= B[i] <= 10^9

解析：田忌赛马
*/
type ar struct {
	ind int
	val int
}

func advantageCount(A []int, B []int) []int {
	tmp := make([]int, len(A))

	arrow := make([]ar, 0)
	for i := 0; i < len(B); i++ {
		arrow = append(arrow, ar{i, B[i]})
	}

	for i := 0; i < len(arrow); i++ {
		tar := arrow[i].val
		var pos int
		tmp[arrow[i].ind], pos = erfen(A, tar)
		A = append(A[:pos], A[pos+1:]...)
	}
	fmt.Println("reality", tmp)
	return tmp
}

func erfen(a []int, v int) (int, int) {
	sus := -1
	var h int
	for h, t := 0, len(a)-1; h < t; {
		mid := h + (t-h)/2
		if a[mid] > v {
			t = mid
			if sus != -1 && a[mid] < sus {
				sus = a[mid]
			} else if sus == -1 {
				sus = a[mid]
			}
		} else {
			h = mid + 1
		}
	}
	return sus, h
}

