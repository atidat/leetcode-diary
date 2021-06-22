package main

func main() {

}

func minCostToMoveChips(position []int) int {
	// odd 奇数 even 偶数
	oddNum, evenNum := 0, 0
	for i := range position {
		if position[i] % 2 != 0 {
			oddNum++
		} else {
			evenNum++
		}
	}
	if evenNum >= oddNum {
		return oddNum
	}
	return evenNum
}