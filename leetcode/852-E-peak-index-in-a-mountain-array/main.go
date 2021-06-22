package main

func main() {

}

/*
3 <= arr.length <= 104
0 <= arr[i] <= 106
题目数据保证 arr 是一个山脉数组
 

进阶：很容易想到时间复杂度 O(n) 的解决方案，你可以设计一个 O(log(n)) 的解决方案吗？
*/
func peakIndexInMountainArray(arr []int) int {
	h, t := 0, len(arr)-1
	m := (h+t)/2
	for ; m >= h; {
		if (arr[m]-arr[m-1]) * (arr[m]-arr[m+1]) > 0 {
			return m
		}
		if arr[m] > arr[m-1] { // 上升
			h = m+1
		} else { // 下降
			t = m
		}
		m = (h+t)/2
	}
	return -1
}