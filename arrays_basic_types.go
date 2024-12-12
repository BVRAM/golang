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
	aggregateTypes()

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

//aggregate types - Arrays and structs

/*

What is an Array?
	An array in Go is a collection of elements of the same data type, stored in contiguous memory locations.
	It's a fixed-size data structure, meaning its size is determined at compile time.


Declaring and Initializing Arrays:
	var arrayName [size]dataType

	arrayName: The name of the array.
	size: The number of elements in the array.
	dataType: The data type of the elements.

Initializing Arrays:

	Explicitly: var arrayName = [size]dataType{element1, element2, element3}
	Implicitly: arrayName := [size]dataType{element1, element2, element3}

Accessing Array Elements:

	access individual elements using their index, which starts from 0:

	fmt.Println(arrayName[0])

Iterating over Arrays:

	use a for loop to iterate over each element of an array

	for i:=0;i<len(arrayName);i++ {
		fmt.Println(arrayName[i])
	}

Passing Arrays to Functions:

	pass arrays to functions as arguments

	arrays are passed by value - any modifications made inside the function won't affect the original array

	func printArray(arr [5]int) {
		for _, num := range arr {
			fmt.Println(num)
    }
}

Limitations of Arrays:

	Fixed size: Once declared, the size of an array cannot be changed.
	Inefficient for large datasets: Arrays can be inefficient for large datasets due to memory allocation and copying.
*/

func aggregateTypes() {

	//Explicit way of declaring arrays
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers[4])
	//Implicit wat of declaring arrays
	strings := [5]string{"Arjuna", "Bhima", "Chanakya", "Damodhar", "Eswar"}
	fmt.Println(strings[0])

	//for loop iteration for numbers and strings arrays

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}

	printArray(numbers)
	printStringArray(strings)
}

func printArray(arr [5]int) {
	for _, num := range arr {
		fmt.Println(num)
	}
}

func printStringArray(arr [5]string) {
	for _, num := range arr {
		fmt.Println(num)
	}
}
