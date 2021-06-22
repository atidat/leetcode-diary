package main

import "fmt"

func main() {

}

/*
  1 <= maxChoosableInteger <= 20
  0 <= desiredTotal <= 300
*/
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger >= desiredTotal {
		return true
	}
	if maxChoosableInteger * (maxChoosableInteger+1) < 2*desiredTotal {
		return false
	}
	fmt.Println(maxChoosableInteger, desiredTotal)

	cache := make(map[int]bool)
	return exec(0, maxChoosableInteger, desiredTotal, cache)
}

func exec(bits int, mci int, dt int, cache map[int]bool) bool {
	if _, ok := cache[bits]; ok {
		return cache[bits]
	}

	for i := 0; i < mci; i++ {
		if bits & (1 << i) > 0 {
			continue
		}
		// dp[n] 为 false，表示 先手负，为 true，表示 先手胜
		if i + 1 >= dt || !exec(bits | (1 << i), mci, dt-i-1, cache) {
			cache[bits] = true
			return true
		}
	}
	cache[bits] = false
	return false
}

