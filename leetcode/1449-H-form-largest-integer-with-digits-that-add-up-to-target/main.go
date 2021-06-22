package main

func main() {
	cost := []int{4,3,2,5,6,7,2,5,5}
	target := 9
	largestNumber(cost, target)
}

type st struct {
	v int // value
	l int // length
}

func largestNumber(cost []int, tar int) string {
	dp := make([]st, tar+1)
	dp[0] = st{0, 0}
	for i := 1; i <= tar; i++ {
		dp[i] = st{-1, -1}
	}

	for i := 1; i <= tar; i++ {
		for j := 0; j < 9; j++ {
			s := i-cost[j]
			if s >= 0 && dp[s].l != -1 && dp[s].l+1 >= dp[i].l {
				dp[i] = st{j+1, dp[s].l+1}
			}
		}
	}

	if dp[tar].l == -1 {
		return "0"
	}

	var res string
	for tar > 0 {
		res += string(dp[tar].v+'0')
		tar -= cost[dp[tar].v-1]
	}
	return res
}

/*func largestNumber(cost []int, tar int) string {
	dp := make([]string, tar+1)
	dp[0] = ""
	for i := 1; i <= tar; i++ {
		dp[i] = "0"
		for j := 0; j < 9; j++ {
			if i-cost[j] >= 0 && dp[i-cost[j]] != "0" &&
				len(dp[i]) <= 1+len(dp[i-cost[j]]) {
				dp[i] = string('1'+j) + dp[i-cost[j]]
			}
		}
	}
	return dp[tar]
}*/

/*
func largestNumber(cost []int, target int) string {
	rec := make(map[int]string)
	rec[0] = ""

	detail(cost, target, rec)
	if rec[target] == "*" {
		return "0"
	}
	return rec[target]
}

func detail(cost []int, tar int, rec map[int]string) string {
	if v, o := rec[tar]; o {
		return v
	}
	s, t := "", ""
	for i := 0; i < 9; i++ {
		if cost[i] <= tar {
			t = detail(cost, tar-cost[i], rec)
			if t != "*" {
				t = string('1' + i) + t
				if compare(t, s) {
					s = t
				}
			}
		}
	}
	if s == "" {
		s = "*"
	}
	rec[tar] = s
	return rec[tar]
}

func compare(a, b string) bool {
	if len(a) == len(b) {
		return a>=b
	}
	return len(a)>len(b)
}
*/