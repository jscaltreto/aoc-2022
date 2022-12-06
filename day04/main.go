package day04

import (
	"strings"

	"github.com/jscaltreto/aoc-2022/lib"
)

func ParseElf(sec string) (min, max [2]int) {
	els := strings.Split(sec, ",")
	for i, el := range els {
		secs := strings.Split(el, "-")
		min[i], max[i] = lib.StrToInt(secs[0]), lib.StrToInt(secs[1])
	}
	return
}

func PartA(filename string) (overlap int) {
	for _, sec := range lib.SlurpFile(filename) {
		min, max := ParseElf(sec)
		if (min[0] >= min[1] && max[0] <= max[1]) || (min[1] >= min[0] && max[1] <= max[0]) {
			overlap++
		}
	}
	return
}

func PartB(filename string) (overlap int) {
	for _, sec := range lib.SlurpFile(filename) {
		min, max := ParseElf(sec)
		if !(min[1] > max[0] || max[1] < min[0]) {
			overlap++
		}
	}
	return
}
