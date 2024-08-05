package main

import "fmt"

func main() {
	var nome string = "John"
	var idade int = 45
	var salario float64 = 8790
	var casado bool = true
	teste := "teste"

	fmt.Println(nome, idade, salario, casado, teste)
	showMsg("Olá, mundo!")
}

func showMsg(msg string) {
	fmt.Println(msg)
}
