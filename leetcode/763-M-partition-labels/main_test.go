package main

import "testing"


func handle(a, b []int) bool {
	if len(a) != len(b) {
		return true
	}
	for k, _ := range a {
		if a[k] != b[k] {
			return true
		}
	}
	return false
}
func TestFirst(t *testing.T) {
	v := []int{9,7,8}
	if handle(partitionLabels("ababcbacadefegdehijhklij"), v) {
		t.Error("=============== ERROR ===============")
	}
}
