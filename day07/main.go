package day07

import (
	"strings"

	"github.com/jscaltreto/aoc-2022/lib"
)

const (
	Total    = 70000000
	Required = 30000000
)

type file struct {
	parent   *file
	children map[string]*file
	size     int
	allDirs  []*file
}

func (f *file) addChild(name string, size int) *file {
	fh := &file{
		parent:   f,
		size:     size,
		children: make(map[string]*file),
	}
	f.children[name] = fh
	c := f
	for c != nil {
		c.size += size
		c = c.parent
	}
	return fh
}

func ParseInput(data []string) *file {
	root := &file{
		children: make(map[string]*file),
		allDirs:  make([]*file, 0),
	}
	cwd := root

CMD:
	for i := 0; i < len(data); i++ {
		line := data[i]
		switch line[2:4] {
		case "ls":
			for i++; i < len(data); i++ {
				de := strings.Split(data[i], " ")
				switch de[0] {
				case "$":
					i--
					continue CMD
				case "dir":
					root.allDirs = append(root.allDirs, cwd.addChild(de[1], 0))
				default:
					cwd.addChild(de[1], lib.StrToInt(de[0]))
				}
			}
		case "cd":
			switch line[5:] {
			case "/":
				cwd = root
			case "..":
				cwd = cwd.parent
			default:
				cwd = cwd.children[line[5:]]
			}
		}
	}

	return root
}

func PartA(filename string) int {
	root := ParseInput(lib.SlurpFile(filename))

	totalSize := 0
	for _, f := range root.allDirs {
		if f.size <= 100000 {
			totalSize += f.size
		}
	}

	return totalSize
}

func PartB(filename string) (overlap int) {
	root := ParseInput(lib.SlurpFile(filename))

	needed := Required - (Total - root.size)
	smallest := root.size
	for _, f := range root.allDirs {
		if f.size > needed && f.size < smallest {
			smallest = f.size
		}
	}

	return smallest
}
