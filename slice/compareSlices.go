package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Comparing slices.. ")

	/*
		Go doesn't directly support comparing slices for equality using the == operator.
		This is because slices are reference types, and comparing them would check if they point to the same underlying array,
		not if they contain the same elements.

		Key Points to Remember:
		Element-wise Comparison >>>> Using the reflect Package
		Performance Considerations: For large slices, element-wise comparison is generally more efficient than using reflect.DeepEqual.
	*/

	// Element-wise Comparison:
	// This is the most common approach for comparing slices.

	sl1 := []int{1, 2, 3, 4}
	sl2 := []int{1, 2, 3, 4}
	sl3 := []string{"Apple", "Bat", "Cat"}
	sl4 := []string{"Apple", "Bat", "Cat"}

	fmt.Println(equalSlices(sl1, sl2))  //true
	fmt.Println(equalSlices1(sl3, sl4)) //true

	// Using the reflect Package:
	// This function can be used for more complex data structures, but it's less efficient for simple slices.
	if reflect.DeepEqual(sl1, sl2) {
		fmt.Println("Slices are equal")
	} else {
		fmt.Println("Slices are not equal")
	}
}

func equalSlices(slice1, slice2 []int) bool {

	if len(slice1) != len(slice2) {
		return false
	}

	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

func equalSlices1(slice1, slice2 []string) bool {

	if len(slice1) != len(slice2) {
		return false
	}

	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}
