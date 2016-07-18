package main

import (
	"testing"
)

func TestEvenFibSum(t *testing.T) {
	res := EvenFibSum(35)
	if res != 44 {
		t.Error("evenFibSum(35) expected 44 but got", res)
	}
}
