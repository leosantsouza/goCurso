package main

import "fmt"

func main()  {
	// declarar map e inicializar
	var capitais map[string]string = make(map[string]string)

	capitais["Brasil"] = "Brasilia"
	capitais["França"] = "Paris"
	capitais["Japão"] = "Tóquio"
	capitais["Alemanha"] = "Berlim"

	fmt.Println(capitais)
	fmt.Println(capitais["Alemanha"],",",capitais["Brasil"])

}