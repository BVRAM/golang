package main

import (
	"fmt"
	"slices"
)

func main() {

	var s []string
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println("uunit", s,s==nil, len(s)==0)

	fmt.Println()
	
	var t []int
	t = make([]int,3)
	t[0]=1
	fmt.Println("t",t)

	s=make([]string,3)
	fmt.Println("emp:",s,"len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println(s)
	fmt.Println(len(s))

	s = append(s, "d")
	s = append(s, "e","f")

	fmt.Println("s: ",s, "length: ",len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	l := s[2:5]
  fmt.Println("sl1: ", l)

	l = s[:5]
  fmt.Println("sl2:", l)

	l = s[2:]
  fmt.Println("sl3:", l)

	t1 := []string{"g", "h", "i"}
	fmt.Println("dcl:", t1)
	t2 := []string{"g", "h", "i"}
	fmt.Println("dcl:", t2)

	if slices.Equal(t1,t2){
		fmt.Println("t1==t2")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)
}

//output
/*
[]
0
uunit [] true true

t [1 0 0]
emp: [  ] len: 3 cap: 3
[a b c]
3
s:  [a b c d e f] length:  6
cpy:  [a b c d e f]
sl1:  [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
dcl: [g h i]
t1==t2
2d:  [[0] [1 2] [2 3 4]]
*/
