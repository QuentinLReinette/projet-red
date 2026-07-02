package main

import (
	"fmt"
	"slices"
)

type character struct {
	name                     string
	maxHP, currHP, lvl, gold int
	class
	*inventory
	skills []skill
}

func initChar(name string, class class, maxHP int, newInventory *inventory) *character {
	return &character{name: name, maxHP: maxHP, currHP: maxHP * 4 / 10, lvl: 1, gold: 100, inventory: newInventory, class: class}
}

func (c *character) displayInfo() {
	fmt.Printf(" Name: %s\n HP: %d/%d\n Class: %s\n Gold: %d\n", c.name, c.currHP, c.maxHP, c, c.gold)
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
	case fireBall:
		c.skills = append(c.skills, newFireBall(c))
	}
	for _, b := range c.inventory.books {
		if b.skillType == skill {
			c.inventory.books = slices.Delete(c.inventory.books, slices.IndexFunc(c.inventory.books, func(b *book) bool { return b.skillType == skill }), 1)
			break
		}
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
