package main

func main() {

}

func minDeletionSize(a []string) int {
	if len(a) == 0 || len(a[0]) == 0 {
		return 0
	}
	res := 0
	stLen := len(a[0])
	aLen := len(a)
	for i := 0; i < stLen; i++ {
		for j := 0; j < aLen-1; j++ {
			if a[j][i] > a[j+1][i] {
				res++
				break
			}
		}
	}
//	fmt.Println(res)
	return res
}