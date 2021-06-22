package main

import "fmt"

func main() {
	comb := Constructor("abc", 2)
	fmt.Println(comb.Next())
	fmt.Println(comb.HasNext())
	fmt.Println(comb.Next())
	fmt.Println(comb.HasNext())
	fmt.Println(comb.Next())
	fmt.Println(comb.HasNext())
}

type CombinationIterator struct {
	loc []int
	s string
	c int
}


func Constructor(characters string, combinationLength int) CombinationIterator {
	return CombinationIterator{
		loc: make([]int, 0),
		s: characters,
		c: combinationLength,
	}
}


func (this *CombinationIterator) Next() string {
	if len(this.loc) == 0 {
		for i := 0; i < this.c; i++ {
			this.loc = append(this.loc, i)
		}
		return this.s[:this.c]
	}

	for i := 0; i <= len(this.loc)-1; i++ {
		if this.loc[len(this.loc)-1-i] == len(this.s)-1-i {
			continue
		}
		//fmt.Println()
		//fmt.Println(this.loc)
		this.loc[len(this.loc)-1-i] += 1
		for j := len(this.loc)-i; j < len(this.loc); j++ {
			this.loc[j] = this.loc[j-1]+1
		}
		break
		//fmt.Println(this.loc)
	}
	tmp := make([]byte, 0)
	for i :=0; i < len(this.loc); i++ {
		tmp = append(tmp, this.s[this.loc[i]])
	}
	//fmt.Println(this.loc)
	return string(tmp)
}

func (this *CombinationIterator) HasNext() bool {
	curs := make([]byte, 0)
	for i := 0; i < len(this.loc); i++ {
		curs = append(curs, this.s[this.loc[i]])
	}
	return string(curs) != this.s[len(this.s)-this.c:]
}


/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */