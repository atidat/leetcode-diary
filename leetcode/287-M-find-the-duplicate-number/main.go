package main

import "fmt"

func main() {

}

func findDuplicate(nums []int) int {
	tmp := firstStep(nums, 1, len(nums) - 1)
	fmt.Printf("result is %d\n", tmp)
	return tmp
}

func firstStep(nums[] int, head, tail int) int {
	diff := 0
	mid := head + (tail - head)/2
	for _, v := range nums {
		if v <= mid {
			diff++
		}
	}

	if diff == len(nums) {
		return mid
	}
	fmt.Printf("drawerPick is %d\n", diff)
	return drawerPick(nums, head, tail, diff, len(nums))
}

func drawerPick(nums []int, head, tail int, headDIff, tailDIff int) int {

	if head + 1 == tail {
		fmt.Println(head, tail, headDIff, tailDIff)
		if tail - head < tailDIff - headDIff {
			return tail
		}
		return head

	}

	mid := head + (tail - head)/2
	diff := 0
	for _, v := range nums {
		if v <= mid {
			diff++
		}
	}

	/*fmt.Printf("index:[%d], diff:[%d]\n", tail, tailDIff)
	fmt.Printf("index:[%d], diff:[%d]\n", mid, diff)
	fmt.Printf("index:[%d], diff:[%d]\n", head, headDIff)*/

	if tail - mid < tailDIff - diff {
		return drawerPick(nums, mid, tail, diff, tailDIff)
	}

	hDIff := 0
	for _, v := range nums {
		if v <= head {
			hDIff++
		}
	}
	return drawerPick(nums, head, mid, hDIff, diff)
}


