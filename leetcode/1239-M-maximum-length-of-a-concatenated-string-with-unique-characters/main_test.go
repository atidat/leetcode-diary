package main

import "testing"

func Test1(t *testing.T) {
	arr := []string{"un","iq","ue"}
	if 4 != maxLength(arr) {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	arr := []string{"cha","r","act","ers"}
	if 6 != maxLength(arr) {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	arr := []string{"abcdefghijklmnopqrstuvwxyz"}
	if 26 != maxLength(arr) {
		t.Error("=============== ERROR 3===============")
	}
}

func Test4(t *testing.T) {
	arr := []string{"a", "abc", "d", "de", "def"}
	if 6 != maxLength(arr) {
		t.Error("=============== ERROR 4===============")
	}
}

func Test5(t *testing.T) {
	arr := []string{"zboesm","rmtsyhhbpiezsxrbxcj","iufwwhifynsor",
		"xikybtblzpfkjaivifmpm","rpfhy","kfxmanukdmyhxht",
		"wsotfdrvjupavrlyrpit","gtleqgdbatyaq"}
	if 11 != maxLength(arr) {
		t.Error("=============== ERROR 5===============")
	}
}