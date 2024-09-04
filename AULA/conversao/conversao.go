package main

import (
	"fmt"
	"strconv"
)


func main()  {
	// conversao int/float
	var valor int = 42
	var valorConvertido float64 = float64(valor)

	fmt.Println("inteiro", valor)
	fmt.Println("convertido em float", valorConvertido)

	// conversao int/string
	var valorString string = strconv.Itoa(valor)

	fmt.Println(valorString)

	// convesao string/int
	var valor2 string = "230"
	var valorInteiro, err = strconv.Atoi(valor2)

	if err != nil {
		fmt.Println("ERRO AO CONVERTER")
	} else{
		fmt.Println(valorInteiro + 10)
	}

	fmt.Println(valor2)

	// conversao string/float

	var pi string = "3.14159"
	var piConvertido, error2 = strconv.ParseFloat(pi, 64)
	
	if error2 != nil {
		fmt.Println("ERRO AO CONVERTER")
	} else{
		fmt.Println(piConvertido)
	}
	
}