package day08

import (
	"github.com/jscaltreto/aoc-2022/lib"
)

func ParseInput(data []string) [][]rune {
	grid := make([][]rune, len(data))
	for row, trees := range data {
		grid[row] = make([]rune, len(trees))
		for col, tree := range trees {
			grid[row][col] = tree
		}
	}
	return grid
}

func isVisible(grid [][]rune, row, col int) bool {
	me := grid[row][col]
	u, d, l, r := row-1, row+1, col-1, col+1
	var stuck bool
	for (u >= 0 && d < len(grid) && l >= 0 && r < len(grid)) && !stuck {
		stuck = true
		if u >= 0 && grid[u][col] < me {
			u--
			stuck = false
		}
		if d < len(grid) && grid[d][col] < me {
			d++
			stuck = false
		}
		if l >= 0 && grid[row][l] < me {
			l--
			stuck = false
		}
		if r < len(grid[row]) && grid[row][r] < me {
			r++
			stuck = false
		}
	}
	return !stuck
}

func PartA(filename string) int {
	grid := ParseInput(lib.SlurpFile(filename))

	visible := (len(grid) * 2) + (len(grid[0]) * 2) - 4
	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[row])-1; col++ {
			if isVisible(grid, row, col) {
				visible++
			}
		}
	}
	return visible
}

func score(grid [][]rune, row, col int) int {
	me := grid[row][col]
	u, d, l, r := row-1, row+1, col-1, col+1
	cu, cd, cl, cr := 1, 1, 1, 1
	var su, sd, sl, sr bool
	for !(su && sd && sl && sr) {
		if !su {
			if u <= 0 || grid[u][col] >= me {
				su = true
			} else {
				cu++
				u--
			}
		}
		if !sd {
			if d >= len(grid)-1 || grid[d][col] >= me {
				sd = true
			} else {
				cd++
				d++
			}
		}
		if !sl {
			if l <= 0 || grid[row][l] >= me {
				sl = true
			} else {
				cl++
				l--
			}
		}
		if !sr {
			if r >= len(grid[row])-1 || grid[row][r] >= me {
				sr = true
			} else {
				cr++
				r++
			}
		}
	}
	return cu * cd * cl * cr
}

func PartB(filename string) int {
	grid := ParseInput(lib.SlurpFile(filename))

	var best int
	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[row])-1; col++ {
			if s := score(grid, row, col); s > best {
				best = s
			}
		}
	}
	return best
}
