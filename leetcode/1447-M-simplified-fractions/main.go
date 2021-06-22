package main

import "fmt"

func main() {
	fmt.Println(simplifiedFractions(6))
}

func simplifiedFractions(n int) []string {
	if n < 2 {
		return []string{}
	}

	repo := make([]string, 0)
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if canBeDivided(i, j) {
				tmp := fmt.Sprintf("%d/%d", j, i)
				repo = append(repo, tmp)
			}
		}
	}
	return repo
}

func canBeDivided(a, b int) bool {
	// 判断两数是否有公约数
	for {
		// c := a / b
		d := a % b
		if d == 0 && b == 1 {
			return true
		}
		if d != 0 {
			a = b
			b = d
		} else {
			return false
		}
	}
	return false
}
