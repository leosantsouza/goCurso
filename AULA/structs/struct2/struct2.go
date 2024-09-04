package main

import "fmt"

type Usuario struct {
	Nome string
	Idade int
	status bool
	permissao string

}

func main()  {
	
	usuarios := map[int]Usuario{

		1:{Nome: "Leonardo", Idade: 44, status: true, permissao: "admin"},
		2:{Nome: "Taynam", Idade: 37, status: true, permissao: "user"},
		3:{Nome: "Leon", Idade: 30, status: true, permissao: "user"},
		4:{Nome: "Rayssa", Idade: 26, status: false, permissao: "estagiario"},

	}
	

	i := 1
	for i <= len(usuarios){
		fmt.Printf(" Nome: %s, Idade: %d, Permissão Atual: %s  \n ", usuarios[i].Nome, usuarios[i].Idade, usuarios[i].permissao)
		i++
	}


	//fmt.Println(usuarios)
}