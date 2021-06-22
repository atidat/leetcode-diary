package main

import "fmt"

func main() {

}
func pushDominoes(dominoes string) string {
	l2r := run(dominoes, 'L', 'R')
	r2l := runreverse(dominoes, 'R', 'L')
	return merge(l2r, r2l)
}

func run(d string, still, down byte) []byte {
	var s []byte
	var flag bool
	for i := range d {
		if d[i] == still {
			s = append(s, d[i])
			flag = false
		} else if d[i] == down {
			s = append(s, d[i])
			flag = true
		} else {
			if flag {
				s = append(s, down)
			} else {
				s = append(s, '.')
			}
		}
	}
	return s
}

func runreverse(d string, still, down byte) []byte {
	var s []byte
	var flag bool
	for i := len(d)-1; i >= 0; i-- {
		if d[i] == still {
			s = append(s, d[i])
			flag = false
		} else if d[i] == down {
			s = append(s, d[i])
			flag = true
		} else {
			if flag {
				s = append(s, down)
			} else {
				s = append(s, '.')
			}
		}
	}
	for i, j := 0, len(d)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++;j--
	}
	return s
}


func merge(l2r, r2l []byte) string {
	fmt.Println(string(l2r))
	fmt.Println(string(r2l))
	var s []byte
	var l int
	for i := 0; i < len(l2r); i++{
		//fmt.Printf("%c %c\n", l2r[i], r2l[i])
		if l2r[i] == r2l[i] {
			if l != 0 {
				j := 0
				for ; j < l/2; j++ {
					s = append(s, 'R')
				}
				j = l/2
				if l % 2 != 0 { s = append(s, '.'); j += 1 }
				for ; j < l; j++ {
					s = append(s, 'L')
				}
				l = 0
			}
			s = append(s, l2r[i])
		} else if l2r[i] == 'R' && r2l[i] == 'L' {
			l++
		} else {
			if l2r[i] != '.' {
				s = append(s, l2r[i])
			} else {
				s = append(s, r2l[i])
			}

		}
	}
	//fmt.Println(string(s))
	return string(s)
}