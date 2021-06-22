package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	smallestGoodBase("13")
}

/*
class Solution:
    def smallestGoodBase(self, n: str) -> str:
        num = int(n)
        def check(x, m):
            ans = 0
            for _ in range(m+1):
                ans = ans*x + 1
            return ans
        ans = float("inf")
        for i in range(1, 64):
            l = 2
            r = num
            while l < r:
                mid = l + (r - l)//2
                tmp = check(mid, i)
                if tmp == num:
                    ans = min(ans, mid)
                    break
                elif tmp < num:
                    l = mid + 1
                else:
                    r = mid

        return str(ans)
*/
func smallestGoodBase(s string) string {
	//要求所求进制 >= 2，所以范围为 [2, n]
	tar, _ := strconv.Atoi(s)
	fmt.Println("in: ", tar)
	res := math.MaxInt64

	for b := 2; b <= 64; b++ {
		l, r := 2, tar
		for l < r {
			mid := l+(r-l)/2
			fmt.Println(mid, b)
			tmp := iterate(b, mid)
			if tmp == tar {
				fmt.Println(res, mid)
				res = mmin(res, mid)
				break
			} else if tmp < tar {
				l = mid + 1
			} else {
				r = mid
			}
		}
	}
	fmt.Println(res)
	return strconv.Itoa(res)
}

func iterate(times int, base int) int {
	var res int
	con := 1
	for i := 0; i < times; i++ {
		res += con
		con *= base
	}
	return res

}
 func mmin(a, b int) int {
 	if a < b { return a }
 	return b
 }