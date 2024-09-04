package main

import "fmt"

// Slices são mais flexiveis e poderosos do que arrays. Um slice não tem tamanho fixo, permitindo que voce adicione ou remova elementos dinamicamente

func main()  {

	var lista []int // Declarando um slice de inteiros

	fmt.Println("slices vazio", lista)

	lista = append(lista, 10, 50, 30)

	fmt.Println("Slice com valores", lista)

	// Criar slice a partir de um array existente.

	myArray := [7]int{1,2,3,4,5,6,7}

	fmt.Println("my array", myArray)

	mySlice := myArray[1:5] // Criar um slice incluindo elementos do indice 1 a 4

	fmt.Println("slice a partir do array", mySlice)

	var numeros []int

	numeros = append(numeros, 50, 10, 99, 30)

	fmt.Println("conteudo do slice", numeros)
	
// Usando make
	nomes := make([]string, 0)
	nomes = append(nomes, "Leonardo", "taynam")

	fmt.Println("nomes:", nomes)
}