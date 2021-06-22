package main

func main() {

}

func countArrangement(n int) int {
	vis := make([]bool, n+1)
	var res int
	backtrack(&vis, 1, &res)
	return res
}

func backtrack(v *[]bool, ind int, res *int) {
	if ind == len(*v) {
		*res += 1
		return
	}

	for i := 1; i < len(*v); i++ {
		if (*v)[i] == true {
			continue
		}
		if ind % i != 0 && i % ind != 0 {
			continue
		}

		(*v)[i] = true
		backtrack(v, ind+1, res)
		(*v)[i] = false
	}
}