package main

import "fmt"

// Um ponteiro é uma variavel que armazena o endereço de memória de outra variável.

func main()  {
	var numero int = 60
	var p *int = &numero

	fmt.Println(numero)

	// Imprimir endereço de numero na memoria !
	fmt.Println("Endereço na memoria da var numero:", &numero)
	fmt.Println("Valor do ponteiro p:", p)

	// imprimir o valor do endereço armazenado em p (dereferencia)
	fmt.Println("Valor apontado por p:", *p)


}