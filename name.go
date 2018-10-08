package main

import (
	"math/rand"
	"time"
)

var Nouns = []string{"eye",
	"crest", "crown", "tooth", "fang", "horn", "frill", "skull", "bone", "tongue", "throat", "voice", "nose", "snout", "chin", "eye", "sight", "seer", "speaker", "singer", "song",
	"chanter", "howler", "chatter", "shrieker", "shriek", "jaw", "bite", "biter", "neck", "shoulder", "fin", "wing", "arm", "lifter", "grasp", "grabber", "hand", "paw", "foot", "finger",
	"toe", "thumb", "talon", "palm", "touch", "racer", "runner", "hoof", "fly", "flier", "swoop", "roar", "hiss", "hisser", "snarl", "dive",
}

func randomNoun() string {
	return Nouns[rand.Intn(len(Nouns)-1)]
}

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
func generateName() string {
	rand.Seed(time.Now().UnixNano())
	Noun1 := uppercaseFirstLetter(randomNoun())
	Noun2 := uppercaseFirstLetter(randomNoun())
	Adjective := lowercaseFirstLetter(randomAdjective())
	return Noun1 + Adjective + " " + Noun2
}
