package main

func main() {

}

func maxNumberOfFamilies(n int, s [][]int) int {
	const s1 = 0b0111111110
	const s2 = 0b0000011110
	const s3 = 0b0111100000
	const s4 = 0b0001111000

	var res int
	dict := make(map[int]int, 0)
	for i := 0; i < len(s); i++ {
		if v, ok := dict[s[i][0]]; !ok {
			dict[s[i][0]] = handle(0, s[i][1])
		} else {
			dict[s[i][0]] = handle(v, s[i][1])
		}
	}

	for _, v := range dict {
		if v & s1 == 0 {

			res += 2
		} else if v & s2 == 0 {

			res++
		} else if v & s3 == 0 {

			res++
		} else if v & s4 == 0 {

			res++
		}
	}

	res += 2 * (n-len(dict))
	return res
}

func handle(src int, p int) int {
	if p == 1 || p == 10 {
		return src
	}
	base := 0x0000000001
	for i := 0; i < 10-p; i++ {
		base = base << 1
	}
	return src | base
}