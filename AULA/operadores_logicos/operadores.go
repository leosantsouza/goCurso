package main

import "fmt"

func main()  {
	
	// and (&&)
	// or (||)
	// negaçao (!)

estoque := true

vendaLibarada := true

// para true ambas as partes precisam ser true.
fmt.Println(estoque && vendaLibarada)

// para true apenas um precisa ser true
fmt.Println(estoque || vendaLibarada)

// inverte 
fmt.Println(!estoque)


}