package day15

import (
	"testing"
)

func TestPartA(t *testing.T) {
	result := PartA("data/test", 10)
	exp := 26
	if result != exp {
		t.Fatalf("Result should be %d! Got %d", exp, result)
	}
	t.Log(PartA("data/input", 2000000))
}

func TestPartB(t *testing.T) {
	result := PartB("data/test", 20)
	exp := 56000011
	if result != exp {
		t.Fatalf("Result should be %d! Got %d", exp, result)
	}
	t.Log(PartB("data/input", 4000000))
}
