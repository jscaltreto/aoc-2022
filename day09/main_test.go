package day09

import (
	"testing"
)

func TestPartA(t *testing.T) {
	result := PartA("data/testa")
	exp := 13
	if result != exp {
		t.Fatalf("Result should be %d! Got %d", exp, result)
	}
	t.Log(PartA("data/input"))
}

func TestPartB(t *testing.T) {
	result := PartB("data/testb")
	exp := 36
	if result != exp {
		t.Fatalf("Result should be %d! Got %d", exp, result)
	}
	t.Log(PartB("data/input"))
}
