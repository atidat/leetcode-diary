package main

import "testing"

func TestFirst(t *testing.T) {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	if true == wordBreak(s, wordDict) {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestTwo(t *testing.T) {
	s := "applepenapple"
	wordDict := []string{"apple", "pen"}
	if true == wordBreak(s, wordDict) {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestThree(t *testing.T) {
	s := "catsandog"
	wordDict := []string{"cats", "dog", "and", "cat"}
	if false == wordBreak(s, wordDict) {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestFour(t *testing.T) {
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	if false == wordBreak(s, wordDict) {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestFive(t *testing.T) {
	s := ""
	wordDict := []string{"", "asdf"}
	if true == wordBreak(s, wordDict) {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}

func TestSix(t *testing.T) {
	s := ""
	wordDict := []string{"ad", "asdf"}
	if false == wordBreak(s, wordDict) {
		t.Log("=============== SUCCESS ===============")
	} else {
		t.Error("=============== ERROR ===============")
	}
}
