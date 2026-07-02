package main

import "fmt"

func trainingFight() {
	goblin := initGoblin()
	fight(goblin)
}

type fighter struct {
	name  string
	fight func()
	next  *fighter
}

func fight(monster *monster) {
	round := 1
	charFighter := &fighter{name: Game.character.name, fight: func() { charTurn(monster) }}
	monsterFighter := &fighter{name: monster.name, fight: func() { monster.goblinPattern(round) }}

	charFighter.next, monsterFighter.next = monsterFighter, charFighter

	currentFighter := charFighter
	if Game.character.initiative < monster.initiative {
		currentFighter = monsterFighter
	}
	fmt.Printf("%s takes the initiative!\n", currentFighter.name)

	for ; Game.character.currHP > 0 && monster.currHP > 0; round++ {
		fmt.Printf("Round %d\n", round)

		currentFighter.fight()
		if !Game.character.checkAlive() {
			return
		}
		if monster.currHP <= 0 {
			break
		}
		currentFighter = currentFighter.next

		currentFighter.fight()
		if !Game.character.checkAlive() {
			return
		}
		if monster.currHP <= 0 {
			break
		}
		currentFighter = currentFighter.next
	}
	println("You defeated the monster!")
	Game.character.xp += monster.xp
	Game.character.checkXP()
}

func charTurn(monster *monster) {
	options := []menuOption{
		{"Open Inventory", Game.character.accessInventory},
		{"Attack", func() {
			Game.character.attack(monster)
		}},
	}
	menuPrint(options, false, false)
}
