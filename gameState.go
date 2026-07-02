package main

import "strings"

type gameState struct {
	*character
	*shop
	*blacksmith
}

func newGame() *gameState {
	game := &gameState{&character{}, &shop{}, &blacksmith{}}
	game.charCreation()
	return game
}

// create a default character for testing purposes
func (gs *gameState) makeDefaultChar() {
	gs.character = newHuman("Default")
}

func (gs *gameState) charCreation() {
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
	newChar := &character{}
	options := []menuOption{
		{"Human", func() { newChar = newHuman(name) }},
		{"Elf", func() { newChar = newElf(name) }},
		{"Dwarf", func() { newChar = newDwarf(name) }},
	}
	menuPrint(options, false, false)

	println("Character created!")
	newChar.displayInfo()
	gs.character = newChar
}
