package main

import "fmt"

func main(){
	i := 1
	
	for i<=4{
		fmt.Println(i)
		i=i+1
	}	

	for j :=5; j<=9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n:=1;n<=10;n++{
		if n%2==0 {
			continue
		}
		fmt.Println(n)
	}
}
//output
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// loop
// 1
// 3
// 5
// 7
// 9
