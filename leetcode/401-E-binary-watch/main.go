package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(readBinaryWatch(5))
/*	for i := 1; i < 8; i++ {
		fmt.Println(readBinaryWatch(i))
	}*/
}
/*
var pow = []int{1,2,4,8,16,32}

func readBinaryWatch(turnedOn int) []string {
	hourMap := make(map[int][]int)
	minuteMap := make(map[int][]int)

	hourMap[1] = []int{1,2,4,8}
	minuteMap[1] = []int{1,2,4,8,16,32}


	var res []string
	for hc := 0; hc <= turnedOn; hc++ {
		mc := turnedOn-hc
		huse := make([]int, 4, 4)
		muse := make([]int, 6, 6)
		hmem, mmem := make([]int, 0), make([]int, 0)
		dfs(hc, 0, &huse, 0, &hmem)
		dfs(mc, 0, &muse, 0, &mmem)

		sort.Ints(hmem)
		sort.Ints(mmem)
		hmem = adjust(hmem, 12)
		mmem = adjust(mmem, 60)
		fmt.Println(hmem, mmem)
		res = append(res, gather(hmem, mmem)...)
	}
	return res
}

func adjust(mem []int, tar int) []int {
	var k int
	for k = 0; k < len(mem); k++ {
		if mem[k] >= tar {
			break
		}
	}
	return mem[:k]
}

// t表示1的个数; c表示已使用1的个数；res表示哪些位置上被占用;
func dfs(t int, ind int, res *[]int, val int, mem *[]int) {
	if t <= 0 {
		*mem = append(*mem, val)
		return
	}

	for i := ind; i < len(*res); i++ {
		if (*res)[i] == 1 {
			continue
		}
		(*res)[i] = 1
		dfs(t-1, i+1, res, val+pow[i], mem)
		(*res)[i] = 0
	}
}

func gather(huse, muse []int) []string {
	var res []string

	handel := func(v int, isM bool) (t []byte) {
		if v < 10 {
			if !isM {
				t = append(t, byte(v + '0'))
			} else {
				t = append(t, '0')
				t = append(t, byte(v + '0'))
			}
		} else {
			t = append(t, byte(v/10 + '0'))
			t = append(t, byte(v%10 + '0'))
		}
		return
	}

	for _, hv := range huse {
		for _, mv := range muse {
			var t []byte
			t = handel(hv, false)
			t = append(t, ':')
			t = append(t, handel(mv, true)...)
			res = append(res, string(t))
		}
	}
	return res
}
*/

var (
	res []string

/*	hourPath int
	minPath int*/

	// 可选择的数据。每次选择的时间为：hours[i]+minutes[i]
	hours = []int{1,2,4,8,0,0,0,0,0,0}
	minutes = []int{0,0,0,0,1,2,4,8,16,32}
)

// 解体思路：
// 相当于有2行数组：第一行是：4个元素的数组，表示时钟；第二行是：6个元素的数组表示分钟。
// 那么，解题其实就变化成了，在这个数组中，找到几个数的组合，并且组合是合法的（也即正确的时钟表示）。
//
// 优化1 - 可以将2行数组优化成一维数组。

func readBinaryWatch(num int) []string {
	var res []string
	for h := uint8(0); h < 12; h++ {
		for m := uint(0); m < 60; m++ {
			if bits.OnesCount8(h) + bits.OnesCount(m) == num {
				res = append(res, fmt.Sprintf("%d:%.2d", h, m))
			}
		}
	}
	return res
}