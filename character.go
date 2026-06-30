package main

import "fmt"

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

func (c *character) displayInfo() {
	fmt.Printf(" Name: %s\n HP: %d/%d\n Class: %s\n", c.name, c.currHP, c.maxHP, c.className)
}