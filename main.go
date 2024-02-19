package main

import (
	"fmt"
	"sync"
)

var (
	message string
)

func UpdateMessage(stringMessage string) {
	defer wg.Done()
	message = stringMessage
}

func PrintMessage() {
	fmt.Println(message)
}

var wg sync.WaitGroup

func main() {

	message = "Hello, world!"

	wg.Add(1)
	go UpdateMessage("Hello, universe!")
	wg.Wait()
	PrintMessage()

	wg.Add(1)
	go UpdateMessage("Hello, cosmos!")
	wg.Wait()
	PrintMessage()

	wg.Add(1)
	go UpdateMessage("Hello, world!")
	wg.Wait()
	PrintMessage()
}
