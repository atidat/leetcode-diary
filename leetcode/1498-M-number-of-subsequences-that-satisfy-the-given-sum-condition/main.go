package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := []string{"FS", "HOFS"}
	times := 3
	for i := 0; i < times; i++ {
		fmt.Printf("before: %v %p \n", arr, &arr)
		filter(&arr, "FS")
		fmt.Printf("after: %v %p \n", arr, &arr)
		fmt.Println()
	}

}

func numSubseq(nums []int, target int) int {

}