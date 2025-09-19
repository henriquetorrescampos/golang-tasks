// Arquivo principal do programa (entrypoint) 🫡
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

// Importa os pacotes necessários
import (
	"fmt"
	anonymousfunction "projeto-go/internal/anonymous-function"
	// "projeto-go/internal/anamnese"
	// memoriago "projeto-go/internal/memoria-go"
)

// Função principal do programa
func main() {
	fmt.Println("Hello, this is my golang project.")
	// Mensagem inicial da aplicação
	// fmt.Println("🚀 Meu primeiro projeto em Go com estrutura de mercado!")

	// Chamada para função de saudação
	// hello.SayHello()

	// Demonstração: cálculo do 10º número de Fibonacci
	// n := 10
	// Chama a função Fibonacci do pacote fibonacci
	// fibonacci // importado acima
	// Fibonacci(n) // retorna o n-ésimo número da sequência
	// := é usado para declarar e inicializar a variável
	// valor := fibonacci.Fibonacci(n)
	// Imprime o resultado com formatação
	// fmt.Printf("F(%d) = %d\n", n, valor)

	// Demonstração: imprimir a sequência completa até n
	// fibonacci.PrintSequence(n)
	// memoriago.LimparTela()

	// reader := bufio.NewReader(os.Stdin)

	// tamanhoDaSequencia := 4

	// fmt.Println("Iniciando a leitura do palpite...")
	// meuPalpite := memoriago.LerPalpite(reader, tamanhoDaSequencia)

	// fmt.Println("Seu palpite foi:", meuPalpite)

	// seq1 := []int{1, 2, 3, 10}

	// seq2 := []int{1, 9, 4}

	// result := memoriago.Pontuar(seq1, seq2)

	// fmt.Println(result)

	// anamnese.Anamense()

	/*
		função anônima é uma função que não possui nome
		é útil quando você quer criar um função de 1 linha.
		É conhecida por função literal.
	*/
	// func(a, b int, z float64) bool {
	// 	return (a*b < int(z))
	// }(10, 20, 20.2)
	fmt.Println(
		anonymousfunction.La(
			10, 20, 30.4))

}
