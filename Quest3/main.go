package main

import (
	"fmt"
	"Root/piscine"
)

// func main() {
// 	n := 0
// 	piscine.PointOne(&n)
// 	fmt.Println(n)
// }


// func main() {
// 	a := 0
// 	b := &a
// 	n := &b
// 	piscine.UltimatePointOne(&n)
// 	fmt.Println(a)
// }



// func main() {
// 	a := 13
// 	b := 2
// 	var div int
// 	var mod int
// 	piscine.DivMod(a, b, &div, &mod)
// 	fmt.Println(div)
// 	fmt.Println(mod)
// }



// func main() {
// 	a := 13
// 	b := 2
// 	piscine.UltimateDivMod(&a, &b)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }


// func main() {
// 	piscine.PrintStr("Hello World!")
// }


// func main() {
// 	l := piscine.StrLen("Hello World!")
// 	fmt.Println(l)
// }



// func main() {
// 	a := 0
// 	b := 1
// 	piscine.Swap(&a, &b)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }



// func main() {
// 	s := "Hello World!"
// 	s = piscine.StrRev(s)
// 	fmt.Println(s)

// }


// func main() {
// 	fmt.Println(piscine.BasicAtoi("12345"))
// 	fmt.Println(piscine.BasicAtoi("0000000012345"))
// 	fmt.Println(piscine.BasicAtoi("000000"))
// }

// func main() {
// 	fmt.Println(piscine.BasicAtoi2("12345"))
// 	fmt.Println(piscine.BasicAtoi2("0000000012345"))
// 	fmt.Println(piscine.BasicAtoi2("012 345"))
// 	fmt.Println(piscine.BasicAtoi2("Hello World!"))
// }



// func main() {
// 	fmt.Println(piscine.Atoi("12345"))
// 	fmt.Println(piscine.Atoi("0000000012345"))
// 	fmt.Println(piscine.Atoi("012 345"))
// 	fmt.Println(piscine.Atoi("Hello World!"))
// 	fmt.Println(piscine.Atoi("+1234"))
// 	fmt.Println(piscine.Atoi("-1234"))
// 	fmt.Println(piscine.Atoi("++1234"))
// 	fmt.Println(piscine.Atoi("--1234"))
// }



func main() {
	s := []int{5,4,3,2,1,0}
	piscine.SortIntegerTable(s)
	fmt.Println(s)
}