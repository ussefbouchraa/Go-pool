package main

import (
	"fmt"
	"strings"
)

func Split(s[] string, sep string) []string{


var res []string
word := ""

for i:= 0 ; i < len(s) ;i++{
		if  strings.Contains(s[i: ], sep) && s[i : i + len(sep)] == sep{
				if word != "" {
						res = append(res, word)
						word = ""
				}
		}else{
			word += s[i]
		}
	}
	if word != "" {
			res = append(res, word)
	}

	return res
}

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}