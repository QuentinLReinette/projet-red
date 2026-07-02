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

func makeDeafultChar() {
	char1 = initChar("newCharacter", elf, 100, &inventory{})
	char1.inventory.potions = []*potion{newHealthPot(char1, 3)}
	char1.skills = []skill{newPunch(char1)}
}

func mainMenu() {
	options := []menuOption{
		{"Show character's info", char1.displayInfo},
		{"Show inventory", char1.accessInventory},
		{"Shop", func() { shopMenu(char1) }},
		{"Blacksmith", func() { blacksmithMenu(char1) }},
	}
	menuPrint(options, true)
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

	for {
		println("\nChoose a class for your character:\n" +
			" 1. Human\n" +
			" 2. Elf\n" +
			" 3. Dwarf")
		Scanner.Scan()
		input := strings.TrimSpace(Scanner.Text())
		switch input {
		case "1":
			char1 = newHuman(name)
		case "2":
			char1 = newElf(name)
		case "3":
			char1 = newDwarf(name)
		default:
			println("Invalid input, please try again")
			continue
		}
		break
	}
	char1.inventory.potions = []*potion{newHealthPot(char1, 3)}
	char1.skills = []skill{newPunch(char1)}
}
