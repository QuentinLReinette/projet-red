package main

import (
	"fmt"
)

type inventory struct {
	potions []*potion
	books   []*book
}

func (inv *inventory) accessInventory() {
	options := []menuOption{}
	for _, pot := range inv.potions {
		options = append(options, menuOption{fmt.Sprintf("%s potion (+50 HP) x%d", pot, pot.quantity), pot.action})
	}
	for _, book := range inv.books {
		options = append(options, menuOption{fmt.Sprintf("Spell book: %s", book), book.learn})
	}
	menuPrint(options, false)
}
