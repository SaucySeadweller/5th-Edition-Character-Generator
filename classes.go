package main

type Class struct {
	ClassName     string
	HitDie        []int
	HitPoints     int
	Proficiencies []string
	ClassFeatures []string
	Equipment     []string
	Spells        []Spell
	SpellSlots    int
	Initiative    float64
}
