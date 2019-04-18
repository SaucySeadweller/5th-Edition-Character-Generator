package main

//Character create a struct that holds all the values of a character
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

// NewCharacter generates a new character
func NewCharacter() *Character {
	sex := pickSex()
	Character := &Character{
		Name:          generateName(sex),
		Race:          pickRace(),
		Age:           generateAge(),
		Sex:           sex,
		Class:         pickClass(),
		AbilityScores: generateAbilityScores(),
	}
	return Character
}
