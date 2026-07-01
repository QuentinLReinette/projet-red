package main

import (
	"fmt"
	"math"
	"slices"
)

type potion struct {
	potionName string
	quantity   int
	action     func()
	*character
}

func newHealthPot(char *character, quantity int) *potion {
	newPot := &potion{"Health", quantity, nil, char}
	newPot.action = newPot.heal
	return newPot
}

func (p *potion) heal() {
	hp, maxHP := &p.character.currHP, &p.character.maxHP
	*hp = int(math.Min(float64(*hp+50), float64(*maxHP)))
	p.potionRemove()
	fmt.Printf("You healed 50 HP. Current health: %d/%d\n", p.character.currHP, *maxHP)
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
	return p.potionName == other.potionName && p.character == other.character
}
