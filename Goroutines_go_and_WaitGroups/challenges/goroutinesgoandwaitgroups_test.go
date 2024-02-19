package goroutinesgoandwaitgroups

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSometh(t *testing.T){
	stdOut := os.Stdout

	read, write, _ :=os.Pipe()
	os.Stdout = write

	var wg sync.WaitGroup
	wg.Add(1)

	go PrintSomething("alpha", &wg)

	wg.Wait()

	_ = write.Close()

	result, _ := io.ReadAll(read)

	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "alpha"){
		t.Errorf("Expected to find alpha, but it was not found")
	}
}