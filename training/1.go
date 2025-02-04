package main

import (
	"fmt"
	// "unicode"
)

func main() {
	s := "aaa\4"
	runes := []rune(s)
	for i:=0; i < len(s); i++ {
		fmt.Println(runes[i])
	}
}
