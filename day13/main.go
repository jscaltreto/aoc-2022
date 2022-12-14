package day13

import (
	"encoding/json"
	"log"
	"sort"

	"github.com/jscaltreto/aoc-2022/lib"
)

const (
	Int = iota
	List
)

type Packet struct {
	data interface{}
	str  string
}

func NewPacket(p string) Packet {
	var d interface{}
	err := json.Unmarshal([]byte(p), &d)
	if err != nil {
		log.Fatalf("Invalid! %v", err)
	}
	return Packet{data: d, str: p}
}

func (p *Packet) Type() int {
	switch p.data.(type) {
	case float64:
		return Int
	case []interface{}:
		return List
	}
	panic("Something went wrong")
}

func (p *Packet) AsList() []interface{} {
	if p.Type() == Int {
		return []interface{}{p.data.(float64)}
	}
	return p.data.([]interface{})
}

func (p1 *Packet) Cmp(p2 Packet) int {
	if p1.Type() == Int && p2.Type() == Int {
		return int(p1.data.(float64) - p2.data.(float64))
	}
	p1l, p2l := p1.AsList(), p2.AsList()
	for i := range p1l {
		if i == len(p2l) {
			return 1
		}
		np1 := Packet{data: p1l[i]}
		np2 := Packet{data: p2l[i]}
		if cmp := np1.Cmp(np2); cmp != 0 {
			return cmp
		}
	}
	if len(p1l) == len(p2l) {
		return 0
	}
	return -1
}

type PacketList []Packet

func (pl PacketList) Len() int {
	return len(pl)
}

func (pl PacketList) Swap(i, j int) {
	pl[i], pl[j] = pl[j], pl[i]
}

func (pl PacketList) Less(i, j int) bool {
	return pl[i].Cmp(pl[j]) < 0
}

func PartA(filename string) int {
	data := lib.SlurpFile(filename)
	sum := 0
	pid := 0
	for i := 0; i < len(data); i += 3 {
		pid++
		p1 := NewPacket(data[i])
		p2 := NewPacket(data[i+1])
		if cmp := p1.Cmp(p2); cmp < 0 {
			sum += pid
		}
	}
	return sum
}

func PartB(filename string) int {
	const (
		DivA = "[[2]]"
		DivB = "[[6]]"
	)
	packets := PacketList{
		NewPacket(DivA),
		NewPacket(DivB),
	}

	for _, data := range lib.SlurpFile(filename) {
		if data != "" {
			packets = append(packets, NewPacket(data))
		}
	}

	sort.Sort(packets)

	result := 1
	for idx, packet := range packets {
		if packet.str == DivA || packet.str == DivB {
			result *= idx + 1
		}
	}

	return result
}
