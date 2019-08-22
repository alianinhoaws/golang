package main

import (
	"errors"
	"fmt"
	"math"
)

func main () {
	result, err := sqrt(-2)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	// Struct
	p:= person{name: "andrii", age: 25}
	fmt.Println(p)

	//Pointers
	i:=7
	inc(&i)
	fmt.Println(i)

	//sort array
	arr2 := []int{2,6,1,3,5,1,2}
	sort(arr2)
	max:=2
	fmt.Println(max)

}

func inc (x *int){
	*x++
}

func sqrt (x float64) (r float64, err error){
	if x < 0 {
		err = errors.New ("Under 0")
	} else {
		r = math.Sqrt(x)
	}
	return
}
	type person struct {
		name string
		age int
	}
	//dosent work=)
	func sort (l [] int) (max int){
		max = 0
		for x := range l {
			if x > max {
				max = x
			}
		}
		return max
	}
