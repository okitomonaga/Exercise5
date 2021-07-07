package main

import (
	//"reflect"
	//"fmt"
	"testing"
)

func Test_CalcPlus1(t *testing.T) {
	result := calcPlus(3, 6)
	if result != 9 {
		t.Error("Test_CalcPlus1")
	}
}
func Test_CalcPlus2(t *testing.T) {
	result := calcPlus(3, 2)
	if result != 5 {
		t.Error("Test_CalcPlus1")
	}
}
