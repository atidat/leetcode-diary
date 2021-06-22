package main

import "testing"

func TestFirst(t *testing.T) {
	d := ".L.R...LR..L.."
	o := "LL.RR.LLRRLL.."
	if o != pushDominoes(d) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	d := "RR.L"
	o := "RR.L"
	if o != pushDominoes(d) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	d := "RL"
	o := "RL"
	if o != pushDominoes(d) {
		t.Error("=============== ERROR 3===============")
	}
}