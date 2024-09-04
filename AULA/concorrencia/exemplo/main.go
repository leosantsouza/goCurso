package main

import "fmt"


func soma(slice []int, channel chan int)  {
	sum := 0
	for _, value := range slice{
		sum += value
	}
	channel <- sum
}

func main()  {
	slice := []int{1,2,3,4,5,6,7,8,9,10}

	channel := make(chan int)

 // dividir o slice em duas partes

	mid := len(slice) / 2
	p1 := slice[:mid]
	p2 := slice[mid:]

	go soma(p1, channel)
	go soma(p2, channel)

	//receber resultado dos channels
	soma1 := <- channel
	soma2 := <- channel

	totalSoma := soma1 + soma2

	fmt.Printf("Total da soma: %d  ",totalSoma)



	// for _, value := range p1{
	// 	fmt.Printf("Valores: %d \n", value)
	// }
	// fmt.Println("----------------")
	// for _, value := range p2{
	// 	fmt.Printf("Valores: %d \n", value)
	// }


}


