package mutexraceconditions

import (
	"fmt"
	"sync"
)

var (
	msg   string
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()

	m.Lock()
	msg = s
	m.Unlock()
}

func Mutexraceconditions() {
	msg = "Hi there!"

	wg.Add(2)
	go updateMessage("Hello, world!", &mutex)
	go updateMessage("Hello, universe!", &mutex)
	wg.Wait()

	fmt.Println(msg)
}
