package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
/*	values, labels := []int{3,0,3,0,6}, []int{0,2,1,1,0}
	num_wanted, use_limit := 4, 1
	_ = largestValsFromLabels(values, labels, num_wanted, use_limit)*/
	s := "hello"
	fmt.Printf("main: %p\n", &s)
	hello(&s)
	fmt.Println(s)
}

func hello(s *string) {
	fmt.Printf("func: %p\n", &s)
	*s = "world"
	time.Sleep(1)

}
/*func largestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {

	rec := make(map[int]int, 0)
	res := 0
	for i := 0; i < len(values); i++ {
		res = mmax(res, dfs(&values, &labels, rec, num_wanted, use_limit, -1))
	}
	fmt.Println(res)
	return res
}

func mmax(a, b int) int {
	if a < b { return b }
	return a
}

func dfs(val, lab *[]int, rec map[int]int, rest, limit, ind int) int {
	if rest == 0 {
		return 0
	}

	m := 0
	for i := ind+1; i < len(*val); i++ {
		t, ok := rec[(*lab)[i]]
		if ok && t >= limit {
			continue
		}

		if !ok {
			rec[(*lab)[i]] = 1
		} else {
			rec[(*lab)[i]] += 1
		}
		m = mmax(m, (*val)[i] + dfs(val, lab, rec, rest-1, limit, i))
		rec[(*lab)[i]] -= 1
	}
	return m
}*/

/*func largestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {
	//fmt.Println(values, use_limit)
	labelRevs := classify(&values, &labels)
	labelUses := initLabelCnt(&labels)
	//fmt.Println(labelRevs, labelUses)

	var res int
	for num_wanted > 0 {
		res += chutouniao(labelRevs, labelUses, use_limit)
		num_wanted--
	}
	return res
}

// 筛选属于每个标签的元素，并从大到小排序
func classify(values, labels *[]int) map[int][]int {
	labelRevs := make(map[int][]int, 0)

	// 筛选
	for lab_k, lab_v := range *labels {
		if vals, ok := labelRevs[lab_v]; !ok {
			labelRevs[lab_v] = []int{(*values)[lab_k]}
		} else {
			labelRevs[lab_v] =append(vals, (*values)[lab_k])
		}
	}

	// 排序
	for _, v := range labelRevs {
		sort.Slice(v, func(i, j int) bool {
			return v[i] > v[j]
		})
	}
	return labelRevs
}

func initLabelCnt(labels *[]int) map[int]int {
	labelUseCnt := make(map[int]int, 0)
	for _, v := range *labels {
		labelUseCnt[v] = 0
	}
	return labelUseCnt
}

// 选出每个标签对应的元素组中最大的那个，并从中删除
func chutouniao(labelRevs map[int][]int, labelUses map[int]int, limit int) (res int) {
	elites := make([]int, 0)
	m, k := math.MinInt32, -1

	for lab, vals := range labelRevs {
		cnt, ok := labelUses[lab]
		if len(vals) < 1 || ok && cnt >= limit {
			continue
		}

		if m < vals[0] {
			m = vals[0]
			k = lab
		}
		elites = append(elites, vals[0])
	}

	if len(elites) == 0 {
		return 0
	}

	sort.Slice(elites, func(i, j int) bool {
		return elites[i] > elites[j]
	})
	//fmt.Println(elites)

	if len(labelRevs[k]) > 0 {
		labelRevs[k] = labelRevs[k][1:]
	} else {
		labelRevs[k] = []int{}
	}


	labelUses[k] += 1
	if labelUses[k] >= limit {
		delete(labelUses, k)
		delete(labelRevs, k)
	}
	return elites[0]
}
*/

type st struct {
	val int
	lab int
}

func largestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {
	labelCnt := make(map[int]int, 0)
	for _, v := range labels {
		if _, ok := labelCnt[v]; !ok {
			labelCnt[v] = 0
		}
	}

	ll := len(values)

	mixes := make([]st, ll, ll)
	for i := 0; i < ll; i++ {
		mixes[i] = st{values[i], labels[i]}
	}
	sort.Slice(mixes, func(i, j int) bool {
		return mixes[i].val > mixes[j].val
	})

	var res int
	for i := 0; i < ll; i++ {
		if labelCnt[mixes[i].lab] >= use_limit {
			continue
		}

		if num_wanted <= 0 {
			break
		}

		res += mixes[i].val
		labelCnt[mixes[i].lab] += 1
		num_wanted--
	}
	return res
}

func mmax(a, b int) int {
	if a > b { return a }
	return b
}