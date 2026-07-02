package main

import "fmt"

func trainingFight(char *character) {
	goblin := initGoblin()
	fight(char, goblin)
}

func fight(char *character, monster *monster) {
	for round := 1; char.currHP > 0 && monster.currHP > 0; round++ {
		monsterAttack := monster.goblinPattern(round)
		char.currHP -= monsterAttack
		fmt.Printf("Round %d\nGoblin attacks for %d damage! You have %d HP left.\n", round, monsterAttack, char.currHP)
		if !char.checkAlive() {
			return
		}
		charTurn(char, monster)
	}
	if monster.currHP <= 0 {
		fmt.Println("You defeated the monster!")
		return
	}
}

func charTurn(char *character, monster *monster) {
	options := []menuOption{
		{"Open Inventory", char.accessInventory},
		{"Attack", func() {
			monster.currHP -= 5
			fmt.Printf("%s attacks the %s for 5 damage! The monster has %d HP left.\n", char.name, monster.name, monster.currHP)
		}},
	}
	menuPrint(options, false)
}
