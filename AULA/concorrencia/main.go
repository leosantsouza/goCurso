package main

import (
	"fmt"
	"time"
)



func numeros()  {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d",i)
		time.Sleep(time.Millisecond * 200)
	}
}
func letras()  {
	for l := 'a'; l < 'j'; l++ {
		fmt.Printf("%c", l)
		time.Sleep(time.Millisecond * 200)
	}
}





func main()  {
	go letras()
	go numeros()

	time.Sleep(3 * time.Second)
	fmt.Println("\n Finalizado!!!")
}