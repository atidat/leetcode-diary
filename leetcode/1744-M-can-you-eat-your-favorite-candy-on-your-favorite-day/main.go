package main

import "fmt"

func main() {
	fmt.Println(canEat([]int{7,4,5,3,8}, [][]int{{0,2,2},{4,2,4},{2,13,1000000000}}))
	fmt.Println(canEat([]int{5,2,6,4,1}, [][]int{{3,1,2},{4,10,3},{3,10,100},
	{4,100,30},{1,3,1}}))

}

func canEat(candiesCount []int, queries [][]int) []bool {
	cl := len(candiesCount)
	sum := make([]int, cl)
	sum[0] = candiesCount[0]
	for i := 1; i < cl; i++ {
		sum[i] = sum[i-1] + candiesCount[i]
	}

	res := make([]bool, len(queries))
	for i, _ := range queries {
		if judge(sum, queries, i) {
			res[i] = true
		}
	}
	return res
}

func judge(s []int, q [][]int, i int) bool { // 是否限定每日吃糖数固定
	fmt.Println(s, q[i][1])

	// q [糖类型  目标天数  单日吃糖限制量]
	days := q[i][1] + 1
	for t := 1; t <= q[i][2]; t++ {
		fmt.Printf("%d ", days*t)
		if days*t > s[q[i][0]] {
			if t == 1 {
				fmt.Println("---")
				return false
			}
			fmt.Println("+++ +++")
			return true

		}

		// s[q[i][0] >= limit * 天数+1 > s[q[i][0]-1]

		if days*t <= s[q[i][0]] &&
			(q[i][0] != 0 && s[q[i][0]-1] < days*t || q[i][0] == 0) {
			fmt.Println("+++")
			return true
		}
	}
	fmt.Println("--- ---")
	return false
}