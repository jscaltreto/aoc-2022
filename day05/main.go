package day05

import (
	"strings"

	"github.com/jscaltreto/aoc-2022/lib"
)

func MoveCrates(data []string, moveFunc func([]byte, int) []byte ) string {
	cursor := 0
	stacks := [][]byte{}
	for ; cursor < len(data); cursor++ {
		if data[cursor+1] == "" {
			break
		}
		line := data[cursor]
		for stack := 0; true; stack++ {
			lc := 1 + (stack * 4)
			if lc >= len(line) {
				break
			}
			if stack >= len(stacks) {
				stacks = append(stacks, []byte{})
			}
			if line[lc] != ' ' {
				stacks[stack] = append(stacks[stack], line[lc])
			}
		}
	}

	for cursor += 2; cursor < len(data); cursor++ {
		line := strings.Split(data[cursor], " ")
		num := lib.StrToInt(string(line[1]))
		src := lib.StrToInt(string(line[3])) - 1
		dst := lib.StrToInt(string(line[5])) - 1
		nextStack := moveFunc(stacks[src], num)
		nextStack = append(nextStack, stacks[dst]...)
		stacks[dst] = nextStack
		stacks[src] = stacks[src][num:]
	}

	overlap := ""
	for _, stack := range stacks {
		overlap += string(stack[0])
	}
	return overlap
}

func PartA(filename string) string {
	return MoveCrates(lib.SlurpFile(filename), func(src []byte, num int) []byte {
		nextStack := []byte{}
		for i := num - 1; i >= 0; i-- {
			nextStack = append(nextStack, src[i])
		}
		return nextStack
	})
}

func PartB(filename string) (overlap string) {
	return MoveCrates(lib.SlurpFile(filename), func(src []byte, num int) []byte {
		nextStack := make([]byte, num)
		copy(nextStack, src)
		return nextStack
	})
}
