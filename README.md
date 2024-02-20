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