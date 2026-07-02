package main

type monster struct {
	name                              string
	maxHP, currHP, attack, initiative, xp int
}

func initGoblin() *monster {
	return &monster{"Training Goblin", 40, 40, 5, 10, 10}
}

func (m *monster) goblinPattern(round int, char *character) {
	attack := m.attack
	if round%3 == 0 {
		attack *= 2
	}
	char.currHP -= attack
	println("Goblin attacks for", attack, "damage! You have", char.currHP, "HP left.")
}
