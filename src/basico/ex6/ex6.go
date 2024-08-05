package main

import "fmt"

func main() {
	x := -20

	if x > 0 {
		fmt.Println("Positivo")
	} else if x == 0 {
		fmt.Println("Zero")
	} else {
		fmt.Println("Negativo")
	}
}

//6. **Objetivo**: Utilizar estruturas de controle `if`.
//- **Tarefa**: Modifique a função `main` para verificar se um número é positivo, negativo ou zero e imprimir a
//mensagem correspondente.
//- **Pergunta**: Como Go lida com a condição `if` sem parênteses ao redor das condições?
