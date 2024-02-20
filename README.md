# Golang-concurrency
Go (Golang) simplifies concurrency with "go" for GoRoutines. Tools like WaitGroup, Mutex, and Channels manage concurrent tasks. This repo covers these in detail, addressing common issues. Classic problems illustrate concepts, and a real-world example demonstrates efficient concurrent task handling.


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