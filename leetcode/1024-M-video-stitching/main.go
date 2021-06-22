package main

import (
	"fmt"
	"math"
)

func main() {
	clips := [][]int{{0,2},{4,6},{8,10},{1,9},{1,5},{5,9}}
	T := 10
	videoStitching(clips, T)
}

/*贪心算法
func videoStitching(clips [][]int, T int) int {
	if T == 0 {
		return 0
	}
	allLong := preHandle(clips, T)
	if allLong[0] >= T {
		return 1
	} else if allLong[0] == math.MinInt32 {
		return -1
	}
	//fmt.Println(allLong)
	cnt := handle(allLong, T)
	//fmt.Println(cnt)
	return cnt
}

func preHandle(clips [][]int, T int) []int {
	allLong := make([]int, T+1)
	for i := 0; i <= T; i++ {
		allLong[i] = math.MinInt32
	}
	for i := range clips {
		if clips[i][0] <= T {
			allLong[clips[i][0]] = mmax(allLong[clips[i][0]], clips[i][1])
		}
	}
	return allLong
}

func handle(allLong []int, T int) int {
	cnt := 1
	i := 0
	for ;i <= T; {
		curFar := allLong[i]
		if curFar == math.MinInt32 {
			i++
			continue
		}
		far := math.MinInt32
		for j := i+1; j <= T && j <= curFar; j++ {
			tmpFar := allLong[j]
			if tmpFar > curFar {
				if tmpFar > far {
					far = tmpFar
					i = j
				}
			}
		}
		//fmt.Printf("far is %d, cnt is %d\n", far, cnt)
		if far <= curFar{
			return -1
		}
		cnt++
		if far >= T {
			break
		}
	}
	return cnt
}

func mmax(a, b int) int {
	if a < b {
		return b
	}
	return a
}
 */

/*动态规划*/
func videoStitching(clips [][]int, T int) int {
	dp := make([]int, T+1)
	for i := 1; i <= T; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i <= T; i++ {
		for j := range clips {
			if clips[j][0] < i && i <= clips[j][1] {
				dp[i] = mmin(dp[clips[j][0]]+1, dp[i])
			}
		}
	}
	fmt.Println(dp)
	if dp[T] == math.MaxInt32 {
		dp[T] = -1
	}
	return dp[T]
}

func mmin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
