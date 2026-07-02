package main

import (
	"fmt"
	"slices"
)

type character struct {
	name                                     string
	maxHP, currHP, lvl, xp, gold, initiative int
	class
	*inventory
	skills    []skill
	equipment []*equipment
}

func initChar(name string, class class, maxHP, initiative int) *character {
	newChar := &character{name: name, maxHP: maxHP, currHP: maxHP * 4 / 10, lvl: 1, xp: 0, gold: 100, initiative: initiative, class: class, inventory: &inventory{capacity: 10}}
	newChar.skills = []skill{newPunch(newChar)}
	newChar.potions = []*potion{newHealthPot(newChar, 3)}
	return newChar
}

func (c *character) displayInfo() {
	fmt.Printf(" Name: %s\n HP: %d/%d\n Class: %s\n Gold: %d\n Initiative: %d\n", c.name, c.currHP, c.maxHP, c.class, c.gold, c.initiative)
}

func (c *character) checkAlive() bool {
	if c.currHP <= 0 {
		fmt.Printf("%s died. They're back with half of their max HP\n", c.name)
		c.currHP = c.maxHP / 2
		return false
	}
	return true
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

func (c *character) checkXP() {
	if c.xp >= 100 {
		c.lvlUP()
	}
}

func (c *character) lvlUP() {
	for c.xp >= 100 {
		c.lvl++
		c.xp -= 100
		c.maxHP += 10
	}
	c.currHP = c.maxHP
	fmt.Printf("%s leveled up! They are now level %d and have %d HP\n", c.name, c.lvl, c.maxHP)
}
