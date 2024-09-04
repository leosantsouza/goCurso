package main

import "fmt"

// string
// bool > True and False
// int | int8, int16, int64, uint, uint16, uint64
// float32, float64
// Bytes, Runes

func main()  {

	var nome string = "Meu nome"
	var isYou bool = true
	var idade int = 44
	var valor float32 = 19.90
	var teste byte = 'a'
	var runteste rune = '&'


	fmt.Println(nome)
	fmt.Println(isYou)
	fmt.Println(idade)
	fmt.Println(valor)
	fmt.Println(teste)
	fmt.Println(runteste)

}