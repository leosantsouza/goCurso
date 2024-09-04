package main

import "fmt"



func main()  {

	//for i := 0; i < 5; i++ {
	//	fmt.Println(i)
	//}

	numero := 0

	for numero < 5 {
		fmt.Println("passando no for", numero)
		numero++
	}

	for i := 0; i < 5; i++ {
		
		if i == 3{
			fmt.Println("valor e igual a 3")
			continue
		}

		if i == 4{
			break
		}

		fmt.Println("valor do i", i)

	}

		
}