package main

import "fmt"

// Arrays > Um array em Go é uma sequência de elementos de um único tipo, com um tamnho fixo definido na declaração.

func main()  {

	var numeros [4]int

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 15
	numeros[3] = 05


	fmt.Println(numeros)

	var valores = [3]int{4, 3, 2}

	fmt.Println(valores)

	nomes := [3]string {"Leonardo", "Rayssa", "Taynam"}

	fmt.Println(nomes)

}