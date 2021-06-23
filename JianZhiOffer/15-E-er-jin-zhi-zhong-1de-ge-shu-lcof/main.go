package main

func main() {
	hammingWeight(11111111111111111111111111111101)
}


func hammingWeight(num uint32) int {
	var cnt int
	for num != 0 {
		cnt++
		num >>= 1
	}

	return cnt
}