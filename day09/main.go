package day09

import (
	"math"

	"github.com/jscaltreto/aoc-2022/lib"
)

type Pos [2]float64

var dirs map[byte]Pos = map[byte]Pos{
	'R': {1, 0},
	'L': {-1, 0},
	'D': {0, 1},
	'U': {0, -1},
}

type Rope struct {
	knots   []Pos
	visited map[Pos]struct{}
}

func (r *Rope) MoveHead(m Pos) {
	r.knots[0][0] += m[0]
	r.knots[0][1] += m[1]

	for knot := range r.knots[1:] {
		r.MoveKnot(knot, knot+1)
	}
	tail := r.knots[len(r.knots)-1]
	r.visited[Pos{tail[0], tail[1]}] = struct{}{}
}

func (r *Rope) MoveKnot(idx_l, idx_f int) {
	leader, follower := r.knots[idx_l], r.knots[idx_f]
	dx := leader[0] - follower[0]
	dy := leader[1] - follower[1]
	dist := math.Sqrt(math.Pow(math.Abs(dx), 2) + math.Pow(math.Abs(dy), 2))
	if math.Floor(dist) > 1 {
		if dx > 0 {
			follower[0] += 1
		} else if dx < 0 {
			follower[0] -= 1
		}
		if dy > 0 {
			follower[1] += 1
		} else if dy < 0 {
			follower[1] -= 1
		}
		r.knots[idx_f] = follower
	}
}

func SimulateRope(filename string, length int) int {
	rope := &Rope{
		knots:   make([]Pos, length),
		visited: map[Pos]struct{}{{0, 0}: {}},
	}
	for _, m := range lib.SlurpFile(filename) {
		for i := 1; i <= lib.StrToInt(m[2:]); i++ {
			rope.MoveHead(dirs[m[0]])
		}
	}
	return len(rope.visited)
}

func PartA(filename string) int {
	return SimulateRope(filename, 2)
}

func PartB(filename string) int {
	return SimulateRope(filename, 10)
}
