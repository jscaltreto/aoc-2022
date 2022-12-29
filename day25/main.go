package day25

import (
	"fmt"
	"math"

	"github.com/jscaltreto/aoc-2022/lib"
)

func Stoi(s string) int {
	i := 0
	for p := 0; p < len(s); p++ {
		pl := int(math.Pow(5, float64(len(s)-p-1)))
		n := s[p]
		switch n {
		case '-':
			i -= pl
		case '=':
			i -= pl * 2
		default:
			i += int(n-'0') * pl
		}
	}
	return i
}

func Itos(i int) string {
	f := float64(i)
	l5 := math.Floor(math.Log(f) / math.Log(5))
	str := ""
	for l5 >= 0 {
		pow := math.Pow(5, l5)
		co := math.Round(f / pow)
		if math.Abs(co) > 2 {
			l5++
		} else {
			switch co {
			case -1:
				str += "-"
			case -2:
				str += "="
			default:
				str += fmt.Sprint(math.Abs(co))
			}
			f -= co * pow
			l5--
		}
	}
	return str
}

func PartA(filename string) string {
	sum := 0
	for _, s := range lib.SlurpFile(filename) {
		sum += Stoi(s)
	}
	return Itos(sum)
}
