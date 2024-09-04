package main

import "fmt"


func somar(num *int){
	*num = 150
}


func main()  {
	numero := 10
	fmt.Println("Valor inicial da variavel numero:",numero)	

	somar(&numero)

	fmt.Println("Valor atual da variavel numero:",numero)

}