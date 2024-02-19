package goroutinesgoandwaitgroups

import (
	"fmt"
	"sync"
)

func PrintSomething(s string, wg *sync.WaitGroup) {
	//wg.Done() decrements the waitGroup by 1 each time it is called
	defer wg.Done() //defer:= whatever comes after this, don't execute until the current function exits
	fmt.Println(s)
}

func ChampterTwo() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gama",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	wg.Add(len(words))

	for index, value := range words {
		go PrintSomething(fmt.Sprintf("%d: %s", index, value), &wg)
	}

	wg.Wait()

	wg.Add(1)
	PrintSomething("second to be printed", &wg)

}
