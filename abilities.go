package main

type abilityscores struct {
	strength     []int
	dexterity    []int
	constitution []int
	intelligence []int
	wisdom       []int
	charisma     []int
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
