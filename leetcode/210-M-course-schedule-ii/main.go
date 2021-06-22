package main

import "math/rand"

func main() {

}

type st struct {
	d int
	s int
}

func findOrder(numCourses int, pre [][]int) []int {

	dirs := make(map[st]bool)
	starts := make([]int, numCourses, numCourses)
	dests := make([]int, numCourses, numCourses)
	for _, v := range pre {
		dirs[st{v[0], v[1]}] = true
		starts[v[1]] = 1
		dests[v[0]] = 1
	}

	// find start point
	sp := make([]int, numCourses, numCourses)
	for i := 0; i < len(pre); i++ {
		
	}
}