package main

import "fmt"

// Variaveis do tipo Var e const

func main()  {
	var nome = "Leonardo Santana"  // inferência de tipo automatica.
	var idade = 28
	var texto string

	fmt.Println("variaveis")
	fmt.Println(nome)
	fmt.Println(idade)


	idade = 29
	texto = "Valor atribuido"

	fmt.Println("Agora a idade e", idade)
	fmt.Println(texto)
}