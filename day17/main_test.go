package day17

import (
	"testing"
)

func TestPartA(t *testing.T) {
	result := PartA("data/test")
	exp := 3068
	if result != exp {
		t.Fatalf("Result should be %d! Got %d", exp, result)
	}
	t.Log(PartA("data/input"))
}

func TestPartB(t *testing.T) {
	result := PartB("data/test")
	exp := 1514285714288
	if result != exp {
		t.Fatalf("Result should be %d! Got %d", exp, result)
	}
	t.Log(PartB("data/input"))
}
