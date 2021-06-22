package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mctFromLeafValues([]int{6,2}))
	fmt.Println(mctFromLeafValues([]int{6,2,4}))
	fmt.Println(mctFromLeafValues([]int{6,2,4, 8}))
	fmt.Println(mctFromLeafValues([]int{8,10,6,5,3,2,4,5,13,15,13,15,4,3,11,1,14,9,9,4}))
}

func mctFromLeafValues(arr []int) int {
	m := make(map[string]st)
	return executor(arr, m).sum
}

type st struct {
	mv int
	sum int
}

func executor(a []int, m map[string]st) st {
	if get(a, m) != INV { return get(a, m)}
	if len(a) == 1 {set(a, a[0], 0, m); return st{a[0], 0}}
	if len(a) == 2 {set(a, mmax(a[0], a[1]), a[0]*a[1], m); return st{mmax(a[0], a[1]), a[0] * a[1]}}

	res := math.MaxInt32
	mv := math.MinInt32
	for i := 1; i < len(a); i++ {
		l := executor(a[:i], m)
		r := executor(a[i:], m)
		res = mmin(res, l.mv*r.mv + l.sum + r.sum)
		mv = mmax(mv, mmax(l.mv, r.mv))
	}
	set(a, mv, res, m)
	return st{mv, res}
}

func mmax(a, b int) int {
	if a < b {return b}
	return a
}

func mmin(a, b int) int {
	if a < b {return a}
	return b
}

func set(a []int, im, is int, m map[string]st) {
	p := pre(a)
	if _, ok := m[p]; !ok {
		m[p] = st{im, is}
	}
}

var INV st = st{-1, -1}

func get(a []int, m map[string]st) st {
	if v, ok := m[pre(a)]; ok {
		return v
	}
	return INV
}

func pre(a []int) string {
	b := make([]byte, 0)
	for i := range a {
		if a[i] < 10 {
			b = append(b, '0'+byte(a[i]))
		} else {
			b = append(b, 'a'+byte(a[i]-10))
		}
	}
	return string(b)
}