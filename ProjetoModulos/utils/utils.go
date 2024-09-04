package utils

import "fmt"

// função para exibir mensagem de boas vindas!
func Mensagens(nome string)  {
	status := statusSistema()
	fmt.Printf("Ola!! %s, seja bem vindo!!\n", nome)
	fmt.Println(status)
}

func statusSistema() string{
	return "Stauts OK!!!!!"
}