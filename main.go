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
	makeDeafultChar()
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

}
