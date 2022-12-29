package day23

import (
	"math"

	"github.com/jscaltreto/aoc-2022/lib"
)

const (
	Y int = iota
	X
)

var maxCoord = Coord{math.MaxInt, math.MaxInt}

type Coord [2]int

func (c Coord) N() Coord { return Coord{c[Y] - 1, c[X]} }
func (c Coord) S() Coord { return Coord{c[Y] + 1, c[X]} }
func (c Coord) E() Coord { return Coord{c[Y], c[X] + 1} }
func (c Coord) W() Coord { return Coord{c[Y], c[X] - 1} }

type Elf struct {
	pos Coord
}

func (el *Elf) GetNext(m *Map) Coord {
	p := el.pos
	n, s, e, w := p.N(), p.S(), p.E(), p.W()
	ne, nw, se, sw := n.E(), n.W(), s.E(), s.W()
	if m.AllEmpty(n, s, e, w, ne, nw, se, sw) {
		return p
	}
	check := [][]Coord{
		{n, ne, nw},
		{s, se, sw},
		{w, nw, sw},
		{e, ne, se},
	}
	for i := 0; i < 4; i++ {
		cs := check[(i+m.round)%4]
		if m.AllEmpty(cs...) {
			return cs[0]
		}
	}
	return p
}

type Map struct {
	elves map[Coord]*Elf
	round int
}

func (m *Map) Empty(c Coord) bool {
	return m.elves[c] == nil
}

func (m *Map) AllEmpty(coords ...Coord) bool {
	for _, c := range coords {
		if !m.Empty(c) {
			return false
		}
	}
	return true
}

func (m *Map) MoveTo(e *Elf, d Coord) {
	if e.pos == d {
		return
	}
	delete(m.elves, e.pos)
	m.elves[d] = e
	e.pos = d
}

func (m *Map) Round() int {
	collisions := map[Coord]int{}
	moves := map[Coord]*Elf{}
	for _, e := range m.elves {
		if e == nil {
			continue
		}
		if p := e.GetNext(m); p != e.pos {
			moves[p] = e
			collisions[p]++
		}
	}

	nm := 0
	for c, e := range moves {
		if n := collisions[c]; n == 1 {
			m.MoveTo(e, c)
			nm++
		}
	}
	m.round++
	if nm == 0 {
		return 0
	}
	return m.round
}

func (m *Map) Bounds() (Coord, Coord) {
	min, max := maxCoord, Coord{}

	for d := range m.elves {
		if dx := d[X]; dx < min[X] {
			min[X] = dx
		} else if dx > max[X] {
			max[X] = dx
		}
		if dy := d[Y]; dy < min[Y] {
			min[Y] = dy
		} else if dy > max[Y] {
			max[Y] = dy
		}
	}
	return min, max
}

func (m *Map) Score() int {
	min, max := m.Bounds()
	return ((max[X] - min[X] + 1) * (max[Y] - min[Y] + 1)) - len(m.elves)
}

func LoadMap(filename string) *Map {
	m := &Map{map[Coord]*Elf{}, 0}
	for y, line := range lib.SlurpFile(filename) {
		for x, c := range line {
			if c == '#' {
				m.MoveTo(&Elf{maxCoord}, Coord{y, x})
			}
		}
	}
	return m
}

func PartA(filename string) int {
	m := LoadMap(filename)
	for m.Round() < 10 {
	}
	return m.Score()
}

func PartB(filename string) int {
	m := LoadMap(filename)
	for m.Round() != 0 {
	}
	return m.round
}
