package main

import "fmt"

type skillType int

const (
	punch skillType = iota
	fireBall
)

var skillNames = map[skillType]string{
	punch:    "Punch",
	fireBall: "Fireball",
}

type skill struct {
	skillType
	char        *character
	mpCost, dmg int
}

type book struct {
	skillType
	learn func()
}

func (s skill) String() string {
	return skillNames[s.skillType]
}

func (t skillType) String() string {
	return skillNames[t]
}

func newPunch(char *character) skill {
	newSkill := skill{skillType: punch, char: char, mpCost: 5, dmg: 8}
	return newSkill
}

func newFireBall(char *character) skill {
	newSkill := skill{skillType: fireBall, char: char, mpCost: 20, dmg: 18}
	return newSkill
}

func (s skill) skillUse(target *monster) {
	if s.char.currMP < s.mpCost {
		fmt.Printf("Not enough MP to use %s\n", s)
		return
	}
	s.char.currMP -= s.mpCost
	target.currHP -= s.dmg
	fmt.Printf("%s uses %s! It does %d damage to %s. %s has %d HP left.\n", s.char.name, s, s.dmg, target.name, target.name, target.currHP)
}
