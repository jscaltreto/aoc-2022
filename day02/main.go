package day02

import (
	"github.com/jscaltreto/aoc-2022/lib"
)

const (
	OPP_ROCK = "A"
	OPP_PAPER = "B"
	OPP_SCISSORS = "C"
	RES_ROCK = "X"
	RES_PAPER = "Y"
	RES_SCISSORS = "Z"
	LOSE = "X"
	DRAW = "Y"
	WIN = "Z"
)

var PTS map[string]int = map[string]int{
	RES_ROCK: 1,
	RES_PAPER: 2,
	RES_SCISSORS: 3,
	OPP_ROCK + RES_ROCK: 3,
	OPP_PAPER + RES_PAPER: 3,
	OPP_SCISSORS + RES_SCISSORS: 3,
	OPP_ROCK + RES_PAPER: 6,
	OPP_PAPER + RES_SCISSORS: 6,
	OPP_SCISSORS + RES_ROCK: 6,
}

func PartA(filename string) int {
	var pts int
	for _, turn := range lib.SlurpFile(filename) {
		opp, res :=  string(turn[0]), string(turn[2])
		pts += PTS[res] + PTS[opp + res]
	}
	return pts
}

var WIN_LOSE map[string]string = map[string]string{
	OPP_ROCK + WIN: RES_PAPER,
	OPP_ROCK + DRAW: RES_ROCK,
	OPP_ROCK + LOSE: RES_SCISSORS,
	OPP_SCISSORS + WIN: RES_ROCK,
	OPP_SCISSORS + DRAW: RES_SCISSORS,
	OPP_SCISSORS + LOSE: RES_PAPER,
	OPP_PAPER + WIN: RES_SCISSORS,
	OPP_PAPER + DRAW: RES_PAPER,
	OPP_PAPER + LOSE: RES_ROCK,
}

func PartB(filename string) int {
	var pts int
	for _, turn := range lib.SlurpFile(filename) {
		opp, outcome :=  string(turn[0]), string(turn[2])
		res := WIN_LOSE[opp + outcome]
		pts += PTS[res] + PTS[opp + res]
	}
	return pts
}
