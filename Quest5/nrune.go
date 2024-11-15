package piscine

func NRune(s string, n int) rune {
	if len(s) == 0 || n <= 0 || n > len(s){
		return 0
	}
	runes := []rune(s)
	return runes[n -1]
}