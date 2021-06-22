package main

import (
	"fmt"
)

func binBytes2Uint(s []byte) uint32 {
	var res uint32
	scope := 1
	for i := 0; i < len(s); i++ {
		res += uint32(scope * int(s[i]-'0'))
		scope *= 2
	}
	return res
}

func reverseBits(num uint32) uint32 {
	fmt.Println(num)
	oldBitStr := fmt.Sprintf("%b", num)
	//fmt.Println(oldBitStr)
	bitsBytes := []byte(oldBitStr)
	for i := len(bitsBytes); i < 32; i++ {
		bitsBytes = append([]byte{'0'}, bitsBytes...)
	}
	fmt.Println(string(bitsBytes))
	newNum := binBytes2Uint(bitsBytes)
	fmt.Println(newNum)
	return newNum
}

func main() {
	_ = reverseBits(43261596)
	_ = reverseBits(4294967293)
}
