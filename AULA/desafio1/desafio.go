package main

import "fmt"

func main()  {

	// Idade atual com base no ano de nascimento e o ano atual
	var anoNascimento int = 1980
	var anoAtual int = 2024

	var idadeAtual int = anoAtual - anoNascimento

	fmt.Println(idadeAtual)



	// Converter horas em minutos
	var horas int = 1
	var totalMinutos int = horas * 60
	
	fmt.Println(totalMinutos)
	fmt.Printf("%d hora tem %d minutos", horas, totalMinutos)
	fmt.Println(".")



	// Calculadora de Juros simples.
	

	valorCapital := 2000
	taxaJuros := 0.03
	periodoMes := 12



	var jurosFinal float64 = float64(valorCapital) * taxaJuros * float64(periodoMes)

	fmt.Printf("testando o %.2f", jurosFinal)

	
}