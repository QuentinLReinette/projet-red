package main

import (
	"bufio"
	"os"
	"strings"
)

var (
	char1   *character
	Scanner = bufio.NewScanner(os.Stdin)
)

func main() {
	charCreation()
	for {
		mainMenu()
	}
}

// create a default character for testing purposes
func makeDefaultChar() {
	char1 = newHuman("Default")
}

func mainMenu() {
	options := []menuOption{
		{"Show character's info", char1.displayInfo},
		{"Show inventory", char1.accessInventory},
		{"Shop", func() { shopMenu(char1) }},
		{"Blacksmith", func() { blacksmithMenu(char1) }},
		{"Training fight", func() { trainingFight(char1) }},
	}
	menuPrint(options, true, true)
}

func charCreation() {
	name := ""

nameFor:
	for {
		println("Enter a name for your character:")
		Scanner.Scan()
		name = strings.TrimSpace(Scanner.Text())
		if name == "" {
			println("The name can't be empty")
			continue
		}
		name = strings.ToLower(name)
		arr := []rune(name)
		for i, c := range arr {
			if i == 0 && c >= 'a' && c <= 'z' {
				arr[i] -= 'a' - 'A'
				name = string(arr)
			}
			if !(c >= 'a' && c <= 'z') && !(c >= 'A' && c <= 'Z') {
				println("The name can only contain letters")
				continue nameFor
			}
		}
		break
	}
	options := []menuOption{
		{"Human", func() { char1 = newHuman(name) }},
		{"Elf", func() { char1 = newElf(name) }},
		{"Dwarf", func() { char1 = newDwarf(name) }},
	}
	menuPrint(options, false, false)
	println("Character created!")
	char1.displayInfo()
}
