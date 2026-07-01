package main

import (
	"bufio"
	"os"
)

var (
	char1   *character
	Scanner = bufio.NewScanner(os.Stdin)
)

func main() {
	makeDeafultChar()
	for {
		mainMenu()
	}
}

func makeDeafultChar() {
	char1 = initChar("newCharacter", "elf", 100, &inventory{})
	char1.inventory.potions = []*potion{newHealthPot(char1, 3)}
	char1.skills = []skill{newPunch(char1)}
}

func mainMenu() {
	options := []menuOption{
		{"Show character's info", char1.displayInfo},
		{"Show inventory", char1.accessInventory},
		{"Shop", func() { shopMenu(char1) }},
	}
	menuPrint(options, true)
}
