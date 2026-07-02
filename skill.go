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

func newPunch() skill {
	newSkill := skill{skillType: punch, mpCost: 5, dmg: 8}
	return newSkill
}

func newFireBall() skill {
	newSkill := skill{skillType: fireBall, mpCost: 20, dmg: 18}
	return newSkill
}

func (s skill) skillUse(target *monster) {
	if Game.character.currMP < s.mpCost {
		fmt.Printf("Not enough MP to use %s\n", s)
		return
	}
	Game.character.currMP -= s.mpCost
	target.currHP -= s.dmg
	fmt.Printf("%s uses %s! It does %d damage to %s. %s has %d HP left.\n", Game.character.name, s, s.dmg, target.name, target.name, target.currHP)
}
