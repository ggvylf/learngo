package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	res := Add(2, 3)
	if res != 5 {

		t.Errorf("res not 5")
	}

}

func BenchmarkSum(b *testing.B) {
	res := Sum(5)
	if res != 25 {

		b.Errorf("Sum is not 25\n")
	}
}
