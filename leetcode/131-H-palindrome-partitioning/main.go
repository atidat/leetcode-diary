package main

import (
	"fmt"
)

func main() {
	fmt.Println(partition("aab"))
	fmt.Println(partition("a"))
	fmt.Println(partition("aabab"))
}

func judge(s string) *[][]bool {
	sl := len(s)
	dp := make([][]bool, sl)
	for i := range dp {
		dp[i] = make([]bool, sl)
	}
	dp[sl-1][sl-1] = true
	for i := sl-2; i >= 0; i-- {
		dp[i][i] = true
		for j := i+1; j < sl; j++ {
			if i+1 > j-1 {
				dp[i][j] = s[i] == s[j]
				continue
			}
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
			}
		}
	}
	return &dp
}

func partition(s string) [][]string {
	dp := judge(s)
	//fmt.Println(dp)
	cache := make([][][]string, len(s))
	cache[len(s)-1] = [][]string{{s[len(s)-1:]}}

	for i := len(s)-2; i >= 0; i-- {
		post := cache[i+1]
		for j := range post {
			tmp := []string{s[i:i+1]}
			tmp = append(tmp, post[j]...)
			cache[i] = append(cache[i], tmp)
		}

		for j := i+1; j < len(s); j++ {
			if !(*dp)[i][j] {
				continue
			}
			if j == len(s)-1 {
				tmp := []string{s[i:]}
				cache[i] = append(cache[i], tmp)
				break
			}

			/*if j > len(s)-1 {
				break
			}*/
			post := cache[j+1]
			for k := range post {
				tmp := []string{s[i:j+1]}
				tmp = append(tmp, post[k]...)
				cache[i] = append(cache[i], tmp)
			}
		}
		//fmt.Println(cache)
	}
	return cache[0]
}