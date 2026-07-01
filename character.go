package main

import "fmt"

type character struct {
	name               string
	maxHP, currHP, lvl int
	class
	*inventory
}

func initChar(name, className string, maxHP int, newInventory *inventory) *character {
	newClass := class{className}
	return &character{name: name, maxHP: maxHP, currHP: maxHP * 4 / 10, lvl: 1, inventory: newInventory, class: newClass}
}

func (c *character) displayInfo() {
	fmt.Printf(" Name: %s\n HP: %d/%d\n Class: %s\n", c.name, c.currHP, c.maxHP, c.className)
}

func (c *character) checkAlive() {
	if c.currHP <= 0 {
		fmt.Printf("%s died. They're back with half of their max HP\n", c.name)
		c.currHP = c.maxHP / 2
	}
}
