package main

func main() {

}

type board struct {
	stop bool
	on []int
	off []int
}

func carPooling(trips [][]int, capacity int) bool {
	sch := make([]board, 1001)
	for _, v := range trips {
		var tmp1 []int
		if len(sch[v[1]].on) == 0 {
			tmp1 = []int{v[0]}
		} else {
			tmp1 = append(sch[v[1]].on, v[0])
		}
		var tmp2 []int
		if len(sch[v[2]].off) == 0 {
			tmp2 = []int{v[0]}
		} else {
			tmp2 = append(sch[v[2]].off, v[0])
		}

		sch[v[1]] = board{true, tmp1, sch[v[1]].off}
		sch[v[2]] = board{true, sch[v[2]].on, tmp2}
	}
	//fmt.Println(sch)
	for _, v := range sch {
		if !v.stop {
			continue
		}
		capacity += calc(v.off)
		capacity -= calc(v.on)
		if capacity < 0 {
			break
		}
	}
	if capacity >= 0 {
		return true
	}
	return false
}

func calc(a []int) int {
	res := 0
	for _, v := range a {
		res += v
	}
	return res
}
