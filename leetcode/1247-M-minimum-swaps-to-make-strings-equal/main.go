package main

import "fmt"

func main() {

}

func minimumSwap(s1 string, s2 string) int {
	re1, re2 := []byte(s1), []byte(s2)
	//fmt.Printf("%c\n", re1)
	//fmt.Printf("%c\n", re2)
	//fmt.Println("====")
	var cnt int
	var exchange bool = false
	for i := 0; i < len(re1); i++ {
		if re1[i] == re2[i] {
			continue
		}
		var NotfindSame bool = true
		for j := i+1; j < len(re2); j++ {
			if re1[j] == re2[j] {
				continue
			}
			if re2[i] == re2[j] {
				re1[i], re2[j] = re2[j], re1[i]
				cnt++
				NotfindSame = false
				exchange = false
				break
			}
		}
		if NotfindSame && !exchange {
			// 没交换过且没找到相同的
			re1[i], re2[i] = re2[i], re1[i]
			cnt++
			exchange = true
			i--
		} else if exchange {
			// 交换过都没找到相同的
			return -1
		}
	//	fmt.Printf("%c\n", re1)
	//	fmt.Printf("%c\n", re2)
	//	fmt.Println()
	}
	fmt.Println(cnt)
	return cnt
}
