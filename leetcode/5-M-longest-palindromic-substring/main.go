package main

import "fmt"

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	left, scope := 0, 0
	for i := 0; i < len(s); i++ {
		//fmt.Println(dp)
		//fmt.Println("=========")
		for j := 0; j < i; j++ {
			dp[j][i] = s[i] == s[j] && (i-1 <= j || dp[j+1][i-1])
			if dp[j][i] && scope < i-j+1 {
				left, scope = j, i-j+1
			}
		}
	}
	return s[left : left+scope]
}

func main() {
	fmt.Println(longestPalindrome("cbbd"))
	//fmt.Println(longestPalindrome("babad"))
	//fmt.Println(longestPalindrome("asdlbvdaasdfghjkkjhgfdsabgtuhbnjhgghjuyttyujhgzxswplbhfds"))
}
