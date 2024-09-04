package main

import (
	"fmt"
	"math/rand"
	"time"
)



func PrepareOrder(orderID int, completedOrders chan string)  {
	fmt.Printf("Preparando o pedido #%d... \n", orderID)

	//Simulaçao de tempo de preparaçao (1 a 3 seg)
	time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
	//Enviando pedido para o channel
	completedOrders <- fmt.Sprintf("Pedido #%d concluido!!", orderID)
}


func main()  {
	//rand.NewSource(time.Now().UnixNano())

	orderCount := 4
	completedOrders := make(chan string, orderCount)

	// Lançar uma goroutine para cada pedido
	for i := 1; i <= orderCount; i++{
		go PrepareOrder(i, completedOrders)
	}

	for i := 1; i <= orderCount; i++ {
		fmt.Println(<- completedOrders)
	}

	//fechando o channel apos uso !
	close(completedOrders)
}