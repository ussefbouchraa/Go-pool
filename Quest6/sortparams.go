package main
import (
	"os"
	"github.com/01-edu/z01"
)
func printArg(arg string) {
	for _, val := range arg {
		if val == '.' || val == '/' {
			continue
		}
		z01.PrintRune(val)
	}
	z01.PrintRune('\n')
}
func main() {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if args[i] > args[j] {
				tmp := args[i]
				args[i] = args[j]
				args[j] = tmp
			}
		}
	}
	for i := 0; i < len(args); i++ {
		printArg(args[i])
	}
}
