package day25

import (
	"testing"
)

func TestPartA(t *testing.T) {
	result := PartA("data/test")
	exp := "2=-1=0"
	if result != exp {
		t.Fatalf("Result should be %s! Got %s", exp, result)
	}
	t.Log(PartA("data/input"))
}
