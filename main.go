package main

import (
	"bufio"
	"os"
)

var (
	Game    *gameState
	Scanner = bufio.NewScanner(os.Stdin)
)

func main() {
	Game = newGame()
	for {
		mainMenu()
	}
}

func mainMenu() {
	options := []menuOption{
		{"Show character's info", Game.character.displayInfo},
		{"Show inventory", Game.character.accessInventory},
		{"Shop", Game.shop.shopMenu},
		{"Blacksmith", Game.blacksmith.blacksmithMenu},
		{"Training fight", trainingFight},
	}
	menuPrint(options, true, true)
}
