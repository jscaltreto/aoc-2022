package day05

import (
	"testing"
)

func TestPartA(t *testing.T) {
	result := PartA("data/test")
	exp := "CMZ"
	if result != exp {
		t.Fatalf("Result should be %s! Got %s", exp, result)
	}
	t.Log(PartA("data/input"))
}

func TestPartB(t *testing.T) {
	result := PartB("data/test")
	exp := "MCD"
	if result != exp {
		t.Fatalf("Result should be %s! Got %s", exp, result)
	}
	t.Log(PartB("data/input"))
}
