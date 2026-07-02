package main

import (
	"bufio"
	"os"
)

var (
	Game    *gameState
	char1   *character
	Scanner = bufio.NewScanner(os.Stdin)
)

func main() {
	Game = newGame()
	char1 = Game.character
	for {
		mainMenu()
	}
}

func mainMenu() {
	options := []menuOption{
		{"Show character's info", char1.displayInfo},
		{"Show inventory", char1.accessInventory},
		{"Shop", Game.shop.shopMenu},
		{"Blacksmith", func() { blacksmithMenu(char1) }},
		{"Training fight", func() { trainingFight(char1) }},
	}
	menuPrint(options, true, true)
}
