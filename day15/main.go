package day15

import (
	"errors"
	"regexp"

	"github.com/jscaltreto/aoc-2022/lib"
)

const (
	RE = `^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$`
)

var re *regexp.Regexp = regexp.MustCompile(RE)
var outofrange error = errors.New("OutOfRange")

type Sensors []*Sensor

func (ss Sensors) CheckSensors(x, y int) (int, error) {
	for _, s := range ss {
		if nx, err := s.CheckSensor(x, y); err == nil {
			return nx, nil
		}
	}
	return x, errors.New("FoundIt")
}

type Sensor struct {
	x, y, bx, by, r int
}

func NewSensor(input string) *Sensor {
	m := re.FindStringSubmatch(input)
	s := &Sensor{
		x:  lib.StrToInt(m[1]),
		y:  lib.StrToInt(m[2]),
		bx: lib.StrToInt(m[3]),
		by: lib.StrToInt(m[4]),
	}
	s.r = lib.Abs(s.x-s.bx) + lib.Abs(s.y-s.by)
	return s
}

func (s *Sensor) IntersectRow(row int, empty, beacons map[int]bool) {
	dist := lib.Abs(s.y - row)
	rem := s.r - dist
	if s.by == row {
		beacons[s.bx] = true
	}
	for i := s.x - rem; i <= s.x+rem; i++ {
		empty[i] = true
	}
}

func (s *Sensor) CheckSensor(x, y int) (int, error) {
	if dx, dy := lib.Abs(s.x-x), lib.Abs(s.y-y); dx+dy <= s.r {
		return s.x + (s.r - dy) + 1, nil
	}
	return 0, outofrange
}

func AddSensors(filename string) Sensors {
	sensors := Sensors{}
	for _, line := range lib.SlurpFile(filename) {
		sensors = append(sensors, NewSensor(line))
	}
	return sensors
}

func PartA(filename string, row int) int {
	sensors := AddSensors(filename)

	rowBeacons := map[int]bool{}
	rowInts := map[int]bool{}
	for _, s := range sensors {
		s.IntersectRow(row, rowInts, rowBeacons)
	}

	return len(rowInts) - len(rowBeacons)
}

func PartB(filename string, max int) int {
	sensors := AddSensors(filename)
	var err error

	for y := 0; y <= max; y++ {
		for x := 0; x <= max; {
			x, err = sensors.CheckSensors(x, y)
			if err != nil {
				return (x * 4000000) + y
			}
		}
	}
	return 1
}
