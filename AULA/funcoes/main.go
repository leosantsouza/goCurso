package main

import (
	"errors"
	"fmt"
)

// func boasvindas()  {

// 	nome := "Leonardo"
// 	fmt.Printf("Bem vindo %s", nome)
// }

func soma(numero1 int, numero2 int)int{
	resultado := numero1 + numero2

	//fmt.Println("A soma dos numeros é: ",resultado)
	return resultado
}

type Usuario struct{
	Nome string
	Senha string
}

func autentica(user Usuario, senha string) (string, error) {

	if user.Senha != senha {
		return "", errors.New("Senha invalida")
	}
	//fmt.Println("Usuario1 logado é: ", user)
	return user.Nome, nil
}



func boasvindas(nome string)  {

		fmt.Printf("Bem vindo %s", nome)
}


func main()  {
	boasvindas("Leonardo \n")
	resultadoSoma := soma(30, 10)
	
	usuario1 := Usuario{Nome: "Leonardo", Senha: "1234"}

	nome, err := autentica(usuario1, "1234")

	if err != nil {
		fmt.Println("Erro na autenticação:", err)
	}else {
		fmt.Println("Usuario logado:", nome)
	}


	fmt.Println("Resultado da Soma é: ", resultadoSoma)
	
}