package main

import (
	"fmt"
	"strconv"
)

const (
	_ = iota
	b = iota * 10
	c
)

func main (){
	str1 := "Hello"
	str3 := `{hello2}`
	str2 := 's'
	str4 := `{"adadad":"2313"}`
	arrayByte := [] byte (str4)
	var value2 int = 32

	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(str1+" "+str3)
	fmt.Println(str2)
	fmt.Println(arrayByte)

	num,err := strconv.Atoi("20")
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println(num)

	value := strconv.Itoa(b)
	fmt.Println(value)

	value3 := strconv.Itoa(value2)
	fmt.Println(value3)

}
