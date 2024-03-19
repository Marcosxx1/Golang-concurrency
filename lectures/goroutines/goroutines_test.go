package lectures

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {

	// Salvamos o estado atual de os.Stdout na variável stdOut
	stdOut := os.Stdout 

	// Criamos um Pipe para redirecionar a saída padrão para o nosso próprio pipe
	read, write, _ := os.Pipe()
	os.Stdout = write

	// Importante para sincronizar as goroutines
	var wg sync.WaitGroup												

	// Adicionamos 1 ao WaitGroup, indicando que uma goroutine está sendo executada
	wg.Add(1)																
	// Executamos a função printSomething em uma goroutine
	go printSomething("Valor", &wg)								

	// Esperamos que todas as goroutines no WaitGroup terminem
	wg.Wait()															

	// Fechamos o lado de escrita do pipe para liberar recursos
	_ = write.Close()

	// Lemos a saída do pipe
	result, _ := io.ReadAll(read)
	output := string(result)

	// Restauramos o os.Stdout original
	os.Stdout = stdOut

	// Verificamos se a saída contém a string esperada
	if !strings.Contains(output, "Valor") {
		t.Errorf("Expected to find Valor, but it is not there. Value found %s", output)
	}
}

