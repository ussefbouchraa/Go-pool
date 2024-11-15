package main

import (
	"fmt"
)

func main() {
	emoji := "👨‍👩‍👦‍👦"
	fmt.Println("Emoji:", emoji)
	fmt.Println("Length in runes:", len([]rune(emoji)))
	fmt.Println("Length in bytes:", len(emoji))
}
