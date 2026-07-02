package main

import (
	"fmt"
)

type shop struct{}

func (s shop) shopMenu() {
	options := []menuOption{
		{"Health potion (+50 HP) - 3 g", s.addHealthPot},
		{"Mana potion (+30 MP) - 5 g", s.addManaPot},
		{"Poison potion (-10 HP/s) - 7 g", s.addPoisonPot},
		{"Spell book: fireball - 25 g", s.addFireBall},
		{"Wolf hide - 4 g", s.addWolfHide},
		{"Troll hide - 7 g", s.addTrollHide},
		{"Boar leather - 3 g", s.addBoarLeather},
		{"Crow feather - 1 g", s.addCrowFeather},
		{"Upgrade inventory capacity - 30 g", s.upgradeInventory},
	}
	menuPrint(options, true, false)
}

func (s shop) addHealthPot() {
	if Game.character.inventory.isFull() {
		return
	}
	Game.character.inventory.addPotion(newHealthPot(Game.character, 1))
	Game.character.gold -= 3
	println("You bought 1 Health potion.")
}

func (s shop) addManaPot() {
	if Game.character.inventory.isFull() {
		return
	}
	Game.character.inventory.addPotion(newManaPot(Game.character, 1))
	Game.character.gold -= 5
	println("You bought 1 Mana potion.")
}

func (s shop) addPoisonPot() {
	if Game.character.inventory.isFull() {
		return
	}
	Game.character.inventory.addPotion(newPoisonPot(Game.character, 1))
	Game.character.gold -= 7
	println("You bought 1 Poison potion.")
}

func (s shop) addFireBall() {
	if Game.character.inventory.isFull() {
		return
	}
	char := Game.character
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
	Game.character.gold -= 25
	Game.character.inventory.addBook(&book{fireBall, func() { char.spellBook(fireBall) }})
	println("You bought a Fireball spell book.")
}

func (s shop) addWolfHide() {
	if Game.character.inventory.isFull() {
		return
	}
	Game.character.inventory.addMaterial(wolfHide)
	Game.character.gold -= 4
	println("You bought 1 Wolf Hide.")
}

func (s shop) addTrollHide() {
	if Game.character.inventory.isFull() {
		return
	}
	Game.character.inventory.addMaterial(trollHide)
	Game.character.gold -= 7
	println("You bought 1 Troll Hide.")
}

func (s shop) addBoarLeather() {
	if Game.character.inventory.isFull() {
		return
	}
	Game.character.inventory.addMaterial(boarLeather)
	Game.character.gold -= 3
	println("You bought 1 Boar Leather.")
}

func (s shop) addCrowFeather() {
	if Game.character.inventory.isFull() {
		return
	}
	Game.character.inventory.addMaterial(crowFeather)
	Game.character.gold -= 1
	println("You bought 1 Crow Feather.")
}

func (s shop) upgradeInventory() {
	if Game.character.inventory.capacity >= 40 {
		println("Your inventory is already at maximum capacity.")
		return
	}
	Game.character.inventory.capacity += 10
	Game.character.gold -= 30
	println("Your inventory capacity has been increased by 10.")
}
