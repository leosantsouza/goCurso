package main

import "fmt"


//Interface , semelhante a classes, conjunto de metodos posteriormente implementados por um type.
type Animal interface{
	EmitirSom()
	Andar()
}

//=========================================
// Implementando o metodo EmitirSom da Interface Animal no Type Cachorro.
func (c Cachorro) EmitirSom(){
	fmt.Println("Au Au Au!!")
}
// Implementando o metodo Andar da Interface Animal no Type Cachorro.
func (c Cachorro) Andar(){
	fmt.Println("Cachorro andando!!")
}

//==========================================
// Implementando o metodo EmitirSom da Interface Animal no Type Gato.
func (g Gato) EmitirSom(){
	fmt.Println("Miau Miau Miau!!")
}
// Implementando o metodo Andar da Interface Animal no Type Gato.
func (g Gato) Andar(){
	fmt.Println("Gato andando!!")
}

//=========================================
//Types

// Representa um cachorro.
type Cachorro struct{}
// Representa um gato.
type Gato struct{}

//=========================================
//Chama o metodo EmitirSom da interface Animal
func FazerEmitirSom(a Animal)  {
	a.EmitirSom()
}
//Chama o metodo Andar da interface Animal 
func FazerAndar(a Animal)  {
	a.Andar()
}

//=========================================

func main()  {
	
	cachorro1 := Cachorro{}
	gato1 := Gato{}

	FazerEmitirSom(cachorro1)
	FazerAndar(cachorro1)
	FazerEmitirSom(gato1)
	FazerAndar(gato1)
	

}