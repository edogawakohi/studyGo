package main

import (
	"fmt"
	"time"
)

func main() {

	cha1 := make(chan int)
	cha2 := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		cha1 <- 1
	}()

	go func() {
		time.Sleep(1 * time.Second)
		cha2 <- 2
	}()

	// fmt.Println(<-cha1)
	// fmt.Println(<-cha2)

	for i := 1; i <= 2; i++ {
		select {
		case msg := <-cha1:
			fmt.Println(msg)
		case msg2 := <-cha2:
			fmt.Println(msg2)
		}
	}

}
