package main

import "fmt"

func main() {
	restoreArray([][]int{{2,1},{3,4},{3,2}})
	restoreArray([][]int{{4,-2},{1,4},{-3,1}})
	restoreArray([][]int{{100000,-100000}})
	restoreArray([][]int{{4,-10},{-1,3},{4,-3}, {-3,3}})
}

type st struct {
	one int
	two int
}

func restoreArray(adjacentPairs [][]int) []int {
	m, u := record(&adjacentPairs)

	res := make([]int, 0, len(adjacentPairs)+1)
	res = append(res, u)
	for len(res) != len(adjacentPairs)+1 {
		o, t := m[u].one, m[u].two
		if o != 100001 {
			res = append(res, o)
			if v, ok := m[o]; ok {
				if v.one == u {
					m[o] = st{100001, v.two}
				} else if v.two == u {
					m[o] = st{v.one, 100001}
				}
			}
			u = o
		}
		if t != 100001 {
			res = append(res, t)
			if v, ok := m[t]; ok {
				if v.one == u {
					m[t] = st{100001, v.two}
				} else if v.two == u {
					m[t] = st{v.one, 100001}
				}
			}
			u = t
		}
	}
	fmt.Println(res)
	return res
}

func record(adj *[][]int) (map[int]st, int) {
	m := make(map[int]st, 0)
	var uni int
	uni_dic := make(map[int]bool, 0)
	for i := 0; i < len(*adj); i++ {
		if v, ok := m[(*adj)[i][0]]; !ok {
			m[(*adj)[i][0]] = st{(*adj)[i][1], 100001}
		} else {
			m[(*adj)[i][0]] = st{v.one, (*adj)[i][1]}
		}

		if v, ok := m[(*adj)[i][1]]; !ok {
			m[(*adj)[i][1]] = st{(*adj)[i][0], 100001}
		} else {
			m[(*adj)[i][1]] = st{v.one, (*adj)[i][0]}
		}

		if _, ok := uni_dic[(*adj)[i][0]]; ok {
			delete(uni_dic, (*adj)[i][0])
		} else {
			uni_dic[(*adj)[i][0]] = true
		}
		if _, ok := uni_dic[(*adj)[i][1]]; ok {
			delete(uni_dic, (*adj)[i][1])
		} else {
			uni_dic[(*adj)[i][1]] = true
		}

	}

	for k, _ := range uni_dic {
		uni = k
		break
	}
	return m, uni
}