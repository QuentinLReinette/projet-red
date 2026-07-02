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
	return &character{name, 100, 50, 1, 100, human, &inventory{capacity: 10}, []skill{}, []*equipment{}}
}

func newElf(name string) *character {
	return &character{name, 80, 40, 1, 100, elf, &inventory{capacity: 10}, []skill{}, []*equipment{}}
}

func newDwarf(name string) *character {
	return &character{name, 120, 60, 1, 100, dwarf, &inventory{capacity: 10}, []skill{}, []*equipment{}}
}
