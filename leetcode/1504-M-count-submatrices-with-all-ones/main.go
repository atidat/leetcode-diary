package main

import (
	"fmt"
	"math"
)

func main() {
	mat := [][]int{
		{1, 0, 1},
		{1, 1, 0},
		{1, 1, 0},
	}
	numSubmat(mat)
}

func numSubmat(mat [][]int) int {
	pre(&mat)
	fmt.Println(mat)
	return executor(&mat)
}

func pre(m *[][]int) {
	for i := range *m {
		var cnt int
		for j := range (*m)[i] {
			if (*m)[i][j] != 0 {
				if cnt == 0 {
					(*m)[i][j] = 1
				} else {
					(*m)[i][j] = cnt+1
				}
				cnt++
			} else {
				cnt = 0
			}
		}
	}
}

func executor(m *[][]int) int {
	var sum int
	ml, mml := len(*m), len((*m)[0])

	for j := 0; j < mml; j++ {
		v := math.MaxInt32
		h := 1
		for i := ml-1; i >= 0; i-- {
			if (*m)[i][j] == 0 {
				h = 1
				v = math.MaxInt32
				continue
			}
			v = mmin((*m)[i][j], v)
			sum += v * h
			h++
		}
		fmt.Println(sum)
	}
	fmt.Println(sum)
	return sum
}

func mmin(a, b int) int {
	if a <  b { return a }
	return b
}