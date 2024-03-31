package bankbalance

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_BankBalance(t *testing.T) {
	stdOut := os.Stdout
	read, write, _ := os.Pipe()

	os.Stdout = write

	BankBalance()

	_ = write.Close()

	result, _ := io.ReadAll(read)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "34320.00") {
		t.Errorf("wrong balance returned: %v", output)
	}
}
