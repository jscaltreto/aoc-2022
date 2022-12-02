package day01

import (
	"github.com/jscaltreto/aoc-2022/lib"
)

func PartA(filename string) int {
	var cals, maxCals int
	for _, calsStr := range lib.SlurpFile(filename) {
		if calsStr == "" {
			maxCals = CheckPartA(cals, maxCals)
			cals = 0
		} else {
			cals += lib.StrToInt(calsStr)
		}
	}
	return CheckPartA(cals, maxCals)
}

func CheckPartA(cals, maxCals int) int {
	if maxCals < cals {
		return cals
	}
	return maxCals
}

func PartB(filename string) int {
	topThree := make([]int, 3)
	var cals int
	for _, calsStr := range lib.SlurpFile(filename) {
		if calsStr == "" {
			topThree = CheckPartB(cals, topThree)
			cals = 0
		} else {
			cals += lib.StrToInt(calsStr)
		}
	}
	topThree = CheckPartB(cals, topThree)
	return topThree[0] + topThree[1] + topThree[2]
}

func CheckPartB(cals int, topThree []int) []int {
	for idx, top := range topThree {
		if cals > top {
			nextTop := make([]int, idx)
			copy(nextTop, topThree)
			nextTop = append(nextTop, cals)
			nextTop = append(nextTop, topThree[idx:]...)
			return nextTop[:3]
		}
	}
	return topThree
}
