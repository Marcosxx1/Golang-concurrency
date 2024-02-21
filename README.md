# Golang-concurrency
# Learning Go Concurrency: Documentation Repository

In the domain of Golang projects, I've explored goroutines, channels, and sync's waitGroups. To enhance my understanding, I've dedicated this repository to studying Go's concurrency features.

## Purpose

This repository serves as a comprehensive record of my exploration into Go's concurrency mechanisms. Beyond code, it's a documentation hub for my learning journey, chronicling challenges, and resolving errors.

## Rationale for Study

Studying is essential for navigating the complexities of Go's concurrency. The approach here acknowledges that errors are inevitable. Instead of ignoring them, this repository highlights not only correct solutions but also errors, missteps, pitfalls, and the lessons derived.

## Contents

- **Code Exploration:** In-depth examinations of goroutines, channels, and waitGroups through code snippets and examples.

- **Error Documentation:** Transparent documentation of encountered errors. Understanding and learning often stem from addressing challenges.

- **Learning Pathway:** A structured progression through Go's concurrency features, catering to both novices and adept developers.



Go (Golang) simplifies concurrency with "go" for GoRoutines. Tools like WaitGroup, Mutex, and Channels manage concurrent tasks. This repo covers these in detail, addressing common issues. Classic problems illustrate concepts, and a real-world example demonstrates efficient concurrent task handling.

- [1.1 Avoiding Deadlocks](#avoiding-deadlocks)
- [1.2 Handling Negative WaitGroup Counter](#handling-negative-waitgroup-counter)
- [1.3 Invalid memory address or nil pointer dereference](#3-invalid-memory-address-or-nil-pointer-dereference)

# Dealing with Waiting Groups in Go

When working with waiting groups in Go, it's important to manage the size of the waiting group appropriately to avoid potential errors.

## 1.1 Avoiding Deadlocks

### Issue:
If you have a slice with a length of 10, but add a waiting group of 11 without adjusting the waiting group size, you may encounter the following error:

```go
fatal error: All goroutines are asleep - deadlock
```

### Solution:
It is recommended to use wg.Add(len(name_of_slice)) to ensure the waiting group size matches the length of the slice. This way, you will have the exact waiting time needed.

## 1.2 Handling Negative WaitGroup Counter

### Issue:
When you have a slice of length 10 but fail to properly await one or more goroutines, you might encounter the following error:

```go
panic: sync: negative WaitGroup counter
```

### Solution:
Ensure the waiting group size corresponds to the number of goroutines that need to be awaited. If there is an additional goroutine that is not being properly awaited, adjust the waiting group size accordingly, for example, by using wg.Add(1).

## 1.3 Invalid memory address or nil pointer dereference

### Issue:


<details>
<summary><strong></strong><em>(click me) Can you find the problem here?</em></summary>
<br>

```go
var (
	message  string
	wg       *sync.WaitGroup
	messages = []string{
		"Hello, universe!",
		"Hello, cosmos!",
		"Hello, world!",
	}
)

func UpdateMessage(stringMessage string) {
	defer wg.Done()
	message = stringMessage
}

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

```
 
</details>

If we run the above code wi'll have 
```powershell
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xc0000005 code=0x1 addr=0x0 pc=0x7adfc]
```
Pointing to this line:
```go
		wg.Add(1) // inside the for loop
```

### Solution:
We reclared our waitGroup [ wg  *sync.WaitGroup] but anywhere in the code we initialized it 

In Go, the zero value of a pointer is nil. Therefore, when we try to add to the WaitGroup using wg.Add(1) before initializing wg, it results in a nil pointer dereference.

To fix this, we need to initialize the wg variable before using it. You can do this by assigning a new instance of sync.WaitGroup to wg before the loop.

```go 
...func main() {
	wg = &sync.WaitGroup{} // Initialize the WaitGroup

	for _, msg := range messages {
		wg.Add(1)
    ...
```
Now, wg is properly initialized, and the program should run without the nil pointer dereference error.

# 2.0 Understanding `sync.Mutex` and Channels

## `sync.Mutex`

- `Mutex`: "mutual exclusion" - allows us to deal with race conditions.
- Relatively simple to use.
- Useful for dealing with shared resources and concurrent/parallel goroutines.
- Provides `Lock` and `Unlock` methods for managing access to critical sections.
- Enables testing for race conditions when running code or testing.

# Concurrent Access with Mutex in Go

In the example bellow, even though we have two concurrent calls to `updateMessage`, we're accessing `packageMessage` safely using a `Mutex`.

If we were not using the `Lock()` and `Unlock()` methods to synchronize access to `packageMessage`, we would have a race condition where multiple goroutines could simultaneously update the `packageMessage` variable without coordination. In such a scenario, we wouldn't know which goroutine's update would prevail, leading to unpredictable and potentially incorrect results.

It is worth noting that when a goroutine calls `Lock()`, it acquires the lock, and any other goroutine attempting to acquire the lock will be blocked until the first goroutine releases it with `Unlock()`. This ensures EXCLUSIVE ACCESS to the shared variable during the critical section.

<details>
<summary><strong></strong><em>(click me) Here's the example</em></summary>
<br>

```go
// packageMessage is a shared variable that will be updated by multiple goroutines.
var packageMessage string

// wg is a WaitGroup to wait for the completion of goroutines.
var wg sync.WaitGroup

// updateMessage updates the packageMessage variable with the provided incomingMessage.
// It uses a Mutex for synchronization to avoid race conditions.
func updateMessage(incomingMessage string, m *sync.Mutex) {
	defer wg.Done()

	// Lock the Mutex to ensure exclusive access to the critical section.
	m.Lock()
	packageMessage = incomingMessage
	// Unlock the Mutex to release the lock after updating the shared variable.
	m.Unlock()
}

func main() {
	// Initialize packageMessage with an initial value.
	packageMessage = "Hi, there!"

	// Create a Mutex to synchronize access to packageMessage.
	var mutex sync.Mutex

	// Add 2 to the WaitGroup to wait for the completion of two goroutines.
	wg.Add(2)

	// Launch two goroutines to update packageMessage concurrently.
	go updateMessage("Hi!", &mutex)
	go updateMessage("Hi for everyone!", &mutex)

	// Wait for both goroutines to complete.
	wg.Wait()

	// Print the final value of packageMessage.
	fmt.Println(packageMessage)
}
```
</details>

<br><hr>





## Race Conditions

- Race conditions occur when multiple goroutines try to access the same data concurrently.
- They can be difficult to spot when reading code.
- Go provides mechanisms to check for race conditions when running a program or testing code using `go test`.

## Channels

- Channels are a means of having goroutines share data.
- Goroutines can communicate with each other through channels.
- Reflects Go's philosophy of having things share memory by communicating, rather than communicating by sharing memory.
- Addresses the Producer/Consumer problem.
