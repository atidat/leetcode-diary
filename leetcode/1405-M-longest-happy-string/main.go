package main

func main() {

}
func longestDiverseString(a int, b int, c int) string {
	dp := make([]byte, 3)
	bank := ready(a, b, c)
	dp[1] = (*bank)[0].str
	(*bank)[0].val--
	adjust(bank, (*bank)[0].str)
	dp[2] = (*bank)[0].str
	(*bank)[0].val--
	adjust(bank, (*bank)[0].str)
	//fmt.Println(dp, bank)

	aa, bb, cc := &((*bank)[0].val), &((*bank)[1].val), &((*bank)[2].val)
	for i := 3; *aa + *bb + *cc != getMaxFromThree(*aa, *bb, *cc); i++ {
		if (*bank)[0].str == dp[i-1] {
			if dp[i-1] == dp[i-2] {
				dp = append(dp, (*bank)[1].str)
				(*bank)[1].val--
				adjust(bank, (*bank)[1].str)
				//fmt.Println(dp)
				continue
			}
		}
		// 最大数字对应字符 ！= 末尾字符
		dp = append(dp, (*bank)[0].str)
		(*bank)[0].val--
		adjust(bank, (*bank)[0].str)
		//fmt.Println(dp)
	}


	dpl := len(dp)
	for i := 0; i < 2; i++ {
		if (dp[dpl-1] != (*bank)[0].str || dp[dpl-2] != (*bank)[0].str) && (*bank)[0].val > 0 {
			dp = append(dp, (*bank)[0].str)
			(*bank)[0].val--
		}
	}
	//fmt.Println(string(dp[1:]))
	return string(dp[1:])
}

type st struct {
	val int
	str byte
}
func ready(a, b, c int) *[]st {
	bank := make([]st, 3)
	bank[0] = st{a, 'a'}
	bank[1] = st{b, 'b'}
	bank[2] = st{c, 'c'}
	// sort from big -> small
	for i := 0; i < 3; i++ {
		for j := i+1; j < 3; j++ {
			if bank[i].val < bank[j].val {
				bank[i], bank[j] = bank[j], bank[i]
			}
		}
	}
	return &bank
}
func getMaxFromThree(a, b, c int) int {
	return mmax(a, mmax(b, c))
}
func mmax(a, b int) int {
	if a < b { return b }
	return a
}
func adjust(bank *[]st, tar byte) {
	flag := true
	for i := 0; i < 3; i++ {
		if flag && (*bank)[i].str != tar {
			continue
		}
		flag = false
		for j := i+1; j < 3; j++ {
			if (*bank)[i].val < (*bank)[j].val {
				(*bank)[i], (*bank)[j] = (*bank)[j], (*bank)[i]
			}
		}
	}
}