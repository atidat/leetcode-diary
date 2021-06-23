package main

import "fmt"

/*func phoneDict() map[byte][]byte {
	phone := make(map[byte][]byte, 0)
	phone['2'] = []byte{'a', 'b', 'c'}
	phone['3'] = []byte{'d', 'e', 'f'}
	phone['4'] = []byte{'g', 'h', 'i'}
	phone['5'] = []byte{'j', 'k', 'l'}
	phone['6'] = []byte{'m', 'n', 'o'}
	phone['7'] = []byte{'p', 'q', 'r', 's'}
	phone['8'] = []byte{'t', 'u', 'v'}
	phone['9'] = []byte{'w', 'x', 'y', 'z'}
	return phone
}
func iterator(digits string, i int, p map[byte][]byte, repo *[]byte, dict *[]string) {
	if i == len(digits) {
		*dict = append(*dict, string(*repo))
		return
	}
	for j := 0; j < len(p[digits[i]]); j++ {
		*repo = append(*repo, p[digits[i]][j])
		iterator(digits, i+1, p, repo, dict)
		*repo = (*repo)[0:len(*repo)-1]
	}
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	p := phoneDict()
	dict := make([]string, 0)
	tmp := make([]byte, 0)
	iterator(digits, 0, p, &tmp, &dict)
	return dict
}

func main() {
	fmt.Println(letterCombinations("4576"))
}
*/

func phoneDict() map[byte][]byte {
	pd := make(map[byte][]byte, 0)
	pd['2'] = []byte{'a', 'b', 'c'}
	pd['3'] = []byte{'d', 'e', 'f'}
	pd['4'] = []byte{'g', 'h', 'i'}
	pd['5'] = []byte{'j', 'k', 'l'}
	pd['6'] = []byte{'m', 'n', 'o'}
	pd['7'] = []byte{'p', 'q', 'r', 's'}
	pd['8'] = []byte{'t', 'u', 'v'}
	pd['9'] = []byte{'w', 'x', 'y', 'z'}
	return pd
}



func letterCombinations(digits string) []string {
	pd := phoneDict()

	var res []string
	dl := len(digits)
	if dl == 0 { return res }

	var dfs func(int, int, []byte)
	dfs = func(tar int, ind int, tmp []byte) {
		if tar == ind {
			ha := make([]byte, tar)
			copy(ha, tmp)
			res = append(res, string(tmp))
			return
		}

		for i := 0; i < len(pd[digits[ind]]); i++ {
			tmp = append(tmp, pd[digits[ind]][i])
			dfs(tar, ind+1, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(dl, 0, []byte{})
	fmt.Println(res)
	return res
}