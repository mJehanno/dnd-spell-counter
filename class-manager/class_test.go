package classmanager

import (
	"fmt"
	"testing"
)

type classCase struct {
	arg      string
	expected Class
}

func TestFindClassByName(t *testing.T) {
	testCases := []classCase{
		{"Artificer", Class{
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
			SpellByLevel: map[int]map[int]int{
				1:  {1: 2},
				2:  {1: 2},
				3:  {1: 3},
				4:  {1: 3},
				5:  {1: 4, 2: 2},
				6:  {1: 4, 2: 3},
				7:  {1: 4, 2: 3},
				8:  {1: 4, 2: 3, 3: 2},
				9:  {1: 4, 2: 3, 3: 2},
				10: {1: 4, 2: 3, 3: 3},
				11: {1: 4, 2: 3, 3: 3},
				12: {1: 4, 2: 3, 3: 3},
				13: {1: 4, 2: 3, 3: 3, 4: 1},
				14: {1: 4, 2: 3, 3: 3, 4: 1},
				15: {1: 4, 2: 3, 3: 3, 4: 2},
				16: {1: 4, 2: 3, 3: 3, 4: 2},
				17: {1: 4, 2: 3, 3: 3, 4: 3, 5: 1},
				18: {1: 4, 2: 3, 3: 3, 4: 3, 5: 1},
				19: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2},
				20: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2},
			},
		}},
		{"Bard", Class{
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
		}},
		{"Wizard", Class{
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
		}},
	}

	for _, v := range testCases {
		r := FindClassByName(v.arg)
		if v.expected.String() != r.String() {
			t.Log(fmt.Sprintf("error in test ! expected %s , got %s", v.expected, r))
			t.Fail()
		}
	}

}

func TestFindSubClassByName(t *testing.T) {

}
