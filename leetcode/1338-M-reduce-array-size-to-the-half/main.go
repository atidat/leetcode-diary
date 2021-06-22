package main

import (
	"sort"
)

func main() {

}

func minSetSize(arr []int) int {
	// cnt2num: 计数 -> 该计数对应有几个不同数字
	// 如 [1,1,2,3]
	//   次数2：对应1（只有1出现2次）
	//   次数1：对应2（分别为数字2和3）
	// num2cnt: 每一个数字 -> 计数
	// 如 [1,1,2,3]
	//   数字1：对应2（出现2次）
	//   数字2：对应1
	//   数字3：对应1
	// 所以 cnt2num 和 num2cnt 是互为映射关系
	//fmt.Println(arr)
	cnt2num, num2cnt := make(map[int]int, 0), make(map[int]int, 0)
	for k := range arr {
		if v, ok := num2cnt[arr[k]]; !ok {
			num2cnt[arr[k]] = 1

			if v1, ok1 := cnt2num[1]; !ok1 {
				cnt2num[1] = 1
			} else {
				cnt2num[1] = v1 + 1
			}
		} else {
			num2cnt[arr[k]] = v + 1

			if v1, ok1 := cnt2num[v+1]; !ok1 {
				cnt2num[v+1] = 1
			} else {
				cnt2num[v+1] = v1 + 1
			}
			cnt2num[v] = cnt2num[v] - 1
			if cnt2num[v] == 0 {
				delete(cnt2num, v)
			}
		}
	}
	// {3,3,3,3,5,5,5,2,2,7}
	//fmt.Println(num2cnt)
	//fmt.Println(cnt2num)

	type per struct {
		cnt int
		nums int
	}
	res := make([]per, 0)
	for k, v := range cnt2num {
		res = append(res, per{k, v})
	}
	sort.Slice(res, func(a, b int) bool {
		return res[a].cnt > res[b].cnt
	})
	//fmt.Println(res)

	min := 0
	var cnt int
	for k := 0; k < len(res); k++ {
		for j := 1; j <= res[k].nums; j++ {
			min += res[k].cnt
			cnt++
			// fmt.Println(min, len(arr)/2)
			if min >= len(arr)/2 {
	//			fmt.Println("1", cnt)
				return cnt
			}
		}
	}
	//fmt.Println(cnt)
	return cnt
}