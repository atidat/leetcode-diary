package main

func main() {

}

func balancedStringSplit(s string) int {
	res, rcnt, lcnt := 0, 0, 0
	for i := range s {
		if s[i] == 'R' {
			rcnt++
		} else {
			lcnt++
		}
		if rcnt == lcnt {
			rcnt, lcnt = 0, 0
			res++
		}
	}
	return res
}