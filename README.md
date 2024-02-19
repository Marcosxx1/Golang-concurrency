# Golang-concurrency
Go (Golang) simplifies concurrency with "go" for GoRoutines. Tools like WaitGroup, Mutex, and Channels manage concurrent tasks. This repo covers these in detail, addressing common issues. Classic problems illustrate concepts, and a real-world example demonstrates efficient concurrent task handling.


When dealing with waiting groups, let's say, we have a slice that has length of 10, but add a waiting group of 11 we'll have the error

fatal error: All goroutines are asleep - deadlock

it is always better to use wg.Add(len(name_of_slice)) this way we'll have the exact waiting time

on the other hand, if we have a slice of length 10, but have one more goroutine that is not beeing properly awaited we'll have this error:

panic: sync: negative WaitGroup counter

again, we need to have the proper awaiting size to the waiting group, adding one more or where it is needed	wg.Add(1)



panic: runtime error: invalid memory address or nil pointer dereference