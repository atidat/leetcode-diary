package main

import "fmt"

func main() {

}

func maxTurbulenceSize(a []int) int {
	al := len(a)
/*	if al < 3 { return al }*/

	na := make([]int, al+2)
	na[0], na[al-1] = 0, 0
	copy(na[1:al+1], a)
	fmt.Println(a)

	dp := make([]int, al)
	for i := 1; i <= al; i++ {
		if (na[i-1] - na[i]) * (na[i] - na[i+1]) < 0 {
			dp[i-1] = 1
		}
	}

	res := 1
	for i := 0; i < al; {
		if dp[i] == 1 {
			tmp := 1
			j := i+1
			for ; j < al && dp[j] == 1; j++ {
				tmp++
			}
			if i != 0 && j != al {
				tmp += 2
			} else if i == 0 && j == al {
				//
			} else {
				tmp += 1
			}
			res = mmax(res, tmp)
			i = j
		} else {
			i++
		}
	}

	fmt.Println(dp, res)
	return res
}

func mmax(a, b int) int {
	if a < b { return b}
	return a
}