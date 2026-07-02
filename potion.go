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
}

func (p potion) String() string {
	return potionNames[p.potionType]
}

func newHealthPot(quantity int) *potion {
	newPot := &potion{healthPotion, quantity, nil}
	newPot.action = newPot.heal
	return newPot
}

func newPoisonPot(quantity int) *potion {
	newPot := &potion{poisonPotion, quantity, nil}
	newPot.action = newPot.poison
	return newPot
}

func newManaPot(quantity int) *potion {
	newPot := &potion{manaPotion, quantity, nil}
	newPot.action = newPot.mana
	return newPot
}

func (p *potion) heal() {
	hp, maxHP := &Game.character.currHP, &Game.character.maxHP
	*hp = int(math.Min(float64(*hp+50), float64(*maxHP)))
	p.potionRemove()
	fmt.Printf("You healed 50 HP. Current health: %d/%d\n", Game.character.currHP, *maxHP)
}

func (p *potion) poison() {
	target := Game.fight.monster
	fmt.Printf("%s is poisoned! It will lose 10 HP for 3 seconds.\n", target.name)
	for range 3 {
		fmt.Printf("%s loses 10 HP.\n", target.name)
		target.currHP -= 10
		time.Sleep(1 * time.Second)
	}
	p.potionRemove()
	fmt.Printf("%s's health: %d/%d\n", target.name, target.currHP, target.maxHP)
}

func (p *potion) mana() {
	mp, maxMP := &Game.character.currMP, &Game.character.maxMP
	*mp = int(math.Min(float64(*mp+30), float64(*maxMP)))
	p.potionRemove()
	fmt.Printf("You restored 30 MP. Current mana: %d/%d\n", Game.character.currMP, *maxMP)
}

func (p *potion) potionRemove() {
	p.quantity--

	if p.quantity <= 0 {
		idx := slices.IndexFunc(Game.character.potions, func(q *potion) bool {
			return q.potionEquals(p)
		})
		Game.character.potions = slices.Delete(Game.character.potions, idx, idx+1)
	}
}

func (p *potion) potionEquals(other *potion) bool {
	return p.potionType == other.potionType
}
