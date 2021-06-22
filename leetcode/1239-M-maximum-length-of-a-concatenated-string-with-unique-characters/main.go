package main

import "fmt"

func main() {
	arr := []string{"a", "abc", "d", "de", "def"}
	maxLength(arr)
}

/*
func maxLength(arr []string) (ans int) {
    masks := []int{}
outer:
    for _, s := range arr {
        mask := 0
        for _, ch := range s {
            ch -= 'a'
            if mask>>ch&1 > 0 {
			// 若 mask 已有 ch，则说明 s 含有重复字母，无法构成可行解
                continue outer
            }
            mask |= 1 << ch // 将 ch 加入 mask 中
        }
        masks = append(masks, mask)
    }

    var backtrack func(int, int)
    backtrack = func(pos, mask int) {
        if pos == len(masks) {
            ans = max(ans, bits.OnesCount(uint(mask)))
            return
        }
        if mask&masks[pos] == 0 { // mask 和 masks[pos] 无公共元素
            backtrack(pos+1, mask|masks[pos])
        }
        backtrack(pos+1, mask)
    }
    backtrack(0, 0)
    return
}
*/

func maxLength(arr []string) int {
	bucket := make(map[int32]bool)
	var res int
	for i := 0; i < len(arr); i++ {
		res = mmax(res, dfs(arr, i, bucket))
		bucket = map[int32]bool{}
	}

	fmt.Println(res)
	return res
}

func dfs(arr []string, ind int, buc map[int32]bool) int {
	var res int
	for i := ind; i < len(arr); i++ {
		flag := loop(arr[i], buc)
		if flag {
			continue
		}
		res = mmax(res, len(arr[i]) + dfs(arr, i+1, buc))
		pop(arr[i], buc)
	}
	return res
}

func loop(s string, buc map[int32]bool) bool {
	for k, v := range s {
		if _, ok := buc[v]; ok {
			pop(s[:k], buc)
			return true
		}
		buc[v] = true
	}
	return false
}

func pop(s string, buc map[int32]bool) {
	for _, v := range s {
		delete(buc, v)
	}
}

func mmax(a, b int) int {
	if a < b { return b }
	return a
}