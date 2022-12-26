package day18

import (
	"math"
	"strings"

	"github.com/jscaltreto/aoc-2022/lib"
)

type Coord [3]int

type Droplet struct {
	cubes    map[Coord]struct{}
	shell    map[Coord]struct{}
	min, max []int
}

func (d Droplet) CheckNeighbors(c Coord) int {
	surfaceArea := 0
	for axis := 0; axis < 3; axis++ {
		dir := 1
		for i := 0; i < 2; i++ {
			check := c
			check[axis] += dir
			if _, f := d.cubes[check]; !f {
				surfaceArea++
			}
			dir = -1
		}
	}
	return surfaceArea
}

func (d Droplet) CountOuterShell(c Coord) int {
	for axis := 0; axis < 3; axis++ {
		if _, f := d.shell[c]; f || c[axis] < d.min[axis]-1 || c[axis] > d.max[axis]+1 {
			return 0
		}
	}
	if _, f := d.cubes[c]; f {
		return 1
	}
	d.shell[c] = struct{}{}
	surfaceArea := 0
	for axis := 0; axis < 3; axis++ {
		dir := 1
		for i := 0; i < 2; i++ {
			check := c
			check[axis] += dir
			surfaceArea += d.CountOuterShell(check)
			dir = -1
		}
	}
	return surfaceArea
}

func NewDroplet(filename string) Droplet {
	d := Droplet{
		cubes: map[Coord]struct{}{},
		shell: map[Coord]struct{}{},
		min:   []int{math.MaxInt, math.MaxInt, math.MaxInt},
		max:   []int{0, 0, 0},
	}
	for _, coord := range lib.SlurpFile(filename) {
		c := Coord{}
		for i, s := range strings.Split(coord, ",") {
			c[i] = lib.StrToInt(s)
			if c[i] < d.min[i] {
				d.min[i] = c[i]
			} else if c[i] > d.max[i] {
				d.max[i] = c[i]
			}
		}
		d.cubes[c] = struct{}{}
	}
	return d
}

func PartA(filename string) int {
	d := NewDroplet(filename)
	surfaceArea := 0
	for c := range d.cubes {
		surfaceArea += d.CheckNeighbors(c)
	}
	return surfaceArea
}

func PartB(filename string) int {
	d := NewDroplet(filename)
	surfaceArea := d.CountOuterShell(Coord{d.min[0], d.min[1], d.min[2]})
	return surfaceArea
}
