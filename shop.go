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
	defer println("You bought 1 Health potion.")

	for _, p := range s.char.potions {
		if p.potionType == healthPotion {
			p.quantity++
			return
		}
	}
	s.char.potions = append(s.char.potions, newHealthPot(s.char, 1))
}

func (s shop) addPoisonPot() {
	defer println("You bought 1 Poison potion.")

	for _, p := range s.char.potions {
		if p.potionType == poisonPotion {
			p.quantity++
			return
		}
	}
	s.char.potions = append(s.char.potions, newPoisonPot(s.char, 1))
}

func (s shop) addFireBall() {
	char := s.char
	if char.checkHasSkill(fireBall) {
		fmt.Printf("%s already knows %s.", char.name, fireBall)
		return
	}
	char.inventory.books = append(char.inventory.books, &book{fireBall, func() { char.spellBook(fireBall) }})
}
