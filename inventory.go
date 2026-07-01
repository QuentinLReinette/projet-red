package main

import (
	"fmt"
)

type inventory struct {
	potions []*potion
}

func (inv *inventory) accessInventory() {
	options := make([]menuOption, len(inv.potions))
	for i, pot := range inv.potions {
		options[i] = menuOption{fmt.Sprintf("%s potion (+50 HP) x%d", pot, pot.quantity), pot.action}
	}
	menuPrint(options, false)
}
