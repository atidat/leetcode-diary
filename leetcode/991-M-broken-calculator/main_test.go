package main

import (
	"fmt"
	"testing"
)
/*
func TestFirst(t *testing.T) {
	fmt.Println(brokenCalc(2,3))
	if 2 != brokenCalc(2,3) {
		t.Error("=============== ERROR 1===============")
	}
}

func TestSecond(t *testing.T) {
	fmt.Println(brokenCalc(5, 8))
	if 2 != brokenCalc(5, 8) {
		t.Error("=============== ERROR 2===============")
	}
}

func TestThird(t *testing.T) {
	fmt.Println(brokenCalc(3, 10))
	if 3 != brokenCalc(3,10) {
		t.Error("=============== ERROR 3===============")
	}
}

func TestFourth(t *testing.T) {
	fmt.Println(brokenCalc(1024, 1))
	if 1023 != brokenCalc(1024,1) {
		t.Error("=============== ERROR 4===============")
	}
}

func TestFifth(t *testing.T) {
	fmt.Println(brokenCalc(2, 1))
	if 1 != brokenCalc(2,1) {
		t.Error("=============== ERROR 5===============")
	}
}
*/
func BenchmarkHello(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("hello")
    }
}
