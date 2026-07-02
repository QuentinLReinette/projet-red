package main

import "fmt"

type blacksmith struct {
	char *character
}

func blacksmithMenu(char *character) {
	bs := blacksmith{char}
	options := []menuOption{
		{"Adventurer's hat - 5 g", bs.craftHat},
		{"Adventurer's vest - 5 g", bs.craftVest},
		{"Adventurer's boots - 5 g", bs.craftBoots},
	}
	menuPrint(options, false)
}

func (b blacksmith) checkReqs(materials map[materialType]int, gold int) bool {
	if !b.char.inventory.checkHasMaterials(materials) {
		println("You don't have the required materials to craft this item.")
		for m, qtty := range materials {
			fmt.Printf("%s x%d\n", m, qtty)
		}
		return false
	}
	if b.char.gold < gold {
		fmt.Printf("You don't have enough gold to craft this item (%d).\n", gold)
		return false
	}
	return true
}

func (b blacksmith) craftItem(materials map[materialType]int, gold int, itemName string) {
	if !b.checkReqs(materials, gold) {
		return
	}
	b.char.inventory.removeMaterials(materials)
	b.char.gold -= gold
	fmt.Printf("You crafted a %s.\n", itemName)
}

func (b blacksmith) craftHat() {
	b.craftItem(map[materialType]int{crowFeather: 1, boarLeather: 1}, 5, "Adventurer's hat")
}

func (b blacksmith) craftVest() {
	b.craftItem(map[materialType]int{wolfHide: 2, trollHide: 1}, 5, "Adventurer's vest")
}

func (b blacksmith) craftBoots() {
	b.craftItem(map[materialType]int{wolfHide: 1, boarLeather: 1}, 5, "Adventurer's boots")
}
