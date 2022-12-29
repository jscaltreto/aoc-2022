package day24

import (
	"github.com/jscaltreto/aoc-2022/lib"
)

const (
	Y int = iota
	X

	Empty Wind = '.'
	Left  Wind = '<'
	Right Wind = '>'
	Up    Wind = '^'
	Down  Wind = 'v'
)

var Moves map[Wind]Coord = map[Wind]Coord{
	Left:  {0, -1},
	Right: {0, 1},
	Up:    {-1, 0},
	Down:  {1, 0},
}

type Coord [2]int

func (c Coord) NextCoords() []Coord {
	nc := []Coord{c}
	for _, m := range Moves {
		nc = append(nc, Coord{c[Y] + m[Y], c[X] + m[X]})
	}
	return nc
}

type Wind byte

type Map struct {
	m    map[Coord][]Wind
	w, h int
}

func (m *Map) AddWind(c Coord, w Wind) {
	if _, f := m.m[c]; !f {
		m.m[c] = []Wind{}
	}
	if w != Empty {
		m.m[c] = append(m.m[c], w)
	}
}

func (m *Map) NextMap() *Map {
	nm := &Map{map[Coord][]Wind{}, m.w, m.h}
	for c, ws := range m.m {
		nm.AddWind(c, Empty)
		for _, w := range ws {
			tr := Moves[w]
			nc := Coord{(c[Y] + tr[Y] + m.h) % m.h, (c[X] + tr[X] + m.w) % m.w}
			nm.AddWind(nc, w)
		}
	}
	return nm
}

func (m *Map) CanMove(c Coord) bool {
	return c[X] >= 0 && c[X] < m.w && c[Y] >= 0 && c[Y] < m.h && len(m.m[c]) == 0
}

func ShortestDist(m *Map, start, goal Coord) (int, *Map) {
	states := map[Coord]struct{}{start: {}}
	for t := 1; ; t++ {
		nextMap := m.NextMap()
		nextStates := make(map[Coord]struct{})
		for s := range states {
			if s == goal {
				return t, m
			}
			for _, nc := range s.NextCoords() {
				if nc == start || nc == goal || nextMap.CanMove(nc) {
					nextStates[nc] = struct{}{}
				}
			}
		}
		states = nextStates
		m = nextMap
	}
}

func LoadMap(filename string) *Map {
	MapCache = map[int]*Map{}
	data := lib.SlurpFile(filename)
	m := &Map{map[Coord][]Wind{}, len(data[0]) - 2, len(data) - 2}
	for y := 1; y < len(data)-1; y++ {
		line := data[y]
		for x := 1; x < len(line)-1; x++ {
			m.AddWind(Coord{y - 1, x - 1}, Wind(line[x]))
		}
	}
	return m
}

var MapCache map[int]*Map

func PartA(filename string) int {
	m := LoadMap(filename)
	t, _ := ShortestDist(m, Coord{-1, 0}, Coord{m.h, m.w - 1})
	return t - 1
}

func PartB(filename string) int {
	m := LoadMap(filename)
	t1, m := ShortestDist(m, Coord{-1, 0}, Coord{m.h, m.w - 1})
	t2, m := ShortestDist(m.NextMap(), Coord{m.h, m.w - 1}, Coord{-1, 0})
	t3, m := ShortestDist(m.NextMap(), Coord{-1, 0}, Coord{m.h, m.w - 1})

	return t1 + t2 + t3 - 1
}
