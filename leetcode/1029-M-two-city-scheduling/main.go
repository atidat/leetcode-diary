package main

import "fmt"

func main() {

}

func twoCitySchedCost(c [][]int) int {
	var aMax, bMax int
	diff := make([]int, 0)
	for i := 0; i < len(c); i++ {
		aMax += c[i][0]
		bMax += c[i][1]
		diff = append(diff, c[i][0] - c[i][1])
	}
	fmt.Println("aMax, bMax", aMax, bMax)
	for i := 0; i < len(diff)-1; i++ {
		for j := i; j < len(diff); j++ {
			if diff[i] < diff[j] {
				diff[i], diff[j] = diff[j], diff[i]
			}
		}
	}

	for i, j := 0, len(diff)-1; i <= j; {
		aMax -= diff[i]
		bMax += diff[j]
		i++; j--
	}

	fmt.Println("aMax, bMax", aMax, bMax)
	if aMax > bMax {
		return bMax
	}
	return aMax
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return 0 - a
}

func samePosNeg(a, b int) bool {
	if a < 0 && b < 0 || a > 0 && b > 0 {
		return true
	}
	return false
}