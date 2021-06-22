package main

func main() {

}

/*
	1 <= words.length <= 1000
	1 <= words[i].length <= 16
	words[i] 仅由小写英文字母组成。
*/
func longestStrChain(words []string) int {
	resort(words)
	var res int
	dp := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		dp[i] = 1
		for j := 1; j <= i; j++ {
			if suit(words[i-j], words[i]) {
				dp[i] = dp[i-j] + 1
			}
		}
		res = mmax(res, dp[i])
	}
	return res
}

func mmax(a, b int) int {
	if a < b { return b }
	return a
}

func suit(a, b string) bool {
	if len(a) + 1 != len(b) {
		return false
	}
	for i := 0; i < len(b); i++ {
		tmp := b[0:i] + b[i+1:]
		if tmp == a {
			return true
		}
	}
	return false
}

func resort(w []string) {
	for i := 0; i < len(w)-1; i++ {
		for j := i+1; j < len(w); j++ {
			if len(w[i]) > len(w[j]) {
				w[i], w[j] = w[j], w[i]
			}
		}
	}
}