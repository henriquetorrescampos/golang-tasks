package anamnese

import (
	"fmt"
	"math"
)

func Anamense() {
	var peso float64
	var altura float64
	var nome string

	fmt.Print("Digite seu nome: ")
	fmt.Scanln(&nome)

	fmt.Print("Digite seu peso: ")
	fmt.Scanln(&peso)

	fmt.Print("Digite sua altura: ")
	fmt.Scan(&altura)

	imc := peso / math.Pow(altura, 2)

	if imc < 18.5 {
		fmt.Printf("%s, seu IMC é de %.2f e você está abaixo do peso.\n", nome, imc)
	} else if imc <= 24.9 {
		fmt.Printf("%s, seu IMC é de %.2f e você está com peso normal.\n", nome, imc)
	} else {
		fmt.Printf("%s, seu IMC é de %.2f e você está acima do peso.\n", nome, imc)
	}

}
