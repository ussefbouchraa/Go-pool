package piscine

import (
	"github.com/01-edu/z01"
)

func SplitWhiteSpaces(s string) []string {
	var res []string
	word := ""

	for _, val := range s {
		if val == ' ' || val >= 9 && val <= 13 {
			if word != "" {
				res = append(res, word)
				word = ""
			}
		} else {
			word += string(val)
		}
	}
	if word != "" {
		res = append(res, word)
	}

	return res
}

func PrintWordsTables(a []string) {
	for i := 0; i < len(a); i++ {
		for _, char := range a[i] {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')

	}
}
