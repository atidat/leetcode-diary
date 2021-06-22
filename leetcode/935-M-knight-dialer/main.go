package main

import "fmt"

func main() {

}

func knightDialer(n int) int {
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, 10)
	}
	for i := 0; i < 10; i++ {
		dp[0][i] = 1
	}
	if n != 1 { dp[0][5] = 0 }

	m := pre()
	a, b := 0, 1
	for i := 1; i < n; i++ {
		for j := 0; j < 10; j++ {
			tmp := (*m)[j]

			var t int
			for k := 0; k < len(tmp); k++ {
				t += dp[a][tmp[k]]
			}
			dp[b][j] = t % 1000000007
		}
		a, b = b, a
	}

	var res int
	for i := 0; i < 10; i++ {
		res += dp[a][i]

	}
	fmt.Println(res % 1000000007)
	return res % 1000000007
}

func pre() *map[int][]int {
	m := make(map[int][]int)
	m[1] = []int{6,8}
	m[2] = []int{7,9}
	m[3] = []int{4,8}
	m[4] = []int{3,9,0}
	m[5] = []int{}
	m[6] = []int{1,7,0}
	m[7] = []int{2,6}
	m[8] = []int{1,3}
	m[9] = []int{2,4}
	m[0] = []int{4,6}
	return &m
}