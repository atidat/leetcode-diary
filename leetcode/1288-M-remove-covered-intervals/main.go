package main

import (
	"sort"
)

func main() {

}

func removeCoveredIntervals(intervals [][]int) int {
	// 先排序，排序标准
	// 第一个数字小在前；第二个数字大在后
	res := make([][]int, 0)
	//fmt.Println(intervals)  // TODO DELETE
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0] ||
			(intervals[a][0] == intervals[b][0] &&
				intervals[a][1] > intervals[b][1])
	})
	//fmt.Println(intervals)  // TODO DELETE

	for i:= len(intervals)-1; i >= 0; i-- {
		for j := i-1; j >= 0; j-- {
			if intervals[i][0] == intervals[j][0] {
				res = append(res, intervals[i])
				break
			}
			if intervals[i][1] <= intervals[j][1] {
				res = append(res, intervals[i])
				break
			}
		}
	}
	//fmt.Println(res)

	return len(intervals) - len(res)
}