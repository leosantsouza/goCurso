package aluno

import "fmt"

//Tipo Aluno
type Aluno struct {
	Nome string
	N1   float64
	N2   float64
}
//Função para cadastro de aluno!
func Cadastro(aluno Aluno) {
	fmt.Printf("Aluno: %s foi cadastrado com sucesso! \n", aluno.Nome)
}
// Função para calculo de media!
func CalcularMedia(nota1 float64, nota2 float64) float64 {
	media := (nota1 + nota2) / 2

	return media
}

func VerificarAprovacao(media float64){
	if media >= 7.0 {
		fmt.Println("Aprovado")
	} else{
		fmt.Println("Reprovado")
	}

}