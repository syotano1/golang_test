package main

import (
	employee "golang_test/src/003_struct/struct"
)

func main(){
	e := employee.NewEmployee(0,"taro",25,"sales")
	e.DisplayInfo()
}