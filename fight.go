package main

import "fmt"

type fight struct{
	round int
	*monster
}

func (f *fight) trainingFight() {
	goblin := initGoblin()
	f.fightLoop(goblin)
}

type fighter struct {
	name  string
	fight func()
	next  *fighter
}

func (f *fight) fightLoop(monster *monster) {
	f.round = 1
	f.monster = monster
	charFighter := &fighter{name: Game.character.name, fight: f.charTurn}
	monsterFighter := &fighter{name: monster.name, fight: func() { monster.goblinPattern(f.round) }}

	charFighter.next, monsterFighter.next = monsterFighter, charFighter

	currentFighter := charFighter
	if Game.character.initiative < monster.initiative {
		currentFighter = monsterFighter
	}
	fmt.Printf("%s takes the initiative!\n", currentFighter.name)

	for ; Game.character.currHP > 0 && monster.currHP > 0; f.round++ {
		println("================================================================")
		fmt.Printf("Round %d\n", f.round)

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

func (f *fight) charTurn() {
	options := []menuOption{
		{"Open Inventory", Game.character.accessInventory},
		{"Attack", func() {
			Game.character.attack(f.monster)
		}},
	}
	menuPrint(options, false, false)
}
