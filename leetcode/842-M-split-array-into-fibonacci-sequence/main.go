package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(splitIntoFibonacci("123456579"))
	fmt.Println(splitIntoFibonacci("11235813"))
	fmt.Println(splitIntoFibonacci("112358130"))
	fmt.Println(splitIntoFibonacci("0123"))
	fmt.Println(splitIntoFibonacci("1101111"))
	fmt.Println(splitIntoFibonacci("1199110"))
	fmt.Println(splitIntoFibonacci("1011"))
}

func splitIntoFibonacci(S string) []int {
	fmt.Println(S)
	res := make([]int, 0)
	if executor(S, &res, 0) {
		return res
	}
	return []int{}
}

// s="123456579"
//   = [1] + dfs("23456579")
//     = [1,2] + dfs("3456579")
//     = [1,23] + dfs("456579")
//     = ...
//   = [12] + dfs("3456579")
//   = [12, 3] + dfs("456579")
//   = [12, 34] + dfs("56579")
//   = [12, 345] + dfs("56579")
func executor(s string, res *[]int, start int) bool {
	if start == len(s) {
		return len(*res) >= 3
	}
	length := 10
	if s[start] == '0' {
		length = 1
	}

	for i := start+1; i <= len(s) && i <= start + length; i++ {
		fmt.Println(*res)
		tmp, _ := strconv.Atoi(s[start:i])
		if tmp > math.MaxInt32 {
			break
		}
		if len(*res) >= 2 {
			if (*res)[len(*res)-2] + (*res)[len(*res)-1] > tmp {
				continue
			} else if (*res)[len(*res)-2] + (*res)[len(*res)-1] < tmp {
				break
			}
		}
		// len(*res) == 0 || 1
		*res = append(*res, tmp)
		if executor(s, res, i) {
			return true
		}
		*res = (*res)[:len(*res)-1]
	}
	return false
}
