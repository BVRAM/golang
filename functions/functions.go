package main

import (
	"errors"
	"fmt"
)

func main() {
	wish := greet("Venkat")
	fmt.Println(wish)

	value, _ := divide(10, 2)
	fmt.Println(value)

	value1, error := divide(2, 0)
	fmt.Println(value1)
	fmt.Println(error)

	total := sum(1, 2, 2, 3, 3, 3, 4, 4, 4)
	fmt.Println(total)

	func() {
		fmt.Println("Annonmuous function..!!")
	}()

}

func greet(name string) string {
	return "Hello, " + name + "!"
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet1() {
	fmt.Println("Hello, my name is", p.Name)
}
