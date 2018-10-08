package main

import "math/rand"

var Adjectives = []string{"Black",
	"Grey",
	"Violet",
	"Vermillion",
	"Crimson",
	"Red",
	"Green",
	"Brown",
	"Bronze",
	"Yellow", "Orange", "Black", "Blue", "Purple", "White", "Pink", "Maroon", "Cyan", ""}

func randomAdjective() string {
	return Adjectives[rand.Intn(len(Adjectives)-1)]
}
