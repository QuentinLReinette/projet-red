# Projet Red

A text-based RPG written in Go.

## Quick Start

### Prerequisites

- Go 1.21+

### Running the Game

```bash
go run .
```

## Overview

### Character & Classes
Choose a character name and class (Human, Elf, Dwarf) with distinct health, mana, and base stats.

### Combat & Progression
- Battle monsters in initiative-based turn combat.
- Cast active spells (Punch, Fireball), use items, or deal physical attacks.
- Gain XP to raise level, max HP, and max mana.

### Items, Shop & Crafting
- **Inventory**: Slot-limited storage for potions, equipment, materials, and spellbooks.
- **Shop**: Buy consumables, spellbooks, or expand inventory capacity with gold.
- **Blacksmith**: Craft headgear, body armor, and boots from gathered materials.

## Codebase Structure

- `main.go`, `gameState.go`, `menu.go`: Main execution loop, state initialization, and CLI navigation.
- `character.go`, `class.go`, `monster.go`: Player attributes, class templates, leveling, and monster behaviors.
- `fight.go`, `skill.go`: Combat loops, initiative tracking, and active spells.
- `inventory.go`, `potion.go`, `equipment.go`, `material.go`: Storage limits, consumables, armor stats, and materials.
- `shop.go`, `blacksmith.go`: Item purchasing, inventory upgrades, and crafting recipes.
