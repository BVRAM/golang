package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("emp:",a)

	var b[3] int
	fmt.Println(b)

	a[1] = 1 
	fmt.Println(a)
	fmt.Println(a[1])

	fmt.Println(len(a))

	c := [10]int{0,1,2,3,4,5}
	fmt.Println(c) 

	var twoD [2][3]int
	for i :=0; i<2; i++{
		for j:=0; j<3; j++{
			twoD[i][j]=i+j
		}
	}

	fmt.Println("2D : ", twoD)
}

// //output
// emp: [0 0 0 0 0]
// [0 0 0]
// [0 1 0 0 0]
// 1
// 5
// [0 1 2 3 4 5 0 0 0 0]
// 2D :  [[0 1 2] [1 2 3]]
