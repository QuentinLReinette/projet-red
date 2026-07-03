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
	println("================================================================")
	println("Welcome to the shop! What would you like to buy?")
	menuPrint(options, true, false)
}

func (s shop) checkCanBuy(cost int) bool {
	if Game.character.gold < cost {
		println("You don't have enough gold.")
		return false
	}
	if Game.character.inventory.isFull() {
		return false
	}
	return true
}

func (s shop) addHealthPot() {
	if !s.checkCanBuy(3) {
		return
	}
	Game.character.inventory.addPotion(newHealthPot(1))
	Game.character.gold -= 3
	println("You bought 1 Health potion.")
}

func (s shop) addManaPot() {
	if !s.checkCanBuy(5) {
		return
	}
	Game.character.inventory.addPotion(newManaPot(1))
	Game.character.gold -= 5
	println("You bought 1 Mana potion.")
}

func (s shop) addPoisonPot() {
	if !s.checkCanBuy(7) {
		return
	}
	Game.character.inventory.addPotion(newPoisonPot(1))
	Game.character.gold -= 7
	println("You bought 1 Poison potion.")
}

func (s shop) addFireBall() {
	if !s.checkCanBuy(25) {
		return
	}
	char := Game.character
	if char.checkHasSkill(fireBall) {
		fmt.Printf("%s already knows Fireball.\n", char.name)
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
	if !s.checkCanBuy(4) {
		return
	}
	Game.character.inventory.addMaterial(wolfHide)
	Game.character.gold -= 4
	println("You bought 1 Wolf Hide.")
}

func (s shop) addTrollHide() {
	if !s.checkCanBuy(7) {
		return
	}
	Game.character.inventory.addMaterial(trollHide)
	Game.character.gold -= 7
	println("You bought 1 Troll Hide.")
}

func (s shop) addBoarLeather() {
	if !s.checkCanBuy(3) {
		return
	}
	Game.character.inventory.addMaterial(boarLeather)
	Game.character.gold -= 3
	println("You bought 1 Boar Leather.")
}

func (s shop) addCrowFeather() {
	if !s.checkCanBuy(1) {
		return
	}
	Game.character.inventory.addMaterial(crowFeather)
	Game.character.gold -= 1
	println("You bought 1 Crow Feather.")
}

func (s shop) upgradeInventory() {
	if Game.character.gold < 30 {
		println("You don't have enough gold.")
		return
	}
	if Game.character.inventory.capacity >= 40 {
		println("Your inventory is already at maximum capacity.")
		return
	}
	Game.character.inventory.capacity += 10
	Game.character.gold -= 30
	println("Your inventory capacity has been increased by 10.")
}
