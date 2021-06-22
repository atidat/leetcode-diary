package main

import (
	"strings"
)

func main() {

}

func minAddToMakeValid(S string) int {
	newS := S
	for {
		newS = strings.Replace(S, "()", "", -1)
		if newS == S {
			break
		}
		S = newS
	}
	return len(newS)
}