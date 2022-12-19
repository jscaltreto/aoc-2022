package day14

import (
	"math"
	"strings"

	"github.com/jscaltreto/aoc-2022/lib"
)

func Between(s, e int) []int {
	if s == e {
		return []int{s}
	}
	f := 1
	if s > e {
		f = -1
	}
	pts := make([]int, 0)
	for i := s; i != e; i += f {
		pts = append(pts, i)
	}
	pts = append(pts, e)
	return pts
}

type Map struct {
	m     map[int]map[int]bool
	minx  int
	miny  int
	maxx  int
	maxy  int
	floor int
}

func (m *Map) AddRock(x, y int) {
	if _, found := m.m[x]; !found {
		m.m[x] = make(map[int]bool, 0)
	}
	m.m[x][y] = true
	if x < m.minx {
		m.minx = x
	}
	if x > m.maxx {
		m.maxx = x
	}
	if y < m.miny {
		m.miny = y
	}
	if y > m.maxy {
		m.maxy = y
	}
}

func (m *Map) HasRock(x, y int) bool {
	if y == m.floor {
		return true
	}
	_, f := m.m[x][y]
	return f
}

func (m *Map) NextSandPos(x, y int) (cx, cy int) {
	cx, cy = x, y+1
	if !m.HasRock(cx, cy) {
		return
	}
	cx = x - 1
	if !m.HasRock(cx, cy) {
		return
	}
	cx = x + 1
	if !m.HasRock(cx, cy) {
		return
	}
	cx, cy = x, y
	return
}

func (m *Map) AddSand(comp func(int) bool) int {
	addedSand := 0
	for x, y := 500, m.miny-1; comp(y); {
		if nx, ny := m.NextSandPos(x, y); nx == x && ny == y {
			m.AddRock(x, y)
			x, y = 500, m.miny-1
			addedSand++
		} else {
			x, y = nx, ny
		}
	}
	return addedSand
}

func (m *Map) AddLine(p1s, p2s string) {
	p1 := strings.Split(p1s, ",")
	p1x, p1y := lib.StrToInt(p1[0]), lib.StrToInt(p1[1])
	p2 := strings.Split(p2s, ",")
	p2x, p2y := lib.StrToInt(p2[0]), lib.StrToInt(p2[1])
	for _, x := range Between(p1x, p2x) {
		for _, y := range Between(p1y, p2y) {
			m.AddRock(x, y)
		}
	}
}

func NewMap(filename string) *Map {
	cave := &Map{
		m:    map[int]map[int]bool{},
		minx: math.MaxInt,
		miny: math.MaxInt,
	}
	for _, line := range lib.SlurpFile(filename) {
		points := strings.Split(line, " -> ")
		for i := 1; i < len(points); i++ {
			cave.AddLine(points[i-1], points[i])
		}
	}
	return cave
}

func PartA(filename string) int {
	cave := NewMap(filename)

	return cave.AddSand(func(y int) bool { return y <= cave.maxy })
}

func PartB(filename string) int {
	cave := NewMap(filename)
	cave.floor = cave.maxy + 2

	return cave.AddSand(func(y int) bool { return y >= 0 })
}
