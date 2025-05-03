package fun15

import "strings"

/*
It looks that we are not actually tackling the problem, so we need to write more tests
to drive us forward
check TestOld2RomanNumerals

	func ConvertToRoman(arabicNumber int) string {
		if arabicNumber == 2 {
			return "II"
		}
		return "I"
	}
*/
type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabicNumber int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabicNumber >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabicNumber -= numeral.Value
		}
	}
	return result.String()
}

func ConvertToArabic(romanNumber string) int {
	arabicNumber := 0

	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(romanNumber, numeral.Symbol) {
			arabicNumber += numeral.Value
			romanNumber = strings.TrimPrefix(romanNumber, numeral.Symbol)
		}
	}
	return arabicNumber
}
