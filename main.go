package main

import (
	"fmt"
	"sync"
)

var packageMessage string
var wg sync.WaitGroup


/* In this example, even though we have two concurrent calls to updateMessage on LINE 13 and LINE 14,
	 we're accessing packageMessage safely using a Mutex.

	 If we were not using the Lock() and Unlock() methods to synchronize access to packageMessage,
	 we would have a race condition where multiple goroutines could simultaneously update the
	 packageMessage variable without coordination. In such a scenario, we wouldn't know which
	 goroutine's update would prevail, leading to unpredictable and potentially incorrect results.

	 Is it worthing noting that when a goroutine calls Lock(), it acquires the lock, and any other
	 goroutine attempting to acquire the lock will be blocked until the first goroutine releases it
	 with Unlock(). This ensure EXCLUSIVE ACCESS to the shared variable during the critical section
*/

func updateMessage(incomingMessage string, m *sync.Mutex) {
	defer wg.Done()

	m.Lock()
	packageMessage = incomingMessage
	m.Unlock()
}

func main() {
	packageMessage = "Hi, there!"

	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Hi!", &mutex)
	go updateMessage("Hi for everyone!", &mutex)
	wg.Wait()

	fmt.Println(packageMessage)

}

/*
	 sync.Mutex

	* Mutex := "mutual exclusion" -- allow us to deal with race conditions
	* Relative simple to use
	* Dealing with shared resources and concurrent/parallel goroutines
	* Lock/Unlock
	* We can test for race conditions when running code, or testing

	 Race Conditions
	 * Race conditions occur when multiple GoRoutines try to access the same data
	 * Can be difficult to spot when reading code
	 * Go allow us to check for them when running a program, or when testing our code with go test

	 Channels
	 * Channels are a means of having GoRoutines share data
	 * They can talk to each other
	 * This is Go's philosophy of having things share memory by communicating, rathen than communicating by sharing memory
	 * The Producer/Consumre problem
*/
