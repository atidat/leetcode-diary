package main

import "fmt"

func rangeBitwiseAnd(m int, n int) int {
	veri := 0x80000000
	rgtMov := 31
	res := 0

	for m&veri == n&veri {
		if m&veri == veri {
			res |= (1 << rgtMov)
		}
		rgtMov--
		m <<= 1
		n <<= 1
	}
	return res
}

func main() {
	m1 := 5
	n1 := 7
	fmt.Println(rangeBitwiseAnd(m1, n1))

	m2 := 1
	n2 := 2147483647
	fmt.Println(rangeBitwiseAnd(m2, n2))
}
