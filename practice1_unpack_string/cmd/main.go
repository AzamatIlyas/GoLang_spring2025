package main

import (
	"fmt"
	"practice1_unpack_string"
)


func main() {
	fmt.Println("Hello World")
	fmt.Println(practice1_unpack_string.Unpack("a4bc2d5e"))
	fmt.Println(practice1_unpack_string.Unpack("abccd"))
	fmt.Println(practice1_unpack_string.Unpack(""))
	fmt.Println(practice1_unpack_string.Unpack("aaa0b"))
	fmt.Println(practice1_unpack_string.Unpack("d\n5abc"))
}
