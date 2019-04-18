package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type LifeSpan struct {
	min, max int
}
type Race struct {
	RaceName         string
	RacialScoreBonus int
	RacialTrait      string
	Size             string
	Weight           int
	Height           int
	Speed            int
	Languages        []string
	Abilityscores    map[string]int
}

var ageRange = map[string]LifeSpan{
	"Human":             {16, 75},
	"Dwarf":             {50, 350},
	"Duergar":           {50, 350},
	"Mountain Dwarf":    {50, 350},
	"Hill Dwarf":        {50, 350},
	"Elf":               {100, 750},
	"Half-Orc":          {14, 75},
	"Half-Elf":          {20, 180},
	"Halfling":          {20, 150},
	"Aarakocra":         {3, 30},
	"Aasimar":           {16, 160},
	"Fallen Aasimar":    {16, 160},
	"Protector Aasimar": {16, 160},
	"Scourge Aasimar":   {16, 160},
	"DragonBorn":        {15, 80},
	"Eladrin":           {100, 750},
	"Sea Elf":           {100, 750},
	"High Elf":          {100, 750},
	"Shadar-Kai":        {20, 1000},
	"Wood Elf":          {100, 750},
	"Genasi":            {16, 120},
	"Githyanki":         {16, 100},
	"Githzerai":         {16, 100},
	"Gnome":             {40, 350},
	"Goliath":           {16, 75},
	"Kenku":             {12, 60},
	"Kobold":            {6, 120},
	"Lizardfolk":        {14, 60},
	"Orc":               {12, 50},
	"Tabaxi":            {16, 75},
	"Bugbear":           {16, 80},
}

func raceAge() {
	fmt.Println(ageRange)
}
func pickRace() string {
	Race := []string{
		"Human", "Dwarf", "Duergar", "Mountain Dwarf", "Hill Dwarf", "Elf", "Half-Orc", "Half-Elf", "Halfling", "Aarakocra",
		" Aasimar", "Fallen Aasimar", "Protector Aasimar", "Scourge Aasimar", "DragonBorn", "Drow", "Eladrin", "Sea Elf", "High Elf",
		"Shadar-Kai", "Wood Elf", "Genasi", "Githyanki", "Githzerai", "Gnome", "Goliath", "Kenku", "Kobold", "Lizardfolk", "Orc", "Tabaxi", "Triton",
	}
	return Race[rand.Intn(len(Race)-1)]

}

func pickSex() string {
	Sex := []string{"Male", "Female"}

	return Sex[rand.Intn(2)]
}
