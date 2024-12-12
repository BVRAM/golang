package main

import (
	"fmt"
	"time"
)

func main(){
	i:=3
	fmt.Println("write",i,"as")

	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	var day=time.Now().Weekday()
	fmt.Println(day)

	switch day {
	case time.Saturday, time.Sunday:
		fmt.Println("It's a weekend")
	default:
		fmt.Println("It's a weekday.. get to work fast")
	}
	
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Its before noon")
	default:
		fmt.Println("Its after noon")
	}

	whatAmI := func (i interface{})  {
		switch t:= i.(type){
		case bool:
				fmt.Println("I'm a bool")
		case int:
				fmt.Println("I'm an int")
		default:
				fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(1)
	whatAmI(true)
	whatAmI(1.234)
	whatAmI("Goo")
}
//output
// write 3 as
// three
// Wednesday
// It's a weekday.. get to work fast
// Its after noon
// I'm an int
// I'm a bool
// Don't know type float64
// Don't know type string
