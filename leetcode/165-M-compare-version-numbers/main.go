package main

import (
	"fmt"
	"strconv"
	"strings"
)

func minfunc(a, b int) (int, bool) {
	if a < b {
		return a, false
	}
	return b, true
}

func resfunc(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	}
	return -1
}

func compareVersion(version1 string, version2 string) int {
	v1Arr, v2Arr := strings.Split(version1, "."), strings.Split(version2, ".")
	fmt.Println(v1Arr)
	fmt.Println(v2Arr)

	lenMin, fstBig := minfunc(len(v1Arr), len(v2Arr))
	for i := 0; i < lenMin; i++ {
		v1int, _ := strconv.Atoi(v1Arr[i])
		v2int, _ := strconv.Atoi(v2Arr[i])

		res := resfunc(v1int, v2int)
		if 0 != res {
			return res
		}
	}

	restArr := v1Arr[lenMin:]
	if lenMin == len(v1Arr) {
		restArr = v2Arr[lenMin:]
	}

	fmt.Println(restArr)

	for i := 0; i < len(restArr); i++ {
		res, _ := strconv.Atoi(restArr[i])
		if res != 0 {
			if fstBig {
				return 1
			}
			return -1
		}
	}
	return 0
}

func main() {
	_ = compareVersion("1.0.1", "1")
}
