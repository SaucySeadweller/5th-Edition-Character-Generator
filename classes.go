package main

type Class struct {
	R  Rogue
	F  Fighter
	W  Wizard
	WL Warlock
	B  Bard
	P  Paladin
	C  Cleric
	BB Barbarian
	D  Druid
	M  Monk
	S  Sorcerer
	RA Ranger
}

type Rogue struct {
	//d8
	HitDie        []int
	HitPoints     int
	Proficiencies []string
	ClassFeatures []string
	Equipment     []string
	Spells        []Spell
	SpellSlots    int
	Initiative    float64
	level         int
	ClassSkills   []string
}
type Wizard struct {
	//d6
	HitDie        []int
	HitPoints     int
	Proficiencies []string
	ClassFeatures []string
	level         int
	Equipment     []string
	Spells        []Spell
	SpellSlots    int
	Initiative    float64
	ClassSkills   []string
}
type Warlock struct {
	//d8
	HitDie        []int
	HitPoints     int
	Proficiencies []string
	ClassFeatures []string
	Equipment     []string
	Spells        []Spell
	SpellSlots    int
	level         int
	Initiative    float64
	ClassSkills   []string
}
type Paladin struct {
	//d10
	HitDie        []int
	HitPoints     int
	Proficiencies []string
	ClassFeatures []string
	Equipment     []string
	Spells        []Spell
	SpellSlots    int
	level         int
	Initiative    float64
	ClassSkills   []string
}
type Fighter struct {
	//d10
	HitDie        []int
	HitPoints     int
	level         int
	Proficiencies []string
	ClassFeatures []string
	Equipment     []string
	Spells        []Spell
	SpellSlots    int
	Initiative    float64
	ClassSkills   []string
}
type Bard struct {
	//d8
	HitDie        []int
	HitPoints     int
	level         int
	Proficiencies []string
	ClassFeatures []string
	Equipment     []string
	Spells        []Spell
	SpellSlots    int
	Initiative    float64
	ClassSkills   []string
}
type Druid struct {
	// d8
	HitDie        []int
	HitPoints     int
	Proficiencies []string
	ClassFeatures []string
	Equipment     []string
	Spells        []Spell
	SpellSlots    int
	level         int
	Initiative    float64
	ClassSkills   []string
}
type Ranger struct {
	//d10
	HitDie        []int
	HitPoints     int
	Proficiencies []string
	ClassFeatures []string
	Equipment     []string
	Spells        []Spell
	SpellSlots    int
	level         int
	Initiative    float64
	ClassSkills   []string
}
type Monk struct {
	//d8
	HitDie        []int
	HitPoints     int
	Proficiencies []string
	ClassFeatures []string
	Equipment     []string
	Spells        []Spell
	SpellSlots    int
	level         int
	Initiative    float64
	ClassSkills   []string
}
type Cleric struct {
	//d8
	HitDie        []int
	HitPoints     int
	Proficiencies []string
	ClassFeatures []string
	Equipment     []string
	Spells        []Spell
	SpellSlots    int
	level         int
	Initiative    float64
	ClassSkills   []string
}
type Sorcerer struct {
	//d6
	HitDie        []int
	HitPoints     int
	Proficiencies []string
	ClassFeatures []string
	Equipment     []string
	Spells        []Spell
	SpellSlots    int
	level         int
	Initiative    float64
	ClassSkills   []string
}
type Barbarian struct {
	//d12
	HitDie        []int
	HitPoints     int
	Proficiencies []string
	ClassFeatures []string
	Equipment     []string
	Spells        []Spell
	SpellSlots    int
	level         int
	Initiative    float64
	ClassSkills   []string
}
type Proficiencies struct {
	ArmorProficiency       string
	WeaponProficiency      string
	SavingThrowProficiency string
	ProficiencyBonus       int
}

type Currency struct {
	CopperPieces   int
	SilverPieces   int
	ElectrumPieces int
	GoldPieces     int
	PlatinumPieces int
}
type SavingThrows struct {
	StrengthSave     int
	DexteritySave    int
	ConstitutionSave int
	IntelligenceSave int
	WisdomSave       int
	CharismaSave     int
}
