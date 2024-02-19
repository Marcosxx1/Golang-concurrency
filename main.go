package main

import (
	"fmt"
	"sync"
)

var (
	message  string
	wg       *sync.WaitGroup
	messages = []string{
		"Hello, universe!",
		"Hello, cosmos!",
		"Hello, world!",
	}
)

// UpdateMessage updates the global 'message' variable and signals
// that the goroutine is done using the WaitGroup.
func UpdateMessage(stringMessage string) {
	defer wg.Done()
	message = stringMessage
}

// PrintMessage prints the current value of the global 'message' variable.
func PrintMessage() {
	fmt.Println(message)
}

func main() {
	for _, msg := range messages {
		wg.Add(1)

		go UpdateMessage(msg)

		wg.Wait()

		PrintMessage()
	}
}

/* Explanation:
	1 - wg.Add(1) [LINE 32] increments the WaitGroup counter by 1 before each goroutine starts.
	The initial value of the counter is 0.

	2- go UpdateMessage(msg) [LINE 34] starts a new goroutine with the specified message. The defer wg.Done()
	in UpdateMessage will decrement the WaitGroup counter by 1 when the goroutine exits.

	3 - wg.Wait() [LINE 36] is used to wait until the WaitGroup counter becomes 0. This effectively waits for
	the currently running goroutine to finish before moving to the next iteration.

	4 - PrintMessage() [LINE 38] prints the current value of the message variable.

The loop then moves on to the next iteration, repeating the process for each message in the messages slice. */
