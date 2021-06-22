package main

import "fmt"

func main() {
	getSmallestString(3, 27)
	getSmallestString(5, 73)
}

/*
	1 <= n <= 105
	n <= k <= 26 * n
*/
func getSmallestString(n int, k int) string {
	// 'a': 1  'z': 26
	res := make([]byte, n)

	greedy(n, k, &res)
	fmt.Println(string(res))
	return string(res)
}

func greedy(i, k int, arr *[]byte) bool {
	// i: 差值
	if i == 1 {
		if k > 26 || k < 1 {
			return false
		}
		(*arr)[len(*arr) - 1] = byte('a' + k - 1)
		return true
	}

	b := byte('a')
	for ; b <= 'z'; b++ {
		if greedy(i-1, k-(int(b-'a')+1), arr) {
			(*arr)[len(*arr) - (i)] = b
			return true
		}
	}
	return false
}