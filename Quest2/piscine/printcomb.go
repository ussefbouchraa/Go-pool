package piscine

import (
	"github.com/01-edu/z01"
)


func PrintComb(){
		for i:= 0; i < 10; i++{
			for j := i + 1; j < 10 ; j++{
				for x := j + 1; x < 10; x++{
					z01.PrintRune(rune(i + '0'))
					z01.PrintRune(rune(j + '0'))
					z01.PrintRune(rune(x + '0'))
					
					if i != 7 || j != 8 || x != 9{
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
		z01.PrintRune('\n')
}
