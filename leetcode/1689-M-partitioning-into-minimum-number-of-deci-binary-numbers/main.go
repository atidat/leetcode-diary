package main

func main() {

}

/*
	1 <= n.length <= 105
*/
func minPartitions(n string) int {
	var res int
	for i := range n {
		res = mmax(res, int(n[i] - '0'))
	}
	return res
}

func mmax(a, b int) int {
	if a < b {
		return b
	}
	return a
}