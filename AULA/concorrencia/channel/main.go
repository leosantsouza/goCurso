package main

import "fmt"



func sayHello(chname chan string)  {
	chname <- "Bem vindo"
}

func main() {

	ch := make(chan int)
	chname := make(chan string)

	go sayHello(chname)

	go func() {
		ch <- 123
	}()

	valor := <- ch

	fmt.Println("\n",valor)

	msg := <- chname

	fmt.Print(msg)

}
