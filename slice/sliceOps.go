package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("operations with slices..")

	// 1. Copying Slices:
	// You can create a copy of a slice to avoid unintended modifications to the original slice:

	oldSlice := []int{1, 2, 3}
	newSlice := make([]int, len(oldSlice))

	copy(newSlice, oldSlice)

	fmt.Println(oldSlice) // [1,2,3]
	fmt.Println(newSlice) // [1,2,3]

	// 2. Sorting Slices:
	// <sort> package can be used

	s1 := []int{1, 3, 5, 7, 2, 4, 6, 8, 9}
	fmt.Println(s1)
	sort.Ints(s1)
	fmt.Println(s1)

	s2 := []string{"Apple", "App", "Account", "Architecture", "Banana", "Bat", "Ant", "curd", "cat"}
	fmt.Println(s2)
	sort.Strings(s2)
	fmt.Println(s2)

	// 3. Reversing Slices:
	// reverse the order of elements in a slice:

	for i, j := 0, len(s1)-1; i < j; i, j = i+1, j-1 {
		s1[i], s1[j] = s1[j], s1[i]
	}
	fmt.Println(s1)

	// 4. Finding Elements in Slices:

	for i, value := range s1 {
		if value == 1 {
			fmt.Println("Found at index: ", i)
			break
		}
	}

	// 5. Filtering Slices:

	// create a new slice containing only the elements that meet a certain condition

	filteredSlice := make([]int, 0)
	for _, value := range s1 {
		if value > 5 {
			filteredSlice = append(filteredSlice, value)
		}
	}

	fmt.Println(filteredSlice)

	// 6. Mapping Slices

	// create a new slice by applying a function to each element of an existing slice

	newSlice1 := make([]int, len(s1))
	for i, value := range s1 {
		newSlice1[i] = value * 2
	}
	fmt.Println(newSlice1)
}
