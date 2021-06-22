package main

func isHappy(n int) bool {
	rec := make(map[int]bool, 1)
	rec[n] = true

	for {
		new := calc(n)
		if new == 1 {
			return true
		}
		if _, ok := rec[new]; ok {
			return false
		} else {
			rec[new] = true
		}
	}
	return false
}

func calc(n int) int {
	res := 0
	for n > 0 {
		last := n % 10
		res += last * last
		n = n / 10
	}
	return res
}

func main() {

}
