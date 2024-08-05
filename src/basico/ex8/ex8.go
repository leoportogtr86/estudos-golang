package main

func main() {
	x := 2

	switch x {
	case 1:
		println("x é 1")
	case 2:
		println("x é 2")
	case 3:
		println("x é 3")
	default:
		println("x não é 1, 2 ou 3")
	}
}

//8. **Objetivo**: Utilizar estruturas de controle `switch`.
//- **Tarefa**: Crie um `switch` que verifica o valor de uma variável e imprime uma mensagem para valores específicos.
//- **Pergunta**: Como o `switch` em Go difere do `switch` em outras linguagens?
