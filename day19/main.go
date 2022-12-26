package day19

import (
	"math"
	"regexp"
	"sync"

	"github.com/jscaltreto/aoc-2022/lib"
)

const RE = `^Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.$`

var re *regexp.Regexp = regexp.MustCompile(RE)

const (
	Ore = iota
	Clay
	Obsidian
	Geode
)

type State struct {
	time, i           int
	resources, robots map[int]int
}

func (s *State) Total(t int) int {
	return s.resources[t] + (s.time * s.robots[t])
}

func (s State) BestCase() int {
	return s.Total(Geode) + ((s.time * (s.time + 1)) / 2)
}

func (s *State) Clone() *State {
	newState := &State{
		time:      s.time,
		robots:    map[int]int{},
		resources: map[int]int{},
	}
	for resource := range s.robots {
		newState.robots[resource] = s.robots[resource]
		newState.resources[resource] = s.resources[resource]
	}
	return newState
}

func (s *State) TimeToBuild(robot int, cost map[int]int) int {
	timeNeeded := 0
	for resource, cost := range cost {
		needed := cost - s.resources[resource]
		time := 0
		if needed > 0 {
			if s.robots[resource] == 0 {
				return -1
			}
			time = (needed + s.robots[resource] - 1) / s.robots[resource]
		}
		if time > timeNeeded {
			timeNeeded = time
		}
	}
	return timeNeeded
}

func TestBlueprint(line string, time int) (int, int) {
	m := re.FindStringSubmatch(line)
	id := lib.StrToInt(m[1])
	cost := map[int]map[int]int{
		Ore:      {Ore: lib.StrToInt(m[2])},
		Clay:     {Ore: lib.StrToInt(m[3])},
		Obsidian: {Ore: lib.StrToInt(m[4]), Clay: lib.StrToInt(m[5])},
		Geode:    {Ore: lib.StrToInt(m[6]), Obsidian: lib.StrToInt(m[7])},
	}

	// Figure out the maximum number of robots of each type needed
	maxRobots := map[int]int{Geode: math.MaxInt}
	for r1 := range []int{Ore, Clay, Obsidian} {
		for _, cost := range cost {
			if cost[r1] > maxRobots[r1] {
				maxRobots[r1] = cost[r1]
			}
		}
	}

	start := &State{
		time:      time,
		robots:    map[int]int{Ore: 1},
		resources: map[int]int{},
	}

	checkStates := []*State{start}

	best := start
	for len(checkStates) > 0 {
		check := checkStates[0]
		checkStates = checkStates[1:]
		if t := check.Total(Geode); t >= best.Total(Geode) {
			best = check
		}
		for robot := range cost {
			if check.robots[robot] < maxRobots[robot] {
				tn := check.TimeToBuild(robot, cost[robot])
				if tn >= 0 && tn < check.time-1 {
					newState := check.Clone()
					newState.robots[robot]++
					tn++
					newState.time -= tn
					for resource := range cost {
						newState.resources[resource] += (tn * check.robots[resource]) - cost[robot][resource]
					}
					if newState.BestCase() > best.Total(Geode) {
						checkStates = append(checkStates, newState)
					}
				}
			}
		}
	}

	return best.Total(Geode), id
}

func PartA(filename string) int {
	total := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, line := range lib.SlurpFile(filename) {
		wg.Add(1)
		go func(l string) {
			defer wg.Done()
			max, id := TestBlueprint(l, 24)
			mu.Lock()
			defer mu.Unlock()
			total += max * id
		}(line)
	}
	wg.Wait()
	return total
}

func PartB(filename string) int {
	product := 1
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, line := range lib.SlurpFile(filename)[:3] {
		wg.Add(1)
		go func(l string) {
			defer wg.Done()
			max, _ := TestBlueprint(l, 32)
			mu.Lock()
			defer mu.Unlock()
			product *= max
		}(line)
	}
	wg.Wait()
	return product
}
