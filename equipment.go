package main

import "slices"

type equipmentType int

const (
	headGear equipmentType = iota
	bodyArmor
	legArmor
)

var equipmentTypeNames = map[equipmentType]string{
	headGear:  "Head Gear",
	bodyArmor: "Body Armor",
	legArmor:  "Leg Armor",
}

type equipment struct {
	equipmentType
	name    string
	defense int
}

func (e equipmentType) String() string {
	return equipmentTypeNames[e]
}

func (e *equipment) equip() {
	for _, eq := range Game.character.equipment {
		if eq.equipmentType == e.equipmentType {
			eq.unequip()
			break
		}
	}
	Game.character.equipment = append(Game.character.equipment, e)
	Game.character.inventory.removeEquipment(e)
	Game.character.maxHP += e.defense
}

func (e *equipment) unequip() {
	for idx, eq := range Game.character.equipment {
		if eq == e {
			Game.character.inventory.addEquipment(e)
			Game.character.equipment = slices.Delete(Game.character.equipment, idx, idx+1)
			break
		}
	}
	Game.character.maxHP = max(Game.character.maxHP-e.defense, 1)
}

func newEquipment(equipmentType equipmentType, name string, defense int) *equipment {
	return &equipment{equipmentType, name, defense}
}
