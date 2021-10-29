package classmanager

type Class struct {
	Name                                     string
	HasSpell, HasSpellWithSubClass, HasFeats bool
	Feats                                    Feats
	FeatsAmountByLevel                       map[int]int
	SubClasses                               []SubClass
	SpellLevelLimit                          int
	SpellByLevel                             map[int]map[int]int
}

type SubClass struct {
	Name               string
	HasSpell, HasFeats bool
	SpellLevelLimit    int
}

type Abilities string

const (
	None         Abilities = "none"
	Constitution Abilities = "cons"
	Charism      Abilities = "char"
	Strengh      Abilities = "str"
	Agility      Abilities = "agi"
	Intelligence Abilities = "int"
	Wisdom       Abilities = "wis"
	Mastery      Abilities = "mas"
)

type Feats struct {
	Name         string
	CharModifier Abilities
}

var (
	NoFeats           = Feats{"", None}
	BardicInspiration = Feats{"Bardic Inspiration", Charism}
	HemocraftDice     = Feats{"EmocraftDice", None}
	Infusion          = Feats{"Infusion", None}
	Ki                = Feats{"Ki", None}
	PsionicEnergyDice = Feats{"Psionic Energy Dice", Mastery}
	Rage              = Feats{"Rage", None}
	SorceryPoint      = Feats{"Sorcery Point", None}
)

var ClassList []Class = []Class{
	{
		Name:                 "Artificer",
		HasSpell:             true,
		HasSpellWithSubClass: false,
		HasFeats:             true,
		Feats:                Infusion,
		SubClasses: []SubClass{
			{"Alchemist", false, false, 5},
			{"Armorer", false, false, 5},
			{"Artillerist", false, false, 5},
			{"Battle Smith", false, false, 5},
		},
		SpellLevelLimit: 5,
		FeatsAmountByLevel: map[int]int{
			1:  0,
			2:  2,
			3:  2,
			4:  2,
			5:  2,
			6:  3,
			7:  3,
			8:  3,
			9:  3,
			10: 4,
			11: 4,
			12: 4,
			13: 4,
			14: 5,
			15: 5,
			16: 5,
			17: 5,
			18: 6,
			19: 6,
			20: 6,
		},
	},
	{
		Name:                 "Barbarian",
		HasSpell:             false,
		HasSpellWithSubClass: false,
		HasFeats:             true,
		Feats:                Rage,
		SubClasses: []SubClass{
			{"Path of the Ancestral Guardian", false, false, 0},
			{"Path of the Battlerager", false, false, 0},
			{"Path of the Beast", false, false, 0},
			{"Path of the Berserker", false, false, 0},
			{"Path of the Storm Herald", false, false, 0},
			{"Path of the Totem Warrior", false, false, 0},
			{"Path of Wild Magic", false, false, 0},
			{"Path of the Zealot", false, false, 0},
		},
		FeatsAmountByLevel: map[int]int{},
	},
	{
		Name:                 "Bard",
		HasSpell:             true,
		HasSpellWithSubClass: false,
		HasFeats:             true,
		Feats:                BardicInspiration,
		SubClasses: []SubClass{
			{"College of Creation", false, false, 9},
			{"College of Eloquence", false, false, 9},
			{"College of Glamour", false, false, 9},
			{"College of Lore", false, false, 9},
			{"College of Spirits", false, false, 9},
			{"College of Swords", false, false, 9},
			{"College of Valor", false, false, 9},
			{"College of Whispers", false, false, 9},
			{"Mage of Silverquill", false, false, 9},
		},
		SpellLevelLimit: 9,
		SpellByLevel: map[int]map[int]int{
			1:  {1: 2},
			2:  {1: 3},
			3:  {1: 4, 2: 2},
			4:  {1: 4, 2: 3},
			5:  {1: 4, 2: 3, 3: 2},
			6:  {1: 4, 2: 3, 3: 3},
			7:  {1: 4, 2: 3, 3: 3, 4: 1},
			8:  {1: 4, 2: 3, 3: 3, 4: 2},
			9:  {1: 4, 2: 3, 3: 3, 4: 3, 5: 1},
			10: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2},
			11: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1},
			12: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1},
			13: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1},
			14: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1},
			15: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1, 8: 1},
			16: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1, 8: 1},
			17: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1, 8: 1, 9: 1},
			18: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1, 8: 1, 9: 1},
			19: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 2, 7: 1, 8: 1, 9: 1},
			20: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 2, 7: 2, 8: 1, 9: 1},
		},
	},
	{
		Name:                 "Blood Hunter",
		HasSpell:             true,
		HasSpellWithSubClass: false,
		HasFeats:             false,
		SubClasses: []SubClass{
			{"Order of the Ghostslayer", false, false, 0},
			{"Order of the Lycan", false, false, 0},
			{"Order of the Mutant", false, false, 0},
			{"Order of the Profane Soul", false, false, 0},
		},
	},
	{
		Name:                 "Cleric",
		HasSpell:             true,
		HasSpellWithSubClass: false,
		HasFeats:             false,
		SubClasses: []SubClass{
			{"Arcana Domain", false, false, 9},
			{"Death Domain", false, false, 9},
			{"Forge Domain", false, false, 9},
			{"Grave Domain", false, false, 9},
			{"Knowledge Domain", false, false, 9},
			{"Life Domain", false, false, 9},
			{"Light Domain", false, false, 9},
			{"Nature Domain", false, false, 9},
			{"Order Domain", false, false, 9},
			{"Peace Domain", false, false, 9},
			{"Tempest Domain", false, false, 9},
			{"Trickery Domain", false, false, 9},
			{"Twilight Domain", false, false, 9},
			{"War Domain", false, false, 9},
			{"Ambition Domain", false, false, 9},
			{"Solidarity Domain", false, false, 9},
			{"Strength Domain", false, false, 9},
			{"Zeal Domain", false, false, 9},
		},
		SpellLevelLimit: 9,
		SpellByLevel: map[int]map[int]int{
			1:  map[int]int{1: 2},
			2:  map[int]int{1: 3},
			3:  map[int]int{1: 4, 2: 2},
			4:  map[int]int{1: 4, 2: 3},
			5:  map[int]int{1: 4, 2: 3, 3: 2},
			6:  map[int]int{1: 4, 2: 3, 3: 3},
			7:  map[int]int{1: 4, 2: 3, 3: 3, 4: 1},
			8:  map[int]int{1: 4, 2: 3, 3: 3, 4: 2},
			9:  map[int]int{1: 4, 2: 3, 3: 3, 4: 3, 5: 1},
			10: map[int]int{1: 4, 2: 3, 3: 3, 4: 3, 5: 2},
			11: map[int]int{1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1},
			12: map[int]int{1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1},
			13: map[int]int{1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1},
			14: map[int]int{1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1},
			15: map[int]int{1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1, 8: 1},
			16: map[int]int{1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1, 8: 1},
			17: map[int]int{1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1, 8: 1, 9: 1},
			18: map[int]int{1: 4, 2: 3, 3: 3, 4: 3, 5: 3, 6: 1, 7: 1, 8: 1, 9: 1},
			19: map[int]int{1: 4, 2: 3, 3: 3, 4: 3, 5: 3, 6: 2, 7: 1, 8: 1, 9: 1},
			20: map[int]int{1: 4, 2: 3, 3: 3, 4: 3, 5: 3, 6: 2, 7: 2, 8: 1, 9: 1},
		},
	},
	{
		Name:                 "Druid",
		HasSpell:             true,
		HasSpellWithSubClass: false,
		HasFeats:             false,
		SubClasses: []SubClass{
			{"Circle of Dreams", false, false, 9},
			{"Circle of the Land", false, false, 9},
			{"Circle of the Moon", false, false, 9},
			{"Circle of the Shepherd", false, false, 9},
			{"Circle of Spores", false, false, 9},
			{"Circle of Stars", false, false, 9},
			{"Circle of Wildfire", false, false, 9},
			{"Mage of Prismari", false, false, 9},
			{"Mage of Witherbloom", false, false, 9},
		},
		SpellLevelLimit: 9,
		SpellByLevel: map[int]map[int]int{
			1:  {1: 2},
			2:  {1: 3},
			3:  {1: 4, 2: 2},
			4:  {1: 4, 2: 3},
			5:  {1: 4, 2: 3, 3: 2},
			6:  {1: 4, 2: 3, 3: 3},
			7:  {1: 4, 2: 3, 3: 3, 4: 1},
			8:  {1: 4, 2: 3, 3: 3, 4: 2},
			9:  {1: 4, 2: 3, 3: 3, 4: 3, 5: 1},
			10: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2},
			11: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1},
			12: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1},
			13: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1},
			14: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1},
			15: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1, 8: 1},
			16: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1, 8: 1},
			17: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1, 8: 1, 9: 1},
			18: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1, 8: 1, 9: 1},
			19: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 2, 7: 1, 8: 1, 9: 1},
			20: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 2, 7: 2, 8: 1, 9: 1},
		},
	},
	{
		Name:                 "Fighter",
		HasSpell:             false,
		HasSpellWithSubClass: true,
		HasFeats:             false,
		Feats:                NoFeats,
		SubClasses: []SubClass{
			{"Arcane Archer", false, false, 0},
			{"Banneret", false, false, 0},
			{"Battle Master", false, false, 0},
			{"Cavalier", false, false, 0},
			{"Champion", false, false, 0},
			{"Echo Knight", false, false, 0},
			{"Eldritch Knight", true, false, 4},
			{"Psi Warrior", false, true, 0},
			{"Rune Knight", false, false, 0},
			{"Samurai", false, false, 0},
		},
	},
	{
		Name:                 "Monk",
		HasSpell:             false,
		HasSpellWithSubClass: false,
		HasFeats:             true,
		Feats:                Ki,
		SubClasses: []SubClass{
			{"Way of Mercy", false, false, 0},
			{"Way of the Astral Self", false, false, 0},
			{"Way of the Drunken Master", false, false, 0},
			{"Way of the Four Elements", false, false, 0},
			{"Way of the Kensei", false, false, 0},
			{"Way of the Long Death", false, false, 0},
			{"Way of the Open Hand", false, false, 0},
			{"Way of Shadow", false, false, 0},
			{"Way of the Sun Soul", false, false, 0},
			{"Way of the Ascendant Dragon", false, false, 0},
		},
	},
	{
		Name:                 "Paladin",
		HasSpell:             true,
		HasSpellWithSubClass: false,
		HasFeats:             false,
		SubClasses: []SubClass{
			{"Oath of the Ancients", false, false, 5},
			{"Oath of Conquest", false, false, 5},
			{"Oath of the Crown", false, false, 5},
			{"Oath of Devotion", false, false, 5},
			{"Oath of Glory", false, false, 5},
			{"Oath of Redemption", false, false, 5},
			{"Oath of Vengeance", false, false, 5},
			{"Oath of the Watchers", false, false, 5},
			{"Oathbreaker", false, false, 5},
		},
		SpellLevelLimit: 5,
	},
	{
		Name:                 "Ranger",
		HasSpell:             true,
		HasSpellWithSubClass: false,
		HasFeats:             false,
		SubClasses: []SubClass{
			{"Beast Master Conclave", false, false, 5},
			{"Fey Wanderer", false, false, 5},
			{"Gloom Stalker Conclave", false, false, 5},
			{"Horizon Walker Conclave", false, false, 5},
			{"Hunter Conclave", false, false, 5},
			{"Monster Slayer Conclave", false, false, 5},
			{"Swarmkeeper", false, false, 5},
			{"Drakewarden", false, false, 5},
		},
		SpellLevelLimit: 5,
	},
	{
		Name:                 "Rogue",
		HasSpell:             false,
		HasSpellWithSubClass: true,
		HasFeats:             false,
		SubClasses: []SubClass{
			{"Arcane Trickster", false, false, 4},
			{"Assassin", false, false, 0},
			{"Inquisitive", false, false, 0},
			{"Mastermind", false, false, 0},
			{"Phantom", false, false, 0},
			{"Scout", false, false, 0},
			{"Soulknife", false, true, 0},
			{"Swashbuckler", false, false, 0},
			{"Thief", false, false, 0},
		},
	},
	{
		Name:                 "Sorcerer",
		HasSpell:             true,
		HasSpellWithSubClass: false,
		HasFeats:             true,
		Feats:                SorceryPoint,
		SubClasses: []SubClass{
			{"Aberrant Mind", false, false, 9},
			{"Clockwork Soul", false, false, 9},
			{"Draconic Bloodline", false, false, 9},
			{"Divine Soul", false, false, 9},
			{"Shadow Magic", false, false, 9},
			{"Storm Sorcery", false, false, 9},
			{"Wild Magic", false, false, 9},
			{"Pyromancy", false, false, 9},
			{"Mage of Prismari", false, false, 9},
			{"Mage of Quandrix", false, false, 9},
		},
		SpellLevelLimit: 9,
	},
	{
		Name:                 "Warlock",
		HasSpell:             true,
		HasSpellWithSubClass: false,
		HasFeats:             false,
		SubClasses: []SubClass{
			{"Archfey", false, false, 5},
			{"Celestial", false, false, 5},
			{"Fathomless", false, false, 5},
			{"Fiend", false, false, 5},
			{"The Genie", false, false, 5},
			{"Great Old One", false, false, 5},
			{"Hexblade", false, false, 5},
			{"Undead", false, false, 5},
			{"Undying", false, false, 5},
			{"Mage of Silverquill", false, false, 5},
			{"Mage of Witherbloom", false, false, 5},
		},
		SpellLevelLimit: 5,
	},
	{
		Name:                 "Wizard",
		HasSpell:             true,
		HasSpellWithSubClass: false,
		HasFeats:             false,
		SubClasses: []SubClass{
			{"School of Abjuration", false, false, 9},
			{"School of Bladesinging", false, false, 9},
			{"School of Chronurgy", false, false, 9},
			{"School of Conjuration", false, false, 9},
			{"School of Divination", false, false, 9},
			{"School of Enchantment", false, false, 9},
			{"School of Evocation", false, false, 9},
			{"School of Graviturgy", false, false, 9},
			{"School of Illusion", false, false, 9},
			{"School of Necromancy", false, false, 9},
			{"Order of Scribes", false, false, 9},
			{"School of Transmutation", false, false, 9},
			{"School of War Magic", false, false, 9},
			{"Mage of Lorehold", false, false, 9},
			{"Mage of Prismari", false, false, 9},
			{"Mage of Quandrix", false, false, 9},
			{"Mage of Silverquill", false, false, 9},
		},
		SpellLevelLimit: 9,
	},
}

func FindClassByName(name string) Class {
	var class Class
	for _, v := range ClassList {
		if v.Name == name {
			class = v
		}
	}
	return class
}

func FindSubClassByName(name string, className string) SubClass {
	var sub SubClass

	for _, v := range FindClassByName(className).SubClasses {
		if v.Name == name {
			sub = v
		}
	}
	return sub
}
