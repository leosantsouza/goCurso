package main

import "fmt"

func main()  {
	
	nome1 := "Leonardo"
	nome2 := "Taynam"

	numero1 := 5
	numero2 := 10
	
	// Comparaçao
	fmt.Println(nome1 == nome2)
	fmt.Println(nome1 != nome2)
	fmt.Println(numero1 >= numero2)
	fmt.Println(numero1 > numero2)
	fmt.Println(numero1 <= numero2)
	fmt.Println(numero1 < numero2)


	// Tamanho da string

	fmt.Println(len(nome1) )

}