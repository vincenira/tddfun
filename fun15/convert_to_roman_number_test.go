package fun15

import "testing"

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
func TestConvertToRoman(t *testing.T) {
	cases := []struct {
		Description  string
		NumberArabic int
		Want         string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
		{"3 gets converted to III", 3, "III"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.NumberArabic)
			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}
}
