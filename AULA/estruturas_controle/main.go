package main

import "fmt"

// if else | else if

func main()  {
	nota1 := 3.0
	nota2 := 4.0

	media := (nota1 + nota2) / 2

	if media >= 7 {
		fmt.Printf("Aprovado, media %.2f", media)
	}else if media >= 5 && media <= 6.99 {
		fmt.Printf("Recuperaçao, media %.2f", media)
	}else {
		fmt.Printf("Reprovado, media %.2f", media)
	}

	fmt.Println(".")

	dia := 3

	switch dia {
	case 1:
		fmt.Println("Segunda-feira")
	case 2:
		fmt.Println("Terça-feira")
	case 3:
		fmt.Println("Quarta-feira")
	case 4:
		fmt.Println("Quinta-feira")
	case 5:
		fmt.Println("Sexta-feira")
	default:
		fmt.Println("Fim de semana")	
		
	}
	
}