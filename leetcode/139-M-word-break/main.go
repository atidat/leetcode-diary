package main

func judge(s string, dict map[string]bool) bool {
	lens := len(s)
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= lens; i++ {
		for j := 0; j < i; j++ {
			_, ok2 := dict[s[j:i]]

			if dp[j] && ok2 {
				dp[i] = true
				break
			}
		}
	}
	return dp[lens]
}

func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		dict[word] = true
	}
	if len(s) == 0 {
		if _, ok := dict[""]; ok {
			return true
		}
		return false
	}

	return judge(s, dict)
}

func main() {
	// s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	// wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
}
