package main

import (
	"fmt"
	"slices"
)

type inventory struct {
	potions   []*potion
	books     []*book
	equipment []*equipment
	materials []*material
	capacity  int
}

func (inv *inventory) accessInventory() {
	options := []menuOption{}
	for _, pot := range inv.potions {
		options = append(options, menuOption{fmt.Sprintf("%s potion (+50 HP) x%d", pot, pot.quantity), pot.action})
	}
	for _, book := range inv.books {
		options = append(options, menuOption{fmt.Sprintf("Spell book: %s", book), book.learn})
	}
	for _, mat := range inv.materials {
		options = append(options, menuOption{fmt.Sprintf("%s x%d", mat, mat.quantity), mat.use})
	}
	if len(options) == 0 {
		println("Your inventory is empty.")
		return
	}
	menuPrint(options, true, false)
}

func (i *inventory) addPotion(p *potion) {
	if i.isFull() {
		return
	}
	for _, pot := range i.potions {
		if pot.potionType == p.potionType {
			pot.quantity++
			return
		}
	}
	i.potions = append(i.potions, p)
}

func (i *inventory) addBook(b *book) {
	if i.isFull() {
		return
	}
	for _, book := range i.books {
		if book.skillType == b.skillType {
			return
		}
	}
	i.books = append(i.books, b)
}

func (i *inventory) isFull() bool {
	size := len(i.potions) + len(i.books) + len(i.materials) + len(i.equipment)
	if size > i.capacity {
		println("Your inventory is full. You can't carry more items.")
		return true
	}
	return false
}

func (i *inventory) addMaterial(m materialType) {
	if i.isFull() {
		return
	}
	for _, mat := range i.materials {
		if mat.materialType == m {
			mat.quantity++
			return
		}
	}
	i.materials = append(i.materials, &material{m, 1})
}

func (i *inventory) checkHasMaterials(materials map[materialType]int) bool {
	for m, quantity := range materials {
		found := false
		for _, mat := range i.materials {
			if mat.materialType == m {
				found = true
				if mat.quantity < quantity {
					return false
				}
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (i *inventory) removeMaterials(materials map[materialType]int) {
	for m, quantity := range materials {
		for idx, mat := range i.materials {
			if mat.materialType == m {
				mat.quantity -= quantity
				if mat.quantity <= 0 {
					i.materials = slices.Delete(i.materials, idx, idx+1)
				}
				break
			}
		}
	}
}

func (i *inventory) addEquipment(e *equipment) {
	if i.isFull() {
		return
	}
	for _, eq := range i.equipment {
		if eq.equipmentType == e.equipmentType {
			return
		}
	}
	i.equipment = append(i.equipment, e)
}

func (i *inventory) removeEquipment(e *equipment) {
	for idx, eq := range i.equipment {
		if eq == e {
			i.equipment = slices.Delete(i.equipment, idx, idx+1)
			break
		}
	}
}
