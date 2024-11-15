package main

import (
	"os"
	"path/filepath"

	"github.com/01-edu/z01"
)

func main() {
	arg := filepath.Base(os.Args[0])
	for _, val := range arg {
		z01.PrintRune(val)
	}
}
