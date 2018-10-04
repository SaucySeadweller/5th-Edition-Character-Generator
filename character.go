package main

type Character struct {
	CharacterName             string
	Class            string
	Race             string
	Health           int
	Skills           []Skill
	AC               int
	Currency         []int
	CarryingCapacity int
	Age              int
	Feats            []Feat
}
