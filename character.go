package main

type character struct {
	name               string
	maxHP, currHP, lvl int
	class
	inventory
}

func initChar(name, className string, maxHP int) *character {
	newClass := class{className}
	newInventory := inventory{make([]potion, 3)}
	return &character{name: name, maxHP: maxHP, currHP: maxHP * 4 / 10, lvl: 1, inventory: newInventory, class: newClass}
}
