package main

type shop struct {
	char *character
}

func shopMenu(char *character) {
	shop := shop{char}
	options := []menuOption{
		{"Health potion (+50 HP) - free", shop.addHealthPot},
	}
	menuPrint(options, false)
}

func (s shop) addHealthPot() {
	defer println("You bought 1 Health potion.")

	for _, p := range s.char.potions {
		if p.potionType == healthPotion {
			p.quantity++
			return
		}
	}
	s.char.potions = append(s.char.potions, newHealthPot(s.char, 1))
}
