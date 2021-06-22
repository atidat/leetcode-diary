package main

import "testing"

func Test1(t *testing.T) {
	if 3 != minPartitions("32") {
		t.Error("=============== ERROR 1===============")
	}
}

func Test2(t *testing.T) {
	if 8 != minPartitions("82734") {
		t.Error("=============== ERROR 2===============")
	}
}

func Test3(t *testing.T) {
	if 9 != minPartitions("27346209830709182346") {
		t.Error("=============== ERROR 3===============")
	}
}