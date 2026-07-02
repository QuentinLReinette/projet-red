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
	defer func() {
		if err := Scanner.Err(); err != nil {
			println("Error while reading input: ", err)
		}
		Scanner = nil
	}()

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
	println("================================================================")
	println("What would you like to do?")
	menuPrint(options, true, true)
}
