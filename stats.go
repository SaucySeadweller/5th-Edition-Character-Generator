package main

import (
	"fmt"
	"math/rand"
	"time"
)

//creates a random number for Ablitity Scores.
func generateNum() int {
	return rand.Intn(17) + 1
}

func generateAge() int {
	rand.Seed(time.Now().UnixNano())
	min := 18
	max := 79
	fmt.Println(rand.Intn(max-min) + min)
	//	switch min {
	//	case "human":
	//		return rand.Intn(54) + 16
	//	case "Elf":
	//		return rand.Intn(300) + 60
	//	case "Dwarf":
	//		return rand.Intn(50) +
	//	case "Duergar":
	//		return rand.Intn(54) + 16
	//	case "Gnome":
	//		return rand.Intn(54) + 16
	//	case "Hobgoblin":
	//		return rand.Intn(54) + 16
	//	case "Goblin":
	//		return rand.Intn(54) + 16
	//	case "Halfling":
	//		return rand.Intn(54) + 16
	//	case "Lizardfolk":
	//		return rand.Intn(54) + 16
	//	case "Tortle":
	//		return rand.Intn(54) + 16
	//	case "Aasimar":
	//		return rand.Intn(54) + 16
	//	case "Fallen Aasimar":
	//		return rand.Intn(54) + 16
	//	case "Protector Aasimar":
	//		return rand.Intn(54) + 16
	//	case "Scourge Aasimar":
	//		return rand.Intn(54) + 16
	//	case "Mountain Dwarf":
	//		return rand.Intn(54) + 16
	//	case "Hill Dwarf":
	//		return rand.Intn(54) + 16
	//	case "Tiefling":
	//		return rand.Intn(54) + 16
	//	case "Dragonborn":
	//		return rand.Intn(54) + 16
	//	case "Orc":
	//		return rand.Intn(54) + 16
	//	case "Half-Orc":
	//		return rand.Intn(54) + 16
	//	case "Half-Elf":
	//		return rand.Intn(54) + 16
	//	case "Aarakocra":
	//		return rand.Intn(54) + 16
	//	case "Eldarrin":
	//		return rand.Intn(54) + 16
	//	case "Tabaxi":
	//		return rand.Intn(54) + 16
	//	case "Triton":
	//		return rand.Intn(54) + 16
	//	case "Shadar-Kai":
	//		return rand.Intn(54) + 16
	//	case "Wood Elf":
	//		return rand.Intn(54) + 16
	//	case "Drow":
	//		return rand.Intn(54) + 16
	//	case "High Elf":
	//		return rand.Intn(54) + 16
	//	case "Kenku":
	//		return rand.Intn(54) + 16
	//	case "Githyanki":
	//		return rand.Intn(54) + 16
	//	case "Genasi":
	//		return rand.Intn(54) + 16
	//	case "Githzerai":
	//		return rand.Intn(54) + 16
	//	case "Goliath":
	//		return rand.Intn(54) + 16
	//	case "kobold":
	//		return rand.Intn(54) + 16
	//	case "Yuan-Ti Pureblood":
	//		return rand.Intn(54) + 16
	//	case "Bugbear":
	//		return rand.Intn(54) + 16
	//	case "Sea Elf":
	//		return rand.Intn(54) + 16
	//	case "Changling":
	//		return rand.Intn(54) + 16
	//	default:
	//		panic("race doesnt exist")
	//	}
	return 0
}
func pickClass() string {
	class := []string{
		"Rogue", "Sorcerer", "Wizard", "Paladin", "Barbarian", "Fighter", "Cleric", "Warlock", "Ranger", "Bard", "Monk"}
	return class[rand.Intn(len(class)-1)]
}

func generateAbilityScores() AbilityScores {

	return AbilityScores{
		Strength:     rand.Intn(17) + 1,
		Dexterity:    rand.Intn(17) + 1,
		Constitution: rand.Intn(17) + 1,
		Intelligence: rand.Intn(17) + 1,
		Wisdom:       rand.Intn(17) + 1,
		Charisma:     rand.Intn(17) + 1,
	}
}

func AbilityBonuses() {

}
func DetermineHP() {

}
