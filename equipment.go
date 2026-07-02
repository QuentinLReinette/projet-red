package main

import "slices"

type equipmentType int

const (
	headGear equipmentType = iota
	bodyArmor
	legArmor
)

type equipment struct {
	equipmentType
	name    string
	defense int
	*character
}

func (e *equipment) equip() {
	for _, eq := range e.character.equipment {
		if eq.equipmentType == e.equipmentType {
			e.unequip()
			break
		}
	}
	e.character.equipment = append(e.character.equipment, e)
	e.character.inventory.removeEquipment(e)
	e.character.maxHP += e.defense
}

func (e *equipment) unequip() {
	for idx, eq := range e.character.equipment {
		if eq == e {
			e.character.inventory.addEquipment(e)
			e.character.equipment = slices.Delete(e.character.equipment, idx, idx+1)
			break
		}
	}
	e.character.maxHP = max(e.character.maxHP-e.defense, 1)
}

func newEquipment(equipmentType equipmentType, name string, defense int, char *character) *equipment {
	return &equipment{equipmentType, name, defense, char}
}
