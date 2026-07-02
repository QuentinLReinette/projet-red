package main

import (
	"fmt"
	"math"
	"slices"
	"time"
)

type potionType int

const (
	healthPotion potionType = iota
	manaPotion
	poisonPotion
)

var potionNames = map[potionType]string{
	healthPotion: "Health",
	manaPotion:   "Mana",
	poisonPotion: "Poison",
}

type potion struct {
	potionType potionType
	quantity   int
	action     func()
	*character
}

func (p potion) String() string {
	return potionNames[p.potionType]
}

func newHealthPot(char *character, quantity int) *potion {
	newPot := &potion{healthPotion, quantity, nil, char}
	newPot.action = newPot.heal
	return newPot
}

func newPoisonPot(char *character, quantity int) *potion {
	newPot := &potion{poisonPotion, quantity, nil, char}
	newPot.action = newPot.poison
	return newPot
}

func newManaPot(char *character, quantity int) *potion {
	newPot := &potion{potionType: manaPotion, quantity: quantity, character: char}
	newPot.action = newPot.mana
	return newPot
}

func (p *potion) heal() {
	hp, maxHP := &p.character.currHP, &p.character.maxHP
	*hp = int(math.Min(float64(*hp+50), float64(*maxHP)))
	p.potionRemove()
	fmt.Printf("You healed 50 HP. Current health: %d/%d\n", p.character.currHP, *maxHP)
}

func (p *potion) poison() {
	char := p.character
	for range 3 {
		fmt.Printf("%s loses 10 HP.\n", char.name)
		char.currHP -= 10
		time.Sleep(1 * time.Second)
	}
	p.potionRemove()
	fmt.Printf("%s's health: %d/%d\n", char.name, char.currHP, char.maxHP)
}

func (p *potion) mana() {
	mp, maxMP := &p.character.currMP, &p.character.maxMP
	*mp = int(math.Min(float64(*mp+30), float64(*maxMP)))
	p.potionRemove()
	fmt.Printf("You restored 30 MP. Current mana: %d/%d\n", p.character.currMP, *maxMP)
}

func (p *potion) potionRemove() {
	p.quantity--

	if p.quantity <= 0 {
		idx := slices.IndexFunc(p.character.potions, func(q *potion) bool {
			return q.potionEquals(p)
		})
		p.character.potions = slices.Delete(p.character.potions, idx, idx+1)
	}
}

func (p *potion) potionEquals(other *potion) bool {
	return p.potionType == other.potionType && p.character == other.character
}
