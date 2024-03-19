package challengeoneprintsomething

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_printMessge(t *testing.T) {
	sdtOut := os.Stdout

	read, write, _ := os.Pipe()
	os.Stdout = write

	msg = "Hello, world!"
	printMessage()

	_ = write.Close()

	result, _ := io.ReadAll(read)
	output := string(result)

	os.Stdout = sdtOut

	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("Expected Hello, world!, but found %s:", output)
	}
}

func Test_updateMessage(t *testing.T) {

	wg.Add(1)

	go updateMessage("Two")

	wg.Wait()

	if msg != "Two" {
		t.Errorf("Expected to Two Valor, but it is not there. Value found %s", msg)
	}
}

func Test_ChallengeOne(t *testing.T) {
	stdOut := os.Stdout
	read, write, _ := os.Pipe()

	os.Stdout = write

	ChallengeOne()

	_ = write.Close()

	result, _ := io.ReadAll(read)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("Expected:Hello, universe!, actual: %s", output)
	}
	if !strings.Contains(output, "Hello, cosmos") {
		t.Errorf("Expected: Hello, cosmos, actual: %s", output)
	}
	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("Expected: Hello, world!, actual: %s", output)
	}
}
// go test -v .