package main

import (
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	res := longestPalindrome("babad")
	if "bab" == res || "aba" == res {
		t.Logf("[babad] ==> Success")
	} else {
		t.Errorf("[babad] ==> Failed")
	}
}

func TestTwo(t *testing.T) {
	res := longestPalindrome("cbbd")
	fmt.Println(res)
	if "bb" == res {
		t.Logf("[cbbd] ==> Success")
	} else {
		t.Errorf("[cbbd] ==> Failed")
	}
}

func TestThree(t *testing.T) {
	res := longestPalindrome("asdlbvdaasdfghjkkjhgfdsabgtuhbnjhgghjuyttyujhgzxswplbhfds")
	fmt.Println(res)
	if "asdfghjkkjhgfdsa" == res {
		t.Logf("[asdlbvdaasdfghjkkjhgfdsabgtuhbnjhgghjuyttyujhgzxswplbhfds] ==> Success")
	} else {
		t.Errorf("[asdlbvdaasdfghjkkjhgfdsabgtuhbnjhgghjuyttyujhgzxswplbhfds] ==> Failed")
	}
}

func TestFourth(t *testing.T) {
	res := longestPalindrome("")
	if "" == res {
		t.Logf("[] ==> Success")
	} else {
		t.Errorf("[] ==> Failed")
	}
}
