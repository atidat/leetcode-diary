package main

func calcContinuous(nums []byte) (head, tail int) {
	head, tail = -1, -1
	for k := range nums {
		if nums[k] == '1' {
			if head == -1 {
				head = k
				continue
			}
		}
	}
	return head, tail
}

func maximalSquare(matrix [][]byte) int {
	repos := make([][]int, len(matrix))
	for k := range repos {
		repos[k] = make([]int, 2)
		repos[k][0], repos[k][1] = calcContinuous(matrix[k])
	}

}

func main() {
}
