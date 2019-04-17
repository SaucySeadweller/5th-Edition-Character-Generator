package main

type Character struct {
	Name             string
	Class            string
	Race             string
	Health           int
	Skills           []Skill
	AC               int
	Currency         []int
	CarryingCapacity int
	Age              int
	Feats            []Feat
	Sex              string
	AbilityScores    AbilityScores
}

func NewCharacter() *Character {
	sex := pickSex()
	c := &Character{
		Name: generateName(sex),
		Race: pickRace(),
		//	Age:           RaceAge(),
		Sex:           sex,
		Class:         pickClass(),
		AbilityScores: generateAbilityScores(),
	}
	return c
}
