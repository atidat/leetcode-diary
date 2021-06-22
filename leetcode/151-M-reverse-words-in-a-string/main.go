package main

import (
	"fmt"
)

func main() {
	reverseWords("the sky is blue")
	reverseWords("  hello world!  ")
	reverseWords(" ")
}

func reverseWords(s string) string {
	if len(s) == 0 {
        return s
	}
	
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			s = s[i:]
			break
		}
	}
	res := ""
	if s[0] == ' ' {
		return ""
	}
	fmt.Println(s)

	rec := false
	head, tail := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if rec == false {
			if s[i] == ' ' {
				continue
			} else {
				rec = true
				tail = i
			}
		} else {
			if s[i] == ' ' {
				rec = false
				head = i + 1
				res += s[head : tail+1]
				res += " "
				//fmt.Println(res)
			} else {
				continue
			}
		}
	}

	res += s[0 : tail+1]
	fmt.Println(res)
	return res
}
