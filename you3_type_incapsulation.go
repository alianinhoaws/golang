package main

import (
	"auxilary"
	"fmt"
)

type myownint uint8

func main(){
	auxilary.A = 2
	fmt.Println(auxilary.A)
	fmt.Println(auxilary.Add2values(22,44))
}

