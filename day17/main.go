package day17

import (
	"fmt"

	"github.com/jscaltreto/aoc-2022/lib"
)

func CheckCollision(a, b uint16) bool {
	return a&b != 0
}

const (
	Wall uint16 = 0b100000001
)

var (
	Rocks []Rock = []Rock{
		{
			0b00111100,
		},
		{
			0b00010000,
			0b00111000,
			0b00010000,
		},
		{
			0b00111000,
			0b00001000,
			0b00001000,
		},
		{
			0b00100000,
			0b00100000,
			0b00100000,
			0b00100000,
		},
		{
			0b00110000,
			0b00110000,
		},
	}
	NewWall Rock = Rock{Wall, Wall, Wall, Wall}
)

type Rock []uint16

func (r Rock) Blow(dir byte, tower Rock) bool {
	nextShape := make(Rock, len(r))
	tower = tower[len(tower)-len(r):]

	for i, row := range r {
		nextRow := row
		switch dir {
		case '<':
			nextRow <<= 1
		case '>':
			nextRow >>= 1
		}
		if CheckCollision(nextRow, tower[i]) {
			return false
		}
		nextShape[i] = nextRow
	}
	copy(r, nextShape)
	return true
}

type Tower struct {
	tower     Rock
	wind      string
	windIndex int
	nextRock  int
}

func (t *Tower) AddRock() {
	rockHeight := len(Rocks[t.nextRock])
	rock := make(Rock, rockHeight)
	towerAdd := make(Rock, rockHeight)
	copy(rock, Rocks[t.nextRock])
	copy(towerAdd, NewWall)
	t.nextRock = (t.nextRock + 1) % len(Rocks)

	for i := 0; i < 3; i++ {
		rock.Blow(t.GetWind(), towerAdd)
	}

	t.tower = append(t.tower, towerAdd...)
	rockPos := len(t.tower)
	for rockPos >= rockHeight {
		rock.Blow(t.GetWind(), t.tower[:rockPos])
		rockPos--
		if rockPos-rockHeight >= 0 {
			if !rock.Blow('-', t.tower[:rockPos]) {
				break
			}
			if t.tower[rockPos]|Wall == Wall {
				t.tower = t.tower[:rockPos]
			}
		}
	}
	for i := rockHeight - 1; i >= 0; i-- {
		t.tower[rockPos] |= rock[i]
		rockPos--
	}
}

func (t *Tower) GetWind() byte {
	wind := t.wind[t.windIndex]
	t.windIndex = (t.windIndex + 1) % len(t.wind)
	return wind
}

func GetHeightAfter(wind string, totalRocks int) int {
	t := &Tower{
		tower: Rock{},
		wind:  wind,
	}
	cycleTracker := make(map[string][]int)
	var heightPerCycle, cycleLength int
	numRocks := 0
	for ; numRocks < totalRocks; numRocks++ {
		if len(t.tower) > 100 {
			l10 := fmt.Sprintf("%d-%d-%s", t.windIndex, t.nextRock, fmt.Sprint(t.tower[len(t.tower)-10:]))
			if h, f := cycleTracker[l10]; f {
				heightPerCycle = len(t.tower) - h[0]
				cycleLength = numRocks - h[1]
				break
			}
			cycleTracker[l10] = []int{len(t.tower), numRocks}
		}
		t.AddRock()
	}
	numCycles := (totalRocks - numRocks) / cycleLength
	numRocks += numCycles * cycleLength

	for ; numRocks < totalRocks; numRocks++ {
		t.AddRock()
	}

	return (numCycles * heightPerCycle) + len(t.tower)
}

func PartA(filename string) int {
	return GetHeightAfter(lib.SlurpFile(filename)[0], 2022)
}

func PartB(filename string) int {
	return GetHeightAfter(lib.SlurpFile(filename)[0], 1000000000000)
}
