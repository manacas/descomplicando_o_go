package main

import "fmt"

func somar(a, b int) int {
	return a + b
}

func dividir(a, b int) int {
	if b == 0 {
		return 0 // licenca poetica de exercicio de programacao
	}
	return a / b
}

func diminuir(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func main() {
	// TODO: Pedir os numeros para o usuario?
	// #desafiodocursodescomplicandoogo - fmt.Scanf
	// fmt.Scanf

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			fmt.Println(i)
			fmt.Println(j)

			resposta := somar(i, j)
			fmt.Printf("Soma: %d\n", resposta)
			resposta = dividir(i, j)
			fmt.Printf("Divisão: %d\n", resposta)
			resposta = diminuir(i, j)
			fmt.Printf("Subtração: %d\n", resposta)
			resposta = multiplicar(i, j)
			fmt.Printf("Multiplicação: %d\n", resposta)
		}
	}
}
