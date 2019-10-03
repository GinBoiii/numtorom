package main

import (
	"testing"
)

type RomNumsTest struct {
	num    int
	rom    string
	oopsie bool
}

var romNumsLabRat = []RomNumsTest{
	{-1, "", true},
	{0, "", true},
	{1, "I", false},
	{2, "II", false},
	{3, "III", false},
	{4, "IV", false},
	{5, "V", false},
	{6, "VI", false},
	{9, "IX", false},
	{27, "XXVII", false},
	{48, "XLVIII", false},
	{49, "XLIX", false},
	{59, "LIX", false},
	{93, "XCIII", false},
	{141, "CXLI", false},
	{163, "CLXIII", false},
	{402, "CDII", false},
	{575, "DLXXV", false},
	{911, "CMXI", false},
	{1024, "MXXIV", false},
	{3000, "MMM", false},
	{4000, "", true},
}

func testArabToRomanConv(t *testing.T) {
	rats := romNumsLabRat

	for _, test1 := range rats {

		//Tests goToRome function
		current, error := goToRome(test1.num)

		if error != nil && !test1.oopsie {
			t.Errorf("Oh! There should have been an error regarding the goToRome(%d) function call.", test1.num)
		}
		if error == nil && test1.oopsie {
			t.Errorf("Oops! There should't have been an error regarding the goToRome(%d) function call.", test1.num)
		}
		if test1.oopsie && current != test1.rom {
			t.Errorf("This ain't it chief! goToRome(%d):%s isn't the same as %s.", test1.num, current, test1.rom)
		}

		//Test just the hiVal function
		for _, test2 := range rats {
			n := hiVal(test2.num)
			for _, r := range maxi {
				if test2.num < 5 && r != 1 {
					t.Errorf("%d should return 1, but %d", n, r)
				}
				if test2.num < 10 && r != 5 {
					t.Errorf("%d should return 5, but %d", n, r)
				}
				if test2.num < 50 && r != 10 {
					t.Errorf("%d should return 10, but %d", n, r)
				}
				if test2.num < 100 && r != 50 {
					t.Errorf("%d should return 50, but %d", n, r)
				}
				if test2.num < 500 && r != 100 {
					t.Errorf("%d should return 100, but %d", n, r)
				}
				if test2.num < 1000 && r != 500 {
					t.Errorf("%d should return 500, but %d", n, r)
				}
				if test2.num <= 3000 && r != 1000 {
					t.Errorf("%d should return 1000, but %d", n, r)
				}
			}
		}
	}
}

func benchmarkGoToRome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range romNumsLabRat {
			goToRome(test.num)
		}
	}
}
