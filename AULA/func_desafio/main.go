package main

import "fmt"


type Aluno struct{
	Nome string
	Nota1 float64
	Nota2 float64
}

// calculaMedia calculates the average grade of a student based on two given grades.
//
// Parameters:
// aluno (Aluno): A struct representing a student with fields Nome (string) and Nota1, Nota2 (float64).
//
// Returns:
// string: The student's name.
// float64: The calculated average grade.
func calculaMedia(aluno Aluno) (string, float64) {
	media := (aluno.Nota1 + aluno.Nota2) / 2

	return aluno.Nome, media
}

// verificaAprovacao checks if a student's average grade is sufficient for approval.
//
// Parameters:
// media (float64): The calculated average grade of the student.
//
// Returns:
// string: A message indicating whether the student has been approved or is in recovery.
func verificaAprovacao(media float64) string {
	if media >= 7 {
		return "Parabens aluno aprovado!"
	}
	return "Ops, você esta em recuperaçao!!"
}

func main()  {

	aluno1 := Aluno{Nome: "Leonardo", Nota1: 6.0, Nota2: 9.5}

	fmt.Printf("%s, Nota1: %.1f, Nota2: %.1f \n", aluno1.Nome, aluno1.Nota1, aluno1.Nota2)

	nome, media := calculaMedia(aluno1)

	fmt.Printf("Media do %s: %.2f \n", nome, media)

	statusAluno1 := verificaAprovacao(media)

	fmt.Println(statusAluno1)
}