package fun15

import (
	"fmt"
	"testing"
)

func TestOldRomanNumerals(t *testing.T) {
	got := ConvertToRoman(1)
	want := "I"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

/*
Not much to refactor yet.

I know it feels weird just to hard-code the result but with TDD we want to stay out of
"red" for as long as possible. It may feel like we haven't accomplished much but we've
defined our API and got a test capturing one of our rules; even if the "real" code is
pretty dumb.

Now use that uneasy feeling to write a new test to force us to write slightly less dumb code.
*/

func TestOld2RomanNumerals(t *testing.T) {
	t.Run("1 gets converted to I", func(t *testing.T) {
		got := ConvertToRoman(1)
		want := "I"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("2 get converted to II", func(t *testing.T) {
		got := ConvertToRoman(2)
		want := "II"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

/*
Let's take advantage of an slice of struct to define more cases to test
since we have some repetition in our testcase
*/
var cases = []struct {
	Description  string
	NumberArabic int
	NumberRoman  string
}{
	{"1 gets converted to I", 1, "I"},
	{"2 gets converted to II", 2, "II"},
	{"3 gets converted to III", 3, "III"},
	{"4 gets converted to IV (can't repeat more than 3 times)", 4, "IV"},
	{"9 gets converted to IX", 9, "IX"},
	{"10 gets converted to X", 10, "X"},
	{"14 gets converted to XIV", 14, "XIV"},
	{"18 gets converted to XVIII", 18, "XVIII"},
	{"20 gets converted to XX", 20, "XX"},
	{"39 gets converted to XXXIX", 39, "XXXIX"},
	{"40 gets converted to XL", 40, "XL"},
	{"47 gets converted to XLVII", 47, "XLVII"},
	{"49 gets converted to XLIX", 49, "XLIX"},
	{"50 gets converted to L", 50, "L"},
	{"100 gets converted to C", 100, "C"},
	{"90 gets converted to XC", 90, "XC"},
	{"400 gets converted to CD", 400, "CD"},
	{"500 gets converted to D", 500, "D"},
	{"900 gets converted to CM", 900, "CM"},
	{"1000 gets converted to M", 1000, "M"},
	{"1984 gets converted to MCMLXXXIV", 1984, "MCMLXXXIV"},
	{"3999 gets converted to MMMCMXCIX", 3999, "MMMCMXCIX"},
	{"2014 gets converted to MMXIV", 2014, "MMXIV"},
	{"1006 gets converted to MVI", 1006, "MVI"},
	{"798 gets converted to DCCXCVIII", 798, "DCCXCVIII"},
}

func TestConvertToRoman(t *testing.T) {
	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.NumberArabic)
			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}
}
