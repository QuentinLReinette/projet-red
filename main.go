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
		{"Show inventory", Game.accessInventory},
		{"Shop", Game.shopMenu},
		{"Blacksmith", Game.blacksmithMenu},
		{"Training fight", Game.trainingFight},
	}
	menuPrint(options, true, true)
}
