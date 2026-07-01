package main

import "fmt"

type character struct {
	name               string
	maxHP, currHP, lvl int
	class
	*inventory
	skills []skill
}

func initChar(name string, class class, maxHP int, newInventory *inventory) *character {
	return &character{name: name, maxHP: maxHP, currHP: maxHP * 4 / 10, lvl: 1, inventory: newInventory, class: class}
}

func (c *character) displayInfo() {
	fmt.Printf(" Name: %s\n HP: %d/%d\n Class: %s\n", c.name, c.currHP, c.maxHP, c)
}

func (c *character) checkAlive() {
	if c.currHP <= 0 {
		fmt.Printf("%s died. They're back with half of their max HP\n", c.name)
		c.currHP = c.maxHP / 2
	}
}

func (c *character) spellBook(skill skillType) {
	switch skill {
	case punch:
		c.skills = append(c.skills, newPunch(c))
	}
	fmt.Printf("%s learned %s\n", c.name, skill)
}

func (c *character) checkHasSkill(skill skillType) bool {
	for _, s := range c.skills {
		if s.skillType == skill {
			return true
		}
	}
	return false
}
