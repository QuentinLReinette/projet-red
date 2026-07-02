package main

import "fmt"

func trainingFight(char *character) {
	goblin := initGoblin()
	fight(char, goblin)
}

type fighter struct {
	name  string
	fight func()
	next  *fighter
}

func fight(char *character, monster *monster) {
	round := 1
	charFighter := &fighter{name: char.name, fight: func() { charTurn(char, monster) }}
	monsterFighter := &fighter{name: monster.name, fight: func() { monster.goblinPattern(round, char) }}

	charFighter.next, monsterFighter.next = monsterFighter, charFighter

	currentFighter := charFighter
	if char.initiative < monster.initiative {
		currentFighter = monsterFighter
	}
	fmt.Printf("%s takes the initiative!\n", currentFighter.name)

	for ; char.currHP > 0 && monster.currHP > 0; round++ {
		fmt.Printf("Round %d\n", round)

		currentFighter.fight()
		if !char.checkAlive() {
			return
		}
		currentFighter = currentFighter.next

		currentFighter.fight()
		if !char.checkAlive() {
			return
		}
		currentFighter = currentFighter.next
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
