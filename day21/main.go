package day21

import (
	"errors"
	"strconv"

	"github.com/jscaltreto/aoc-2022/lib"
)

const (
	Root  = "root"
	Human = "humn"
)

type OpMap map[byte]func(int, int) int

var eval OpMap = OpMap{
	'+': func(a, b int) int { return a + b },
	'-': func(a, b int) int { return a - b },
	'*': func(a, b int) int { return a * b },
	'/': func(a, b int) int { return a / b },
}

// solve for b in `a ? b = v`
var solveB OpMap = OpMap{
	'+': func(a, v int) int { return v - a },
	'-': func(a, v int) int { return a - v },
	'*': func(a, v int) int { return v / a },
	'/': func(a, v int) int { return a / v },
}

// solve for a in `a ? b = v`
var solveA OpMap = OpMap{
	'+': func(b, v int) int { return v - b },
	'-': func(b, v int) int { return v + b },
	'*': func(b, v int) int { return v / b },
	'/': func(b, v int) int { return v * b },
}

type Monkeys map[string]string

func (m Monkeys) Eval(name string, p2 bool) (int, error) {
	if p2 && name == Human {
		return 0, errors.New(Human)
	}

	str := m[name]
	if i, err := strconv.Atoi(str); err == nil {
		return i, nil
	}

	strs, ints := [2]string{str[:4], str[7:]}, [2]int{}
	var err error
	for i := 0; i < 2; i++ {
		if ints[i], err = m.Eval(strs[i], p2); err != nil {
			return 0, err
		}
	}

	return eval[str[5]](ints[0], ints[1]), nil
}

func (m Monkeys) Solve(name string, val int) int {
	str := m[name]
	m1, m2, op := str[:4], str[7:], str[5]

	i, err := m.Eval(m1, true)
	mm := m2
	solveFunc := solveB[op]
	if err != nil {
		i, _ = m.Eval(m2, false)
		mm = m1
		solveFunc = solveA[op]
	}

	val = solveFunc(i, val)
	if mm == Human {
		return val
	}
	return m.Solve(mm, val)
}

func NewMonkeys(filename string) Monkeys {
	m := Monkeys{}
	for _, l := range lib.SlurpFile(filename) {
		m[l[:4]] = l[6:]
	}
	return m
}

func PartA(filename string) int {
	m := NewMonkeys(filename)
	i, _ := m.Eval(Root, false)
	return i
}

func PartB(filename string) int {
	m := NewMonkeys(filename)
	m1, m2 := m[Root][:4], m[Root][7:]
	val, err := m.Eval(m1, true)
	mm := m2
	if err != nil {
		val, _ = m.Eval(m2, false)
		mm = m1
	}
	return m.Solve(mm, val)
}
