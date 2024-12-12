package main

import "fmt"

/*

Slices in Go: A Deeper Dive
Slices are one of the most powerful data structures in Go.
They provide a flexible way to work with sequences of elements, similar to arrays but with dynamic sizing capabilities.

*/
func main() {
	fmt.Println("Its about slice in golang")
	slices()

}
func slices() {

	//Creating a Slice: var sliceName []dataType
	var slice1 []int = []int{1, 2, 3, 4, 5}

	/*
		Initializing a Slice
	*/

	// (1) Direct intialization
	slice2 := []string{"Apple", "Banana", "Choco"}

	// (2) From Array
	myArray := [5]int{100, 200, 300, 400, 500}
	myslice1 := myArray[:5]

	// (3) Slice of certain length
	myslice2 := make([]int, 3)

	/*
		Accessing slice Elements
	*/

	fmt.Println(slice1[0])   //1
	fmt.Println(slice2[1])   //Banana
	fmt.Println(myslice1[1]) //200
	fmt.Println(myslice2[1]) //0

	/*

		Slice Length and Capacity:

		Length: The number of elements currently in the slice.
		Capacity: The total number of elements the underlying array can hold.

	*/
	length := len(slice1)
	capacity := cap(slice1)

	fmt.Println(length)   //5
	fmt.Println(capacity) //3

	/*
		Slicing a Slice:
	*/

	newSlice := slice1[2:4] // Slice from index 2 (inclusive) to 4 (exclusive)
	fmt.Println(slice1)     //[1,2,3,4,5]
	fmt.Println(newSlice)   //[3,4]

	/*
		Appending to Slice
	*/

	newSlice = append(newSlice, 5, 6, 7)
	fmt.Println(newSlice) //[3,4,5,6,7]

	/*
		Deleting Elements from a Slice:
	*/

	newSlice = newSlice[3:] //Removes starting 3 elements
	fmt.Println(newSlice)   //[6,7]

	/*
		Nil Slices:

		A nil slice is a slice with a length and capacity of 0.
	*/

	var nilSlice []int
	fmt.Println(nilSlice)      //[]
	fmt.Println(len(nilSlice)) // 0
	fmt.Println(cap(nilSlice)) // 0

	/*
		Multi-Dimensional Slices:

		Go supports multi-dimensional slices:
	*/

	var matrix [][]int = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)

}
