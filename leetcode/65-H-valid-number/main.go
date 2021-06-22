package main

import (
	"strings"
	//"fmt"
	"regexp"
)

func main() {

}

func isNumber(s string) bool {
	//fmt.Println("input is", s)
	dotNum := strings.Count(s, ".")
	if dotNum > 1 { return false }

	intPat := "^[+-]?[0-9]+([0-9]*|[Ee]{1}[+-]?[0-9]+)$"
	floatPat1 := "^[+-]?[0-9]+.[0-9]*(|[Ee]{1}[+-]?[0-9]+)$"
	floatPat2 := "^[+-]?.[0-9]+(|[Ee]{1}[+-]?[0-9]+)$"

	if dotNum == 1 {
		return floatJudge(s, floatPat1) || floatJudge(s, floatPat2)
	}
	return intJudge(s, intPat)
}

func intJudge(s string, intPat string) bool {
	flag, err := regexp.Match(intPat, []byte(s))
	if err != nil || !flag {
		return false
	}
	return true
}

func floatJudge(s string, floatPat string) bool {
	flag, err := regexp.Match(floatPat, []byte(s))
	if err != nil || !flag {
		return false
	}
	return true
}