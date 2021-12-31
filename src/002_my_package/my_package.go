package main

import (
	"fmt"
	calc "golang_test/src/099_packages"
)

func main() {
	var f, s int

	fmt.Print("数値を２つ入力してください:")
	fmt.Scan(&f, &s)

	fmt.Println("f:",f,"s:",s)
	fmt.Println("sum:",calc.Plus(f,s))
}
