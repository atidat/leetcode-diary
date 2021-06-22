package main

func main() {

}

func canReach(arr []int, start int) bool {
	// rec[i] = 1 means cannot arrive at 0
	rec := make([]int, len(arr))
	if dfs(&arr, start+arr[start], &rec) || dfs(&arr, start-arr[start], &rec) {
		return true
	}
	return false
}

func dfs(arr *[]int, cur int, rec *[]int) bool {
	if cur >= len(*arr) || cur < 0 || (*rec)[cur] == 1 {
		return false
	}


	(*rec)[cur] = 1
	right := dfs(arr, cur+(*arr)[cur], rec)
	left := dfs(arr, cur-(*arr)[cur], rec)

	return (*arr)[cur] == 0 || right || left
}