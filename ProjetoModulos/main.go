package main

import (
	"fmt"
	"projetoModulos/aluno"
	"projetoModulos/utils"
)

// go mod init


func main()  {

	utils.Mensagens("Leonardo")

	aluno1 := aluno.Aluno{Nome: "Leonardo", N1: 6.5, N2: 6.5}
	aluno.Cadastro(aluno1)
	mediaAluno1 := aluno.CalcularMedia(aluno1.N1, aluno1.N2)

	fmt.Println("Media do aluno1:",mediaAluno1)
	aluno.VerificarAprovacao(mediaAluno1)
	

}