package day20

import (
	"github.com/jscaltreto/aoc-2022/lib"
)

type Number struct {
	i          int
	next, prev *Number
}

func (n *Number) GetNthValue(i int) int {
	cur := n
	for ; i > 0; i-- {
		cur = cur.next
	}
	return cur.i
}

func Decrypt(filename string, key, iters int) int {
	raw := lib.SlurpFile(filename)
	size := len(raw)
	encrypted := make([]*Number, size)
	var prev, zero *Number
	for i, s := range raw {
		num := lib.StrToInt(s)
		encrypted[i] = &Number{i: num * key, prev: prev}
		if num == 0 {
			zero = encrypted[i]
		}
		if prev != nil {
			prev.next = encrypted[i]
		}
		prev = encrypted[i]
	}
	encrypted[0].prev = prev
	prev.next = encrypted[0]

	var on, op *Number
	for iter := 0; iter < iters; iter++ {
		for _, n := range encrypted {
			shift := n.i % (size - 1)
			if shift != 0 {
				for i := 0; i < lib.Abs(shift); i++ {
					on = n.next
					op = n.prev
					on.prev = op
					op.next = on
					if shift > 0 {
						n.next = on.next
						on.next.prev = n
						on.next = n
						n.prev = on
					} else {
						n.prev = op.prev
						op.prev.next = n
						op.prev = n
						n.next = op
					}
				}
			}
		}
	}

	return zero.GetNthValue(1000) + zero.GetNthValue(2000) + zero.GetNthValue(3000)
}

func PartA(filename string) int {
	return Decrypt(filename, 1, 1)
}

func PartB(filename string) int {
	return Decrypt(filename, 811589153, 10)
}
