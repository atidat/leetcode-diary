package main

import "testing"

func Test1(t *testing.T) {
	validStrs := []string{"32.e-80123", ".1", "0", "2", "0089", "-0.1", "+3.14",
		"4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1",
		"53.5e93", "-123.456e789"}
	for _, v := range validStrs {
		if !isNumber(v) {
			t.Errorf("=============== ERROR 1 ==> %s ===============", v)
		}
	}
}

func Test2(t *testing.T) {
	invalidStrs := []string{"6+1", ".0e", ".", "e", "abc", "1a", "1e", "e3", "99e2.5",
		"--6", "-+3", "95a54e53"}
	for _, v := range invalidStrs {
		if isNumber(v) {
			t.Errorf("=============== ERROR 2 ==> %s ===============", v)
		}
	}
}
