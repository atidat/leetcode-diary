package main

import "testing"

func TestFirst(t *testing.T) {
	if "IPv4" != validIPAddress("172.16.254.1") {
		t.Error("=============== ERROR 1 ===============")
	}
}

func TestSix(t *testing.T) {
	if "Neither" != validIPAddress("072.16.254.1") {
		t.Error("=============== ERROR 6 ===============")
	}
}

func TestSecond(t *testing.T) {
	if "IPv6" != validIPAddress("2001:0db8:85a3:0:0:8A2E:0370:7334") {
		t.Error("=============== ERROR 2 ===============")
	}
}

func TestSeven(t *testing.T) {
	if "IPv6" != validIPAddress("2001:0db8:85a3:0000:0:8A2E:0370:733a") {
		t.Error("=============== ERROR 7 ===============")
	}
}


func TestThree(t *testing.T) {
	if "Neither" != validIPAddress("256.256.256.256") {
		t.Error("=============== ERROR 3 ===============")
	}
}

func TestFour(t *testing.T) {
	if "Neither" != validIPAddress("2001:0db8:85a3:0:0:8A2E:0370:7334:") {
		t.Error("=============== ERROR 4 ===============")
	}
}

func TestFive(t *testing.T) {
	if "Neither" != validIPAddress("1e1.4.5.6") {
		t.Error("=============== ERROR 5 ===============")
	}
}

func TestEight(t *testing.T) {
	if "Neither" != validIPAddress("20EE:FGb8:85a3:0:0:8A2E:0370:7334") {
		t.Error("=============== ERROR 8 ===============")
	}
}