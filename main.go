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
	newInventory := &inventory{[]potion{{"pot1", nil}, {"pot2", nil}, {"pot3", nil}}}
	char1 = initChar("newCharacter", "elf", 100, newInventory)
	mainMenu()
}

func mainMenu() {
	options := []menuOption{
		{"Show character's info", char1.displayInfo},
		{"Show inventory", char1.accessInventory},
	}
	menuPrint(options)
}
