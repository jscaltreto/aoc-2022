package day10

import (
	"fmt"
	"strings"

	"github.com/jscaltreto/aoc-2022/lib"
)

func PartA(filename string) int {
	X := 1
	result := 0
	cycle := 1
	nextCycle := 20
	for _, line := range lib.SlurpFile(filename) {
		inst := strings.Split(line, " ")
		nextX := X
		switch inst[0] {
		case "noop":
			cycle++
		case "addx":
			cycle += 2
			nextX += lib.StrToInt(inst[1])
		}
		if cycle >= nextCycle {
			if cycle == nextCycle {
				result += nextX * nextCycle
			} else {
				result += X * nextCycle
			}
			nextCycle += 40
		}
		X = nextX
	}

	return result
}

func PartB(filename string) int {
	X := 1
	Y := 0
	insts := lib.SlurpFile(filename)
	progCtr := 0

	for cycle := 0; cycle < 240; cycle++ {
		scan := cycle % 40
		if scan == 0 {
			fmt.Print("\n")
		}
		if scan >= X - 1 && scan <= X + 1 {
			fmt.Print("█")
		} else {
			fmt.Print(" ")
		}
		inst := strings.Split(insts[progCtr], " ")
		if inst[0] == "addx" {
			if Y == 0 {
				progCtr--
				Y = lib.StrToInt(inst[1])
			} else {
				X += Y
				Y = 0
			}
		}
		progCtr++
	}
	fmt.Println("\n---")
	return 1
}
