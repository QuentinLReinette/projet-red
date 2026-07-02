package main

import "fmt"

type blacksmith struct{}

func (b *blacksmith) blacksmithMenu() {
	options := []menuOption{
		{"Adventurer's hat - 5 g", b.craftHat},
		{"Adventurer's vest - 5 g", b.craftVest},
		{"Adventurer's boots - 5 g", b.craftBoots},
	}
	menuPrint(options, true, false)
}

func (b *blacksmith) checkReqs(materials map[materialType]int, gold int) bool {
	if !Game.character.inventory.checkHasMaterials(materials) {
		println("You don't have the required materials to craft this item.")
		for m, qtty := range materials {
			fmt.Printf("%s x%d\n", m, qtty)
		}
		return false
	}
	if Game.character.gold < gold {
		fmt.Printf("You don't have enough gold to craft this item (%d).\n", gold)
		return false
	}
	return true
}

func (b *blacksmith) craftItem(materials map[materialType]int, gold int, itemName string) {
	if !b.checkReqs(materials, gold) {
		return
	}
	Game.character.inventory.removeMaterials(materials)
	Game.character.gold -= gold
	fmt.Printf("You crafted %s.\n", itemName)
}

func (b *blacksmith) craftHat() {
	b.craftItem(map[materialType]int{crowFeather: 1, boarLeather: 1}, 5, "Adventurer's hat")
	newEq := newEquipment(headGear, "Adventurer's hat", 10, Game.character)
	Game.character.inventory.addEquipment(newEq)
}

func (b *blacksmith) craftVest() {
	b.craftItem(map[materialType]int{wolfHide: 2, trollHide: 1}, 5, "Adventurer's vest")
	newEq := newEquipment(bodyArmor, "Adventurer's vest", 25, Game.character)
	Game.character.inventory.addEquipment(newEq)
}

func (b *blacksmith) craftBoots() {
	b.craftItem(map[materialType]int{wolfHide: 1, boarLeather: 1}, 5, "Adventurer's boots")
	newEq := newEquipment(legArmor, "Adventurer's boots", 15, Game.character)
	Game.character.inventory.addEquipment(newEq)
}
