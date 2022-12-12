package day12

import (
	"container/heap"
	"math"

	"github.com/jscaltreto/aoc-2022/lib"
)

type Node struct {
	x, y, c, i int
	w          byte
	v          bool
}

func (n *Node) Cost() int      { return n.c }
func (n *Node) SetIndex(i int) { n.i = i }

func (n *Node) Update(baseCost int) bool {
	g := baseCost + 1
	if g < n.c {
		n.c = g
		return true
	}
	return false
}

type Map struct {
	terrain [][]*Node
	start   []int
	end     []int
}

var nx [4]int = [4]int{0, 0, -1, 1}
var ny [4]int = [4]int{-1, 1, 0, 0}

func (m Map) Neighbors(n *Node) []*Node {
	ns := []*Node{}
	for i, ox := range nx {
		x := n.x + ox
		y := n.y + ny[i]
		if !(y < 0 || y >= len(m.terrain) || x < 0 || x >= len(m.terrain[y])) {
			if n := m.terrain[y][x]; !n.v {
				ns = append(ns, n)
			}
		}
	}
	return ns
}

func (m Map) FindPathUp() int {
	root := m.terrain[m.start[0]][m.start[1]]
	root.c = 0
	pq := lib.PQ{root}
	heap.Init(&pq)
	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)
		if node.x == m.end[1] && node.y == m.end[0] {
			return node.c
		}
		for _, n := range m.Neighbors(node) {
			if n.w <= node.w+1 {
				if updated := n.Update(node.c); updated {
					heap.Push(&pq, n)
				}
			}
		}
		node.v = true
	}
	return 0
}

func (m Map) FindPathDown() int {
	root := m.terrain[m.end[0]][m.end[1]]
	root.c = 0
	pq := lib.PQ{root}
	heap.Init(&pq)
	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)
		if node.w == 'a' {
			return node.c
		}
		for _, n := range m.Neighbors(node) {
			if node.w <= n.w+1 {
				if updated := n.Update(node.c); updated {
					heap.Push(&pq, n)
				}
			}
		}
		node.v = true
	}
	return 0
}

func LoadMap(filename string) *Map {
	data := lib.SlurpFile(filename)
	m := &Map{terrain: make([][]*Node, len(data))}
	for iy, line := range data {
		m.terrain[iy] = make([]*Node, len(line))
		for ix, c := range line {
			node := &Node{
				x: ix,
				y: iy,
				w: byte(c),
				c: math.MaxInt,
			}
			m.terrain[iy][ix] = node
			if node.w == 'S' {
				node.w = 'a'
				m.start = []int{iy, ix}
			} else if node.w == 'E' {
				node.w = 'z'
				m.end = []int{iy, ix}
			}
		}
	}
	return m
}

func PartA(filename string) int {
	m := LoadMap(filename)
	p := m.FindPathUp()
	return p
}

func PartB(filename string) int {
	m := LoadMap(filename)
	p := m.FindPathDown()
	return p
}
