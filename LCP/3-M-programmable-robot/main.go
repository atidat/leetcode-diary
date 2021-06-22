package main

import (
	"strings"
)

func main() {

}

func robot(command string, obstacles [][]int, x int, y int) bool {
	if !pointOnPath(command, x, y) {
		return false
	}
	for _, v := range obstacles {
		if len(v) == 0 {continue}
		if x+y > v[0]+v[1] && pointOnPath(command, v[0], v[1]) {
			return false
		}
	}

	return true
}

func pointOnPath(command string, x, y int) bool {
	tx := strings.Count(command, "R") * ((x+y)/len(command)) +
		strings.Count(command[:(x+y)%len(command)], "R")
	ty := strings.Count(command, "U") * ((x+y)/len(command)) +
		strings.Count(command[:(x+y)%len(command)], "U")
	return x == tx && y == ty

}