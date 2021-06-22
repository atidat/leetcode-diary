package main

import "fmt"

func main() {
	graph := [][]int{{2},{3}, {1},{}}
	allPathsSourceTarget(graph)
}

/*func allPathsSourceTarget(graph [][]int) [][]int {
	tarNum := len(graph) - 1

	m := make(map[int][][]int, 0)
	m[tarNum] = [][]int{{tarNum}}
	for j := tarNum-1; j >= 0; j-- {
		exec(graph[j], j, m)
	}
	fmt.Println(m[0])
	return m[0]
}


func exec(gra []int, t int, m map[int][][]int) {
	m[t] = make([][]int, 0)
	for i := range gra {
		if v, ok := m[gra[i]]; ok {
			for j := range v {
				tmp := make([]int, len(v[j])+1)
				tmp[0] = t
				copy(tmp[1:], v[j])
				m[t] = append(m[t], tmp)
			}
		}
	}
	if v, ok := m[t]; !ok || len(v) == 0 {
		delete(m, t)
	}
}*/

// graph := [][]int{{2},{3}, {1},{}}
func allPathsSourceTarget(graph [][]int) [][]int {
	//fmt.Println(graph)
	tar := len(graph) - 1
	m := make([][]int, 0)
	for i := range graph[0] {
		hist := []int{0}
		exec(graph, graph[0][i], tar, hist, &m)
	}
	fmt.Println(m)
	return m
}

func exec(g [][]int, ind, tar int, hist []int, m *[][]int) {
	if ind == tar {
		tmp := make([]int, 1+len(hist))
		copy(tmp[:len(hist)], hist)
		tmp[len(hist)] = tar
		*m = append(*m, tmp)
		return
	}
	for i := range g[ind] {
		tmp := make([]int, 1+len(hist))
		copy(tmp[:len(hist)], hist)
		tmp[len(hist)] = ind
		exec(g, g[ind][i], tar, tmp, m)

	}
}
