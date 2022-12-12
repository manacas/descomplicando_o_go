package main

import "fmt"

func somar(a, b int) int {
	return a + b
}

func diminuir(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func dividir(a, b int) float64 {
	if b == 0 {
		return 0 // licenca poetica de exercicio de programacao
	}
	return float64(a) / float64(b)
}

func main() {
	// #desafiodocursodescomplicandoogo
	// fmt.Scanf

	var valor_1 int
	var valor_2 int
	var operacao int

	fmt.Println("Entre com o primeiro número: ")
	fmt.Scanf("%d", &valor_1)
	fmt.Println("Entre com o segundo número: ")
	fmt.Scanf("%d", &valor_2)
	fmt.Println("Escolha a operação:\n1) Soma\n2) Subtração\n3) Multiplicação\n4) Divisão")
	fmt.Scanf("%d", &operacao)

	switch operacao {
	case 1:
		fmt.Println("Resultado da soma: ", somar(valor_1, valor_2))
	case 2:
		fmt.Println("Resultado da subtração: ", diminuir(valor_1, valor_2))
	case 3:
		fmt.Println("Resultado da multiplicação: ", multiplicar(valor_1, valor_2))
	case 4:
		fmt.Println("Resultado da divisão: ", dividir(valor_1, valor_2))
	default:
		fmt.Println("Operação inválida")
	}

}
