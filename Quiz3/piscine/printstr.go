package piscine

import "fmt"

func PrintStr(s string) {


	for i:=0 ;i < len(s); i++{
		fmt.Printf("%c", s[i])
	}
}