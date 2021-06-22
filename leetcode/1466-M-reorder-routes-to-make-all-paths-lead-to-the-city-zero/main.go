package main

import "fmt"

func main() {
	n := 6
	connections := [][]int{{0,1},{1,3},{2,3},{4,0},{4,5}}
	minReorder(n, connections)
}

func minReorder(n int, connections [][]int) int {

	re := mapping(&connections)

	vis := make([]bool, n-1)

	var res int
	roots := make([]int, 0)
	roots = append(roots, 0)

	for len(roots) != 0 {
		tar := roots[0]
		roots = roots[1:]
		handleRoot(&roots, tar, &connections, &res, &vis, re)
	}
	fmt.Println(res)
	return res
}

func handleRoot(roots *[]int, tar int, conn *[][]int,
	res *int, vis *[]bool, re map[int][]int) {

	for i := 0; i < len(re[tar]); i++ { // conn[i][0]是起点
		if (*vis)[re[tar][i]] {
			continue
		}

		if (*conn)[re[tar][i]][1] == tar {
			(*vis)[re[tar][i]] = true
			*roots = append(*roots, (*conn)[re[tar][i]][0])
			continue
		}

		// 需要扭转路线
		(*vis)[re[tar][i]] = true
		*res++
		(*conn)[re[tar][i]][0], (*conn)[re[tar][i]][1] = (*conn)[re[tar][i]][1], (*conn)[re[tar][i]][0]
		*roots = append(*roots, (*conn)[re[tar][i]][0])
	}

}


func mapping(conn *[][]int) map[int][]int {
	re := make(map[int][]int, 0)
	for k, vals := range *conn {
		if v, ok := re[vals[0]]; ok {
			re[vals[0]] = append(v, k)
		} else {
			re[vals[0]] = []int{k}
		}

		if v, ok := re[vals[1]]; ok {
			re[vals[1]] = append(v, k)
		} else {
			re[vals[1]] = []int{k}
		}
	}
	return re
}