package main

import (
	"fmt"
	"log"
)

// basic types - numbers, strings, booleans
var myInt int
var myInt16 int16
var myInt32 int32
var myInt64 int64
var myFloat32 float32
var myFloat64 float64
var myString string
var myBoo bool

func main() {
	fmt.Println("Types..Expressions..Composition..Interfaces..")
	basicTypes()

}

// basic types - numbers, strings, booleans
func basicTypes() {
	fmt.Println(myInt, myInt16, myInt32, myInt64, myFloat32, myFloat64)
	log.Println(myInt, myInt16, myInt32, myInt64, myFloat32, myFloat64)
	log.Println(myString)
	myString = "venkat"
	log.Println(myString)
	myString = "venky"
	log.Println(myString)
	log.Println(myBoo)
	var myBool bool
	log.Println(myBool)
	myBool = true
	log.Println(myBool)
}
