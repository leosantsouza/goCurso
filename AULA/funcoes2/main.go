package main

import "fmt"

func novoUsuario(username string)(string, bool)  {
	if username != "" {
		return username, true
	}

	return "Nao cadastrado", false
}

func main()  {

	usuario1 := "Leonardo"

	resultado, status := novoUsuario(usuario1)
	resultado2, status2 := novoUsuario("")

	fmt.Printf("Usuario atual: %s, status atual: %t \n", resultado, status)
	fmt.Printf("Usuario atual: %s, status atual: %t \n", resultado2, status2)

	// função anônima
	func ()  {
		fmt.Println("Função anônima foi chamada")
	}() // <--- indica que foi criada e chamada imediatamente
}
