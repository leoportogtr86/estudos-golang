package main

import "fmt"

func main() {
	fmt.Println(saudacao("Leo"))
}

func saudacao(nome string) string {
	return "Olá, " + nome
}

//5. **Objetivo**: Definir e chamar funções personalizadas.
//- **Tarefa**: Crie uma função `saudacao` que recebe um nome como parâmetro e retorna uma saudação. Chame essa função
//na `main`.
//- **Pergunta**: Qual a diferença entre definir uma função fora e dentro da função `main`?
