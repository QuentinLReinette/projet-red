package main

type monster struct {
	name                              string
	maxHP, currHP, attack, initiative, xp int
}

func initGoblin() *monster {
	return &monster{"Training Goblin", 40, 40, 5, 10, 10}
}

func (m *monster) goblinPattern(round int) {
	attack := m.attack
	if round%3 == 0 {
		attack *= 2
	}
	Game.character.currHP -= attack
	println("Goblin attacks for", attack, "damage! You have", Game.character.currHP, "HP left.")
}
