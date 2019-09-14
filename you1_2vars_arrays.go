package main

import "fmt"


func main (){

	b:=[]int32 {1,2,3,4,5}
	b = append(b,5,4,2,1,3,21,3)
	for range b {
		fmt.Println("a", b)
	}
	fmt.Println(b[len(b)-1])
	fmt.Println(b[0:4])

}