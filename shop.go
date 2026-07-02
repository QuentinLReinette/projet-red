package main

import (
	"fmt"
)

type shop struct {
	char *character
}

func shopMenu(char *character) {
	shop := shop{char}
	options := []menuOption{
		{"Health potion (+50 HP) - free", shop.addHealthPot},
		{"Poison potion (-10 HP/s) - free", shop.addPoisonPot},
		{"Spell book: fireball", shop.addFireBall},
	}
	menuPrint(options, false)
}

func (s shop) addHealthPot() {
	if s.char.inventory.isFull() {
		return
	}
	s.char.inventory.addPotion(newHealthPot(s.char, 1))
	println("You bought 1 Health potion.")
}

func (s shop) addPoisonPot() {
	if s.char.inventory.isFull() {
		return
	}
	s.char.inventory.addPotion(newPoisonPot(s.char, 1))
	println("You bought 1 Poison potion.")
}

func (s shop) addFireBall() {
	if s.char.inventory.isFull() {
		return
	}
	char := s.char
	if char.checkHasSkill(fireBall) {
		fmt.Printf("%s already knows Fireball.", char.name)
		return
	}
	for _, b := range char.inventory.books {
		if b.skillType == fireBall {
			fmt.Printf("%s already has a Fireball spell book.\n", char.name)
			return
		}
	}
	s.char.inventory.addBook(&book{fireBall, func() { char.spellBook(fireBall) }})
	println("You bought a Fireball spell book.")
}
