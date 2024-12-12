package main

import (
	"fmt"
)

//variables in golang

func main() {
	var a string
	a="hello"
	fmt.Println(a)

	var b string = "world"
	fmt.Println(b)

	var c,d int = 1,2
	fmt.Println(c,d)

	var e = true
	fmt.Println(e)

	var f int
	fmt.Println(f)

	g:="Goa"
	fmt.Println(g)

}

//output:
// hello
// world
// 1 2
// true
// 0
// Goa
