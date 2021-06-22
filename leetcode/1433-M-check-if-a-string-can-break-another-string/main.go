package main

import (
	"sort"
)

func main() {

}

func checkIfCanBreak(s1 string, s2 string) bool {
	b1, b2 := []byte(s1), []byte(s2)
	if len(b1) != len(b2) {
		return false
	}

	sort.Slice(b1, func(a, b int) bool {
		return b1[a] < b1[b]
	})
	sort.Slice(b2, func(a, b int) bool {
		return b2[a] < b2[b]
	})
	var pos bool
	var cnt int
	for i := 0; i < len(b1); i++ {
		if b1[i] != b2[i] {
			pos = b1[i] < b2[i]
			cnt = i
			break
		}
	}
	for i := cnt+1; i < len(b1); i++ {
		if pos && b1[i] > b2[i] {
			return false
		}
		if !pos && b1[i] < b2[i] {
			return false
		}
	}
	return true
}