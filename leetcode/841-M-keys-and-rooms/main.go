package main

import "fmt"

func main() {

}
func canVisitAllRooms(rooms [][]int) bool {
	vis := make([]bool, len(rooms))
	vis[0] = true
	dfs(&rooms, &vis, 0)
	fmt.Println(vis)
	for i := 0; i < len(vis); i++ {
		if !vis[i] {
			return false
		}
	}
	return true
}

func dfs(r *[][]int, v *[]bool, i int) {
	for ind := 0; ind < len((*r)[i]); ind++ {
		if (*v)[(*r)[i][ind]] {
			continue
		}
		(*v)[(*r)[i][ind]] = true
		dfs(r, v, (*r)[i][ind])
	}
}