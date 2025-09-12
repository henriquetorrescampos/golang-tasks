package memoriago

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// func GenerateSequence(size, maxNumber int) []int {
// 	seq := make([]int, size)

// 	for i := 0; i < size; i++ {
// 		seq[i] = rand.Intn(maxNumber + 1)
// 	}
// 	return seq
// }

// func LimparTela() {
// 	// Tenta limpar via ANSI + garante espaço em branco como fallback
// 	fmt.Print("\033[H\033[2J")
// 	fmt.Print(strings.Repeat("\n", 10))
// }

func LerPalpite(reader *bufio.Reader, tamanho int) []int {
	for {
		fmt.Printf("Digite a sequência (%d números de 0 a 9, separados por espaço): ", tamanho)
		linha, _ := reader.ReadString('\n')
		linha = strings.TrimSpace(linha)

		partes := strings.Fields(linha)
		if len(partes) != tamanho {
			fmt.Printf("Ops! São exatamente %d números. Tenta de novo.\n", tamanho)
			continue
		}

		palpite := make([]int, tamanho)
		ok := true
		for i, p := range partes {
			v, err := strconv.Atoi(p)
			if err != nil || v < 0 || v > 9 {
				fmt.Println("Use apenas dígitos de 0 a 9.")
				ok = false
				break
			}
			palpite[i] = v
		}
		if ok {
			return palpite
		}
	}
}

func Pontuar(seq, palpite []int) int {
	acertos := 0
	for i := range seq {
		if seq[i] == palpite[i] {
			acertos++
		}
	}
	return acertos
}
