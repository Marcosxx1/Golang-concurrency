package main

import (
	"fmt"
	"sync"
)

var (
	message string
)

func UpdateMessage(stringMessage string, wg *sync.WaitGroup) {
	defer wg.Done()
	message = stringMessage
}

func PrintMessage() {
	fmt.Println(message)
}

func main() { 
	var wg sync.WaitGroup
	
	message = "Hello, world!"

	wg.Add(1)
	go UpdateMessage("Hello, universe!", &wg)
	wg.Wait()
	PrintMessage()

	wg.Add(1)
	go UpdateMessage("Hello, cosmos!",&wg)
	wg.Wait()
	PrintMessage()

	wg.Add(1)
	go UpdateMessage("Hello, world!",&wg)
	wg.Wait()
	PrintMessage()
	
 }
