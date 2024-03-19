package lectures

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup){
	defer wg.Done()

	fmt.Println(s)
}

func Lectures() {

	var wg sync.WaitGroup

	words := []string{
		"um",
		"dois",
		"três",
	}

	wg.Add(len(words))

	for index, value := range words{
		go printSomething(fmt.Sprintf("%d: %s", index, value), &wg)
	}
	wg.Wait()
}
