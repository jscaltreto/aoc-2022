package day16

import (
	"container/heap"
	"regexp"
	"strings"

	"github.com/ernestosuarez/itertools"
	"github.com/jscaltreto/aoc-2022/lib"
)

const (
	RE = `^Valve ([A-Z]{2}) has flow rate=(\d+); tunnel(s)? lead(s)? to valve(s)? (.*)$`
)

var re *regexp.Regexp = regexp.MustCompile(RE)

type Valve struct {
	name string
	flow int
	dist map[string]int
}

func NewValve(input string, vm *ValveMap) *Valve {
	m := re.FindStringSubmatch(input)
	v := &Valve{
		name: m[1],
		flow: lib.StrToInt(m[2]),
		dist: map[string]int{m[1]: 0},
	}
	for _, d := range strings.Split(m[6], ", ") {
		v.dist[d] = 1
	}
	return v
}

type Path struct {
	v       *Valve
	visited map[string]struct{}
	i, dist int
}

func NewPath(v *Valve, p *Path, dist int, visited map[string]struct{}) *Path {
	nextVisited := map[string]struct{}{v.name: {}}
	for v, s := range visited {
		nextVisited[v] = s
	}
	return &Path{
		v:       v,
		dist:    dist,
		visited: nextVisited,
	}
}

func (p *Path) Cost() int      { return p.dist }
func (p *Path) SetIndex(i int) { p.i = i }

type ValveMap struct {
	all  map[string]*Valve
	flow map[string]*Valve
	time int
}

func NewValveMap(filename string, time int) *ValveMap {
	vm := &ValveMap{
		all:  make(map[string]*Valve),
		flow: make(map[string]*Valve),
		time: time,
	}
	for _, line := range lib.SlurpFile(filename) {
		valve := NewValve(line, vm)
		vm.all[valve.name] = valve
		if valve.flow > 0 {
			vm.flow[valve.name] = valve
		}
	}
	// pre-calculate the distances between all valves that have flow.
	// We also need the distances from `AA`...
	vm.flow["AA"] = vm.all["AA"]
	for _, source := range vm.flow {
		for _, dest := range vm.flow {
			vm.GetDistance(source, dest, make(map[string]struct{}))
		}
	}
	// But `AA` has no flow, so we remove it.
	delete(vm.flow, "AA")
	return vm
}

func (vm *ValveMap) GetDistance(source, dest *Valve, visited map[string]struct{}) {
	if _, f := source.dist[dest.name]; f {
		return
	}

	start := NewPath(source, nil, 0, make(map[string]struct{}))
	pq := lib.PQ{start}
	heap.Init(&pq)
	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Path)
		if node.v.name == dest.name {
			source.dist[node.v.name] = node.dist
			node.v.dist[source.name] = node.dist
			return
		}
		for nexts, dist := range node.v.dist {
			if _, f := node.visited[nexts]; !f {
				next := vm.all[nexts]
				if dist += node.dist; dist <= vm.time {
					nextNode := NewPath(next, node, dist, node.visited)
					heap.Push(&pq, nextNode)
				}
			}
		}
	}
}

func (vm ValveMap) BestPathSingle() int {
	return BestFlow(vm.all["AA"], vm.time, 0, vm.flow)
}

func (vm ValveMap) BestPathDouble() int {
	valves := []string{}
	for vi := range vm.flow {
		valves = append(valves, vi)
	}
	best := 0
	for vids := range itertools.CombinationsStr(valves, len(valves)/2) {
		v1, v2 := map[string]*Valve{}, map[string]*Valve{}

		for _, v := range vids {
			v1[v] = vm.all[v]
		}

		for _, v := range valves {
			if _, f := v1[v]; !f {
				v2[v] = vm.all[v]
			}
		}

		b1 := BestFlow(vm.all["AA"], vm.time, 0, v1)
		b2 := BestFlow(vm.all["AA"], vm.time, 0, v2)
		nextBest := b1 + b2
		if nextBest > best {
			best = nextBest
		}
	}
	return best
}

func BestFlow(valve *Valve, time, flow int, closed map[string]*Valve) int {
	nextClosed := map[string]*Valve{}
	for vo, v := range closed {
		if vo != valve.name {
			nextClosed[vo] = v
		}
	}
	flow += time * valve.flow

	bestFlow := flow
	for vn, v := range nextClosed {
		newTime := time - (valve.dist[vn] + 1)
		if newTime >= 0 {
			newFlow := BestFlow(v, newTime, flow, nextClosed)
			if newFlow > bestFlow {
				bestFlow = newFlow
			}
		}
	}
	return bestFlow
}

func PartA(filename string) int {
	m := NewValveMap(filename, 30)
	p := m.BestPathSingle()
	return p
}

func PartB(filename string) int {
	m := NewValveMap(filename, 26)
	p := m.BestPathDouble()
	return p
}
