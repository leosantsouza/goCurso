package main

import "fmt"


func main()  {

	var tarefas []string

	tarefas = append(tarefas, "Estudar", "Jogar", "Ler um livro", "tocar", "sair")
	fmt.Println("Tarefas: ", tarefas)	
	fmt.Println("tamanho do slice", len(tarefas))

	// Slicing
	tarefas = tarefas[1:]

	fmt.Println("Removido a primeira tarefa ", tarefas)

	tarefas = append(tarefas[:1], tarefas[2:]...)

	fmt.Println(tarefas)
}