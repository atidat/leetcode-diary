package main

import (
	"fmt"
	"strconv"
	"strings"
)
func main() {
	fmt.Println(validIPAddress("2001:0db8:85a3:0000:0:8A2E:0370:733a"))
}

func validIPAddress(IP string) string {
	if strings.Contains(IP, ".") {
		return ipv4Judge(IP)
	}
	if strings.Contains(IP, ":") {
		return ipv6Judge(IP)
	}
	return "Neither"
}

/*
  ipv4:
	反向规则
	1、前导0不允许
	2、双0也不允许
*/
func ipv4Judge(IP string) string {
	nums := strings.Split(IP, ".")
	if len(nums) != 4 {
		return "Neither"
	}
	for _, v := range nums {
		if strings.HasPrefix(v, "0") && len(v) != 1 {
			return "Neither"
		}
		vn, err := strconv.Atoi(v)
		if err != nil {
			return "Neither"
		}
		if vn < 0 || vn > 255 {
			return "Neither"
		}
	}
	return "IPv4"
}

/*
  ipv6:
	1、
*/
func ipv6Judge(IP string) string {
	nums := strings.Split(IP, ":")
	if len(nums) != 8 {
		return "Neither"
	}
	for _, v := range nums {
		vLen := len(v)
		if vLen < 1 || vLen > 4 {
			return "Neither"
		}
		for _, e := range v {
			if e >= '0' && e <= '9' {
				continue
			} else if e >= 'a' && e <= 'f' {
				continue
			} else if e >= 'A' && e <= 'F' {
				continue
			}
			return "Neither"
		}
	}
	return "IPv6"
}