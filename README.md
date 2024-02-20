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

- [1. Avoiding Deadlocks](#avoiding-deadlocks)
- [2. Handling Negative WaitGroup Counter](#handling-negative-waitgroup-counter)
- [3. Invalid memory address or nil pointer dereference](#3-invalid-memory-address-or-nil-pointer-dereference)

# Dealing with Waiting Groups in Go

When working with waiting groups in Go, it's important to manage the size of the waiting group appropriately to avoid potential errors.

## 1. Avoiding Deadlocks

### Issue:
If you have a slice with a length of 10, but add a waiting group of 11 without adjusting the waiting group size, you may encounter the following error:

```go
fatal error: All goroutines are asleep - deadlock
```

### Solution:
It is recommended to use wg.Add(len(name_of_slice)) to ensure the waiting group size matches the length of the slice. This way, you will have the exact waiting time needed.

## 2. Handling Negative WaitGroup Counter

### Issue:
When you have a slice of length 10 but fail to properly await one or more goroutines, you might encounter the following error:

```go
panic: sync: negative WaitGroup counter
```

### Solution:
Ensure the waiting group size corresponds to the number of goroutines that need to be awaited. If there is an additional goroutine that is not being properly awaited, adjust the waiting group size accordingly, for example, by using wg.Add(1).

## 3. Invalid memory address or nil pointer dereference

### Issue:
Can you find the problem here?

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