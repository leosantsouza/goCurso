package main

import "fmt"

type Pessoa struct {
	Nome string
	Idade int
	Nacionalidade string

}

func main()  {
	
	leo := Pessoa{Nome: "Leonardo Santana", Idade: 44, Nacionalidade: "Brasil"}
	tay := Pessoa{Nome: "Taynan Ellem", Idade: 37, Nacionalidade: "Brasil"}

	fmt.Println(leo)
	fmt.Println(leo.Nome)
	fmt.Println(tay)
	fmt.Println(tay.Nome)

// Map + Struct
	pessoas := map[int]Pessoa{
		1:{Nome: "Rayssa", Idade: 25, Nacionalidade: "Brasil"},
		2:{Nome: "Gesom", Idade: 26, Nacionalidade: "Brasil"},
	}
	fmt.Println(pessoas[1])

}