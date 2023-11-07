package main

import "fmt"

func main(){
	if 7%2 == 0{
		fmt.Println("7 is even")
	} else{
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 17%2 == 0 && 17%4 == 0 {
		fmt.Println("17 is even and divisible by 4")
	}	else{
		fmt.Println("17 is odd")
	}

	if 8%2 == 0 || 8%5 == 0 {
		fmt.Println("8 is even")
	}	else{
		fmt.Println("17 is odd")
	}

	if num := 9; num < 0 {
			fmt.Println(num, "is negative")
	} else if num < 10 {
			fmt.Println(num, "has 1 digit")
	} else {
			fmt.Println(num, "has multiple digits")
	}
}

// //output
// 7 is odd
// 8 is divisible by 4
// 17 is odd
// 8 is even
// 9 has 1 digit
