package main

type monster struct {
	name                  string
	maxHP, currHP, attack int
}

func initGoblin() *monster {
	return &monster{"Training goblin", 40, 40, 5}
}
