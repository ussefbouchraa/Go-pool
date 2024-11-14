package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main(){

	// for i := 1; i < len(os.Args); i++{
			for j:=  1 ; j < len(os.Args)  ; j++{
				if( len(os.Args) > j + 1 && (os.Args[j]) > (os.Args[j + 1]) ){
					tmp := os.Args[j]
					os.Args[j] = os.Args[j + 1]
					os.Args[j+1] = tmp

				}
			// }
	}
		for i := 1; i < len(os.Args); i++{
		fmt.Println(filepath.Base(os.Args[i]))
		}
}