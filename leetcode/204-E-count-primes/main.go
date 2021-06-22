package main

import "math"

func isPrimeNumber(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); {
		if n%i == 0 {
			return false
		}
		i++
	}

	return true
}

func countPrimes(n int) int {
	count := 1
	if n <= 2 {
		return 0
	}
	for i := 3; i < n; {
		if true == isPrimeNumber(i) {
			count++
		}
		i += 2
	}
	return count
}

func main() {
}
