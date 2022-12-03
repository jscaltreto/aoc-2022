package day03

import (
	"github.com/jscaltreto/aoc-2022/lib"
)

func Score(c rune) int {
	if c >= 97 {
		return int(c) - 96
	} else {
		return int(c) - (65 - 27)
	}
}

func PartA(filename string) int {
	var pts int
	for _, sack := range lib.SlurpFile(filename) {
		pts += PartACheck(sack)
	}
	return pts
}

func PartACheck(sack string) int {
	c1, c2 := sack[:len(sack)/2], sack[len(sack)/2:]
	c1m := map[rune]struct{}{}
	for _, c := range c1 {
		c1m[c] = struct{}{}
	}
	for _, c := range c2 {
		if _, ok := c1m[c]; ok {
			return Score(c)
		}
	}
	panic("No duplicate!")
}

func PartB(filename string) int {
	var pts int
	sacks := lib.SlurpFile(filename)
	for i := 0; i < len(sacks); i += 3 {
		pts += PartBCheck(sacks[i : i+3])
	}
	return pts
}

func PartBCheck(sacks []string) int {
	common := map[rune]int{}
	for si, sack := range sacks {
		for _, c := range sack {
			common[c] = common[c] | 1<<si
			if common[c] == 0b111 {
				return Score(c)
			}
		}
	}
	panic("No common item!")
}
