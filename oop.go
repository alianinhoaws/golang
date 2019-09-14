package oop

import "fmt"

type myObj struct {
	mystring string
	Myint int
}

func main () {
	newobj := myObj{
		mystring: "hello",
		Myint:    1,
	}
	fmt.Println(newobj)

	myObj{}.PrintString ("two")
}

//dosent work
func (m myObj)PrintString(str string){
	fmt.Println(m.mystring, str)
}
func (m *myObj)ChangeString(str string){
	fmt.Println(m.mystring, str)
}
