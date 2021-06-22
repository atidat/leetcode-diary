package main

func main() {

}

type st struct {
	num int // 存放该项作为最长链的数量
	cur int // 存放该项作为最长链的索引
	next int // 存放该项作为最长链的前一项的索引
}

func findLongestChain(pairs [][]int) int {
	if len(pairs) == 1 { return 1 }

/*	// dp memory array
	dp := make([]st, len(pairs))
	for i := range dp {
		dp[i] = st{1, i, -1}
	}
*/

	// resort array
	// [0]小的在前面；[0]相同时，[1]小的在前面
	resort(pairs)

	res := 1

	// mm 记录每一个链的链尾
	mm := make(map[st]bool, 0)
	mm[st{1, len(pairs)-1, 1}] = true

	for i := len(pairs)-2; i >= 0; i-- {
		tmp := st{1, i, -1}
		del := st{1, i, -1}
		for k, _ := range mm {
			if pairs[i][1] < pairs[k.cur][0] && tmp.num < k.num+1 {
				tmp = st{k.num+1, i, k.cur}
				del = st{k.num, k.cur, k.num}
				res = mmax(res, k.num+1)
			}
		}
		delete(mm, del)
		mm[tmp] = true
	}
	return res
}

func resort(pairs [][]int) {
	for i := 0; i < len(pairs)-1; i++ {
		for j := i+1; j < len(pairs); j++ {
			if pairs[i][0] > pairs[j][0] ||
				pairs[i][0] == pairs[j][0] && pairs[i][1] > pairs[j][1] {
				pairs[i], pairs[j] = pairs[j], pairs[i]
			}
		}
	}
}

func mmax(a, b int) int {
	if a < b { return b }
	return a
}