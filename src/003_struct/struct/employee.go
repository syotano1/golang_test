package employee

import (
	"fmt"
)

type Employee struct{
	Id int
	Name string
	Age int
	Division string
}

func NewEmployee(Id int, Name string, Age int, Division string) *Employee{
	return &Employee{Id, Name, Age, Division}
}

func (e Employee) DisplayInfo(){
	fmt.Println(e.Id, e.Name, e.Age, e.Division)
}