package main

import (
	"fmt"
)

func main(){
	// 固定長の配列arrays
	var a[3] int
	// 可変長の配列slice
	var s[] int

	a[0], a[1], a[2] = 1, 2, 3
	
	for i:=1; i<4; i++{
		s = append(s, i)
	}

	fmt.Println("arrays:",a)
	fmt.Println("slice:",s)
	
}