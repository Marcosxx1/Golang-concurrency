package main

/* package challenges

// Importa os pacotes necessários para o teste.
import (
	"io"
	"os"
	"testing"
)

// Função que captura a saída padrão enquanto uma função é executada.
func captureOutput(f func()) string {
	// Salva o valor original da saída padrão.
	originalStdout := os.Stdout

	// Cria um pipe para redirecionar a saída padrão.
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Executa a função fornecida, redirecionando a saída para o pipe.
	f()

	// Fecha o extremo de escrita do pipe.
	_ = w.Close()

	// Lê toda a saída do extremo de leitura do pipe.
	out, _ := io.ReadAll(r)

	// Restaura a saída padrão para o valor original.
	os.Stdout = originalStdout

	// Retorna a saída capturada como uma string.
	return string(out)
}

// Função de teste para a função Challenges.
func TestChallenges(t *testing.T) {
	// Define a saída esperada para a função Challenges.
	expectedOutput := "Hello, universe!\nHello, cosmos!\nHello, world!\n"

	// Chama a função captureOutput, passando a função Challenges, e armazena a saída capturada.
	actualOutput := captureOutput(Challenges)

	// Verifica se a saída capturada é igual à saída esperada.
	if actualOutput != expectedOutput {
		// Gera uma mensagem de erro se as saídas são diferentes.
		t.Errorf("Saída esperada:\n%s\nSaída atual:\n%s", expectedOutput, actualOutput)
	}
}
*/