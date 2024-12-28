package main

import "fmt"

func main() {

	// The & Operator:

	var x int = 10
	var ptr *int = &x // The & operator is used to get the memory address of a variable.

	// Here, ptr now stores the memory address of x.

	fmt.Println(x)   // 10
	fmt.Println(&x)  // 0x00000078wa
	fmt.Println(ptr) // 0x00000078wa

	// The * Operator:
	fmt.Println(*ptr) //10

	// * operator is used to dereference a pointer, which means accessing the value stored at the memory address.
	*ptr = 20

	fmt.Println(x)    //20
	fmt.Println(*ptr) //20

	a := 10
	b := 20
	fmt.Println(a, b)

	swap(&a, &b)
	fmt.Println(a, b)

	// Array indexing with pointers

	arr := [5]int{10, 20, 30, 40, 50}
	ptr1 := &arr[0]

	fmt.Println(*ptr)      // Output: 10
	fmt.Println(*ptr1 + 1) // Output: 11
	fmt.Println(*ptr1 + 2) // Output: 12
}

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
