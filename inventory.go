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
	size := len(i.potions) + len(i.books)
	if size > 10 {
		println("Your inventory is full. You can't carry more items.")
		return true
	}
	return false
}
