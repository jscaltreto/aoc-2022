package day06

import (
	"github.com/jscaltreto/aoc-2022/lib"
)

func CheckChars(data string, num int) int {
	for pos := num - 1; pos < len(data); pos++ {
		chrs := map[byte]struct{}{}
		for i := pos - (num - 1); i <= pos; i++ {
			chrs[data[i]] = struct{}{}
		}
		if len(chrs) == num {
			return pos + 1
		}
	}
	return 0
}

func PartA(filename string) int {
	return CheckChars(lib.SlurpFile(filename)[0], 4)
}

func PartB(filename string) (overlap int) {
	return CheckChars(lib.SlurpFile(filename)[0], 14)
}
