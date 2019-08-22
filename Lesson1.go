package lesson1

import (
	"fmt"
)

func main () {
	//var+if/else
	x := 2
	if x > 6 {
		fmt.Printf("Hello")
	} else {
		fmt.Printf("NO")
	}
	//array
	a := [5]int {0,0,7,0,0}
	fmt.Println (a)
	//dynamic array
	b_slice := [] int {5,4,3,2,1}
	b_slice = append(b_slice, 22,32 )
	fmt.Println(b_slice)
	//dictionary (maps)
	dic := make(map[string]int)
	dic["first"]=1
	dic["second"]=2
	dic["third"]=3

	fmt.Println(dic)
	fmt.Println(dic["first"])
	delete(dic, "second")
	fmt.Println(dic)

	//loops
	for i:= 0; i<5;i++{
		fmt.Println(i)
	}
	//or like while
	i:=0
	for i<5{
		fmt.Println(i)
		i++
	}
	//array iteration
	arr:= [] string{"a","b","c"}
	for index, value := range arr{
		fmt.Println("index", index, "value", value)
	}
	//map iteration
	new_map := make(map[string]string)
	new_map["priest"] = "one"
	new_map["warlock"] = "second"
	new_map["mage"] = "third"
	for key, value := range new_map {
		fmt.Println("key", key, "value", value)
	}

}
