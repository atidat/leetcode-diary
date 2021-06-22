package main

import "fmt"

func main() {

}

type queue struct {
	buffer []int
	l int
}

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	s2b := site2bus(routes)
	q := &queue{
		buffer: []int{source},
		l: 1,
	}
	hisBus := make([]int, len(routes))
	var busNum int

	for !q.empty() {
		fmt.Println("q is not empty")
		busNum++
		for i := q.len() - 1; i >= 0; i-- {
			site := q.pop()
			busesNo := s2b[site]
			for _, bn := range busesNo {
				if hisBus[bn] == 1 { continue }
				q.push(routes[bn])
				for _, s := range routes[bn] {
					if target == s {
						fmt.Println(busNum)
						return busNum
					}
				}
				hisBus[bn] = 1
			}
		}
	}
	fmt.Println(busNum)
	return -1
}

func (q *queue)push(sites []int) {
	q.l += len(sites)
	q.buffer = append(q.buffer, sites...)
}

func (q *queue)pop() int {
	q.l -= 1
	tmp := q.buffer[0]
	q.buffer = q.buffer[1:]
	return tmp
}

func (q *queue)empty() bool {
	if q.l <= 0 { return true }
	return false
}

func (q *queue)len() int {
	return q.l
}

func site2bus(r [][]int) map[int][]int {
	s2b := make(map[int][]int)
	for busNo, _ := range r {
		for _, site := range r[busNo] {
			if _, ok := s2b[site]; ok {
				s2b[site] = append(s2b[site], busNo)
			} else {
				s2b[site] = []int{busNo}
			}
		}
	}
	return s2b
}