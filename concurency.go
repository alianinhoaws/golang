package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		count("sheep")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		count("shark")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		count("Elephant")
	}()
	wg.Wait()
}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
	}
}
