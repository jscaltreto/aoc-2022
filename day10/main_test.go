package day10

import (
	"testing"
)

func TestPartA(t *testing.T) {
	result := PartA("data/test")
	exp := 13140
	if result != exp {
		t.Fatalf("Result should be %d! Got %d", exp, result)
	}
	t.Log(PartA("data/input"))
}

func TestPartB(t *testing.T) {
	result := PartB("data/test")
	exp := 1
	if result != exp {
		t.Fatalf("Result should be %d! Got %d", exp, result)
	}
	t.Log(PartB("data/input"))
}
