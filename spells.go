package main

type Spell struct {
	SpellName     string
	Effect        string
	Damage        int
	Range         int
	SpellLevel    int
	Components    Component
	CastTime      int
	action        string
	School        string
	Duration      int
	Concentration bool
	Ritual        bool
	AreaOfEffect  string
	Target        string
}
type School struct {
	Abjuration    string
	Conjuration   string
	Divination    string
	Destruction   string
	Enchantment   string
	Evocation     string
	Illusion      string
	Necromancy    string
	Transmutation string
}
type Component struct {
	Verbal   string
	Somatic  string
	Material string
}
