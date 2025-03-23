package fun3

func Repeat(letter string) string {
	retLetter := ""
	for range 7 {
		retLetter += letter
	}
	return retLetter
}
