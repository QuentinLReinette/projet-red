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
		{"Health potion (+50 HP) - 3 g", shop.addHealthPot},
		{"Poison potion (-10 HP/s) - 6 g", shop.addPoisonPot},
		{"Spell book: fireball - 25 g", shop.addFireBall},
		{"Wolf hide - 4 g", shop.addWolfHide},
		{"Troll hide - 7 g", shop.addTrollHide},
		{"Boar leather - 3 g", shop.addBoarLeather},
		{"Crow feather - 1 g", shop.addCrowFeather},
	}
	menuPrint(options, false)
}

func (s shop) addHealthPot() {
	if s.char.inventory.isFull() {
		return
	}
	s.char.inventory.addPotion(newHealthPot(s.char, 1))
	s.char.gold -= 3
	println("You bought 1 Health potion.")
}

func (s shop) addPoisonPot() {
	if s.char.inventory.isFull() {
		return
	}
	s.char.inventory.addPotion(newPoisonPot(s.char, 1))
	s.char.gold -= 6
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
	s.char.gold -= 25
	s.char.inventory.addBook(&book{fireBall, func() { char.spellBook(fireBall) }})
	println("You bought a Fireball spell book.")
}

func (s shop) addWolfHide() {
	if s.char.inventory.isFull() {
		return
	}
	s.char.inventory.addMaterial(wolfHide)
	s.char.gold -= 4
	println("You bought 1 Wolf Hide.")
}

func (s shop) addTrollHide() {
	if s.char.inventory.isFull() {
		return
	}
	s.char.inventory.addMaterial(trollHide)
	s.char.gold -= 7
	println("You bought 1 Troll Hide.")
}

func (s shop) addBoarLeather() {
	if s.char.inventory.isFull() {
		return
	}
	s.char.inventory.addMaterial(boarLeather)
	s.char.gold -= 3
	println("You bought 1 Boar Leather.")
}

func (s shop) addCrowFeather() {
	if s.char.inventory.isFull() {
		return
	}
	s.char.inventory.addMaterial(crowFeather)
	s.char.gold -= 1
	println("You bought 1 Crow Feather.")
}
