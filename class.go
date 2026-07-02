package main

type class int

const (
	human class = iota
	elf
	dwarf
)

var classNames = map[class]string{
	human: "human",
	elf:   "elf",
	dwarf: "dwarf",
}

func (c class) String() string {
	return classNames[c]
}

func newHuman(name string) *character {
	return initChar(name, human, 100, 10)
}

func newElf(name string) *character {
	return initChar(name, elf, 80, 15)
}

func newDwarf(name string) *character {
	return initChar(name, dwarf, 120, 5)
}
