package main

import (
	"fmt"
	"slices"
)

type character struct {
	name             string
	maxHP, currHP    int
	maxMP, currMP    int
	lvl, xp          int
	gold, initiative int
	class
	*inventory
	skills    []skill
	equipment []*equipment
}

func initChar(name string, class class, maxHP, maxMP, initiative int) *character {
	newChar := &character{name: name, maxHP: maxHP, currHP: maxHP * 4 / 10, maxMP: maxMP, currMP: maxMP, lvl: 1, xp: 0, gold: 100, initiative: initiative, class: class, inventory: &inventory{capacity: 10}}
	newChar.skills = []skill{newPunch()}
	newChar.potions = []*potion{newHealthPot(3)}
	return newChar
}

func (c *character) displayInfo() {
	fmt.Printf(" Name: %s\n HP: %d/%d\n MP: %d/%d\n Class: %s\n Gold: %d\n Initiative: %d\n", c.name, c.currHP, c.maxHP, c.currMP, c.maxMP, c.class, c.gold, c.initiative)
}

func (c *character) checkAlive() bool {
	if c.currHP <= 0 {
		fmt.Printf("%s died. They're back with half of their max HP and MP\n", c.name)
		c.currHP = c.maxHP / 2
		c.currMP = c.maxMP / 2
		return false
	}
	return true
}

func (c *character) spellBook(skill skillType) {
	switch skill {
	case punch:
		c.skills = append(c.skills, newPunch())
	case fireBall:
		c.skills = append(c.skills, newFireBall())
	}
	for idx, b := range c.inventory.books {
		if b.skillType == skill {
			c.inventory.books = slices.Delete(c.inventory.books, idx, idx+1)
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
		c.maxMP += 5
	}
	c.currHP = c.maxHP
	c.currMP = c.maxMP
	fmt.Printf("%s leveled up! They are now level %d and have %d HP\n", c.name, c.lvl, c.maxHP)
}

func (c *character) attack(monster *monster) {
	options := []menuOption{{"Basic Attack - 5 dmg", func() { c.basicAttack(monster) }}}
	for _, s := range c.skills {
		options = append(options, menuOption{fmt.Sprintf("%s - %d dmg (MP Cost: %d)", s, s.dmg, s.mpCost), func() { s.skillUse(monster) }})
	}
	menuPrint(options, true, false)
}

func (c *character) basicAttack(monster *monster) {
	monster.currHP -= 10
	fmt.Printf("%s attacks %s for 10 damage! %s has %d HP left.\n", c.name, monster.name, monster.name, monster.currHP)
}
