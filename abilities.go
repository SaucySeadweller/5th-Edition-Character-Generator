package main

type AbilityScores struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
}
type Abilities struct {
	Feats  []Feat
	Spells []Spell
}

type Skill struct {
	Athletics      int
	Acrobatics     int
	SleightOfHand  int
	Stealth        int
	Arcana         int
	History        int
	Investigation  int
	Nature         int
	Religion       int
	AnimalHandling int
	Insight        int
	Medicine       int
	Perception     int
	Survival       int
	Deception      int
	Intimidation   int
	Performance    int
	Persuasion     int
}
