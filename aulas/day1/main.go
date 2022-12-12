package main

import "fmt"

// snake-case: dummy_function
// CamelCase: dummyFunction

// privada, dentro do meu pacote
func dummyFunction(s string) string {
	return fmt.Sprintf("Run dummy function. %s\n", s)
}

// publica, dentro de todo o codigo
func DummyPublicFunction(s string) string {
	return dummyFunction(s)
}

func main() {
	// s := dummyFunction("Hello World!")

	// atribuição curta de variável (declaração + atribuição)
	// short assignment
	s := DummyPublicFunction("Hello World!!!!")
	fmt.Println(s)

	var u string
	u = dummyFunction("Hello World again!")
	fmt.Println(u)

	// Controle de fluxo
	// if, else, switch
	var input int
	fmt.Scan(&input)
	if input == 1 {
		fmt.Println("Input is 1.")
	} else if input == 2 {
		// do something
	} else {
		fmt.Println("Input is not 1.")
	}

	switch input {
	case 1:
		fmt.Println("Input is 1.")
	case 2:
		fmt.Println("Input is 2.")
		// não precisa de break!!
	default:
		fmt.Println("Input is not 1.")
	}

	// Estrutura de repetição
	// For
	// Laço - loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("----")

	var i int
	for i = 0; i < 10000; i = i + 2 {
		fmt.Println(i)
	}
	fmt.Println("----")

	// WHILE?
	// Loop infinito com condição de quebra
	// GO NÃO TEM WHILE!
	// Mas tem for que se comporta como while!
	running := true
	for running {
		// do something
		if input == 3 {
			running = false
		}
	}
	fmt.Println(input)
}
