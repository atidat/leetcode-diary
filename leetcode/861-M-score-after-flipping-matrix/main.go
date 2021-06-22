package main

import "fmt"

func main() {
}

func matrixScore(a [][]int) int {
	fmt.Println("in:", a)
	if len(a) == 0 || len(a[0]) == 0 {
		return 0
	}
	// 先保证每一行的高位都是1
	for i := 0; i < len(a); i++ {
		if a[i][0] == 0 {
			for j := 0; j < len(a[0]); j++ {
				reverse(&a[i][j])
			}
			continue
		}
	}
	fmt.Println("phase 1:", a)
	// 其次保证每一列 1 比 0 多
	for i := 1; i < len(a[0]); i++ {
		var zero, one int
		for j := 0; j < len(a); j++ {
			if a[j][i] == 0 {
				zero++
			} else {
				one++
			}
		}
		if one >= zero {
			continue
		}
		for j := 0; j < len(a); j++ {
			reverse(&a[j][i])
		}
		fmt.Println("phase 2:", a)
	}
	return calc(a)
}

func reverse(a *int) {
	if *a == 1 {
		*a = 0
	} else {
		*a = 1
	}
}

func calc(a [][]int) int {
	var res int
	for i := 0; i < len(a); i++ {
		bin := 1
		for j := len(a[0])-1; j >= 0; j-- {
			res += a[i][j] * bin
			bin <<= 1
		}
	}
	fmt.Println(res)
	return res
}