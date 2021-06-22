package main

import (
	"fmt"
	"sort"
)

func main() {
	bagOfTokensScore([]int{100, 200, 300, 400}, 500)
}

/*
You have an initial power P, an initial score of 0 points, and a bag of tokens.

Each token can be used at most once, has a value token[i], and has potentially two ways to use it.

If we have at least token[i] power, we may play the token face up, losing token[i] power, and gaining 1 point.
If we have at least 1 point, we may play the token face down, gaining token[i] power, and losing 1 point.
Return the largest number of points we can have after playing any number of tokens.

Example 1:
	Input: tokens = [100], P = 50
	Output: 0

Example 2:
	Input: tokens = [100,200], P = 150
	Output: 1

Example 3:
	Input: tokens = [100,200,300,400], P = 200
	Output: 2

Note:
	tokens.length <= 1000
	0 <= tokens[i] < 10000
	0 <= P < 10000

*/
func bagOfTokensScore(tokens []int, p int) int {
	// 先找到最小的能量值，用于换取分数
	// 用换到的分数，换取最大的能量
	// 循环以往
	sort.Ints(tokens)
	if p < tokens[0] {
		return 0
	}

	points := 0
//	cnt := 0
	for h, t := 0, len(tokens)-1; h <= t; {
		fmt.Println("h, t, points, p:", h, t, points, p)
		if p > tokens[h] {
			p -= tokens[h]
			points++
			h++
//			cnt++
		}

		if p >= tokens[t] {
			p -= tokens[t]
			points++
			t--
//			cnt++
			continue
		}

		if h == t {
			break
		}

		if points > 0 {
			p += tokens[t]
			points--
			t--
		}
	}
	fmt.Println(points)
	return points
}

func sontask() {

}