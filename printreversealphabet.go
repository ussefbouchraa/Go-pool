package main

import "github.com/01-edu/z01"

func main() {

	for alpha := 'z'; alpha >= 'a' ; alpha-- {
		z01.PrintRune(rune(alpha));
	}
	z01.PrintRune('\n')
}
