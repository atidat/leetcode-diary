package main

func ha() 
{

}
func main() {
    _ = brokenCalc(2,3)
}

/*
在显示着数字的坏计算器上，我们可以执行以下两种操作：

双倍（Double）：将显示屏上的数字乘 2；
递减（Decrement）：将显示屏上的数字减 1 。
最初，计算器显示数字 X。

返回显示数字 Y 所需的最小操作数。
*/
func brokenCalc(ini int, tar int) int {
	// 每次操作都有两个选择
	// 操作1：双倍
	// 操作2：递减
	res := recur(ini, tar, 0)
	return res
}

func recur(ini, tar int, cnt int) int {
	if ini >= tar {
		return cnt + ini - tar
	}
	if ini == tar >> 1 && ini != (tar - 1) >> 1 {
		return cnt + 1
	}
	return recur(ini, tar + 1, cnt + 1)
}

func cmpmin(a, b int) int {
	if a > b {
		return b
	}
	return a
}
