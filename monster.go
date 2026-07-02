package main

type monster struct {
	name                  string
	maxHP, currHP, attack int
}

func initGoblin() *monster {
	return &monster{"Training goblin", 40, 40, 5}
}

func (m *monster) goblinPattern(round int) int {
	if round%3 == 0 {
		return m.attack * 2
	}
	return m.attack
}
