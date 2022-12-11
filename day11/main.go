package day11

import (
	"sort"
	"strings"

	"github.com/jscaltreto/aoc-2022/lib"
)

type Monkey struct {
	items       []int
	op          func(int) int
	cmp         int
	truem       int
	falsem      int
	inspections int
}

func (m Monkey) nextMonkey(i, mod, div int) (int, int) {
	if ni := (m.op(i) % mod) / div; ni%m.cmp == 0 {
		return ni, m.truem
	} else {
		return ni, m.falsem
	}
}

type Monkeys struct {
	ms  []*Monkey
	mod int
	div int
}

func (ms Monkeys) doRound() {
	for _, m := range ms.ms {
		for _, i := range m.items {
			m.inspections++
			ni, nm := m.nextMonkey(i, ms.mod, ms.div)
			ms.ms[nm].items = append(ms.ms[nm].items, ni)
		}
		m.items = []int{}
	}
}

func (ms Monkeys) getBusiness(rounds int) int {
	for i := 1; i <= rounds; i++ {
		ms.doRound()
	}
	sort.Slice(ms.ms, func(i, j int) bool {
		return ms.ms[i].inspections > ms.ms[j].inspections
	})
	return ms.ms[0].inspections * ms.ms[1].inspections
}

func LoadMonkeys(filename string, div int) Monkeys {
	data := lib.SlurpFile(filename)
	monkeys := Monkeys{ms: []*Monkey{}, div: div, mod: 1}
	for i := 0; i < len(data); i += 7 {
		monkey := &Monkey{items: []int{}}
		for _, item := range strings.Split(data[i+1][18:], ", ") {
			monkey.items = append(monkey.items, lib.StrToInt(item))
		}
		opStr := strings.Split(data[i+2][23:], " ")
		if opStr[0] == "*" {
			if opStr[1] == "old" {
				monkey.op = func(o int) int { return (o * o) }
			} else {
				c := lib.StrToInt(opStr[1])
				monkey.op = func(o int) int { return (o * c) }
			}
		} else {
			if opStr[1] == "old" {
				monkey.op = func(o int) int { return (o + o) }
			} else {
				c := lib.StrToInt(opStr[1])
				monkey.op = func(o int) int { return (o + c) }
			}
		}
		monkey.cmp = lib.StrToInt(data[i+3][21:])
		monkey.truem = lib.StrToInt(data[i+4][29:])
		monkey.falsem = lib.StrToInt(data[i+5][30:])
		monkeys.ms = append(monkeys.ms, monkey)
		monkeys.mod *= monkey.cmp
	}
	return monkeys
}

func PartA(filename string) int {
	return LoadMonkeys(filename, 3).getBusiness(20)
}

func PartB(filename string) int {
	return LoadMonkeys(filename, 1).getBusiness(10000)
}
