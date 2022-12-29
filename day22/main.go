package day22

import (
	"math"
	"regexp"

	"github.com/jscaltreto/aoc-2022/lib"
)

const (
	// Directions
	Right Dir = iota
	Down
	Left
	Up
)

var reNum *regexp.Regexp = regexp.MustCompile(`(\d+)([LR]|$)`)

type Dir int

func (d Dir) Left() Dir     { return (d + 3) % 4 }
func (d Dir) Right() Dir    { return (d + 1) % 4 }
func (d Dir) Opposite() Dir { return (d + 2) % 4 }

type Conn struct {
	cell *Cell
	dir  Dir
}

type Cell struct {
	row, col  int
	wall      bool
	neighbors [4]*Conn
}

func LoadMap(filename string, wrap bool) (*Cell, string, int) {
	var start *Cell
	firstRow, lastRow := map[int]*Cell{}, map[int]*Cell{}

	data := lib.SlurpFile(filename)
	row := 0
	nonEmpty := 0
	for ; row < len(data)-2; row++ {
		str := data[row]
		var first, last *Cell
		for col, c := range str {
			cell := &Cell{row: row + 1, col: col + 1}
			switch c {
			case ' ':
				continue
			case '#':
				cell.wall = true
			}
			nonEmpty++
			if start == nil {
				start = cell
			}
			if last == nil {
				first = cell
			} else {
				last.neighbors[Right] = &Conn{cell, Right}
				cell.neighbors[Left] = &Conn{last, Left}
			}
			if lr, f := lastRow[col]; !f {
				firstRow[col] = cell
			} else {
				lr.neighbors[Down] = &Conn{cell, Down}
				cell.neighbors[Up] = &Conn{lr, Up}
			}
			last = cell
			lastRow[col] = cell
		}
		if wrap {
			last.neighbors[Right] = &Conn{first, Right}
			first.neighbors[Left] = &Conn{last, Left}
		}
	}
	if wrap {
		for col := range firstRow {
			firstRow[col].neighbors[Up] = &Conn{lastRow[col], Up}
			lastRow[col].neighbors[Down] = &Conn{firstRow[col], Down}
		}
	}
	return start, data[row+1], int(math.Sqrt(float64(nonEmpty) / 6))
}

func Walk(cur *Cell, path string) int {
	face := Right
	for path != "" {
		m := reNum.FindStringSubmatch(path)
		path = path[len(m[0]):]
		steps := lib.StrToInt(m[1])
		for i := 0; i < steps; i++ {
			if next := cur.neighbors[face]; !next.cell.wall {
				face = next.dir
				cur = next.cell
			} else {
				break
			}
		}
		if dir := m[2]; dir != "" {
			if dir == "L" {
				face = face.Left()
			} else {
				face = face.Right()
			}
		}
	}
	return (1000 * cur.row) + (4 * cur.col) + int(face)
}

func Cubeify(start *Cell, size int) {
	next := &Conn{start, Right}

	connecting := false
	connections := 0
	var connDirA, connDirB Dir

	stack := []*Conn{}

	var cur *Conn
	for connections < 7 {
		side := []*Conn{}

		for i := 0; i < size; i++ {
			cur = next
			if connecting {
				n := len(stack) - 1
				neighbor := stack[n]
				stack = stack[:n]

				cur.cell.neighbors[connDirA] = &Conn{neighbor.cell, connDirB.Opposite()}
				neighbor.cell.neighbors[connDirB] = &Conn{cur.cell, connDirA.Opposite()}
			}
			side = append(side, cur)
			next = cur.cell.neighbors[cur.dir]
		}

		if connecting {
			// If there's nothing in the stack, continue straight.
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				cur = top
			}
			connections++
			connecting = false
		} else {
			stack = append(stack, side...)
		}

		nc := cur.cell.neighbors[cur.dir]
		if nc == nil {
			// right turn
			next = &Conn{cur.cell, cur.dir.Right()}
		} else {
			next = nc
			if nc = next.cell.neighbors[nc.dir.Left()]; nc != nil {
				// left turn; We can form some connections.
				next = nc
				connecting = true
				connDirA = next.dir.Left()
				connDirB = cur.dir.Left()
			}
		}
	}
}

func PartA(filename string) int {
	start, path, _ := LoadMap(filename, true)
	return Walk(start, path)
}

func PartB(filename string) int {
	start, path, size := LoadMap(filename, false)
	Cubeify(start, size)
	return Walk(start, path)
}
