package main

import (
	"context"
	"fmt"
	"time"
)

func cookPho(ctx context.Context, chPho chan<- string) {
	fmt.Println(" Start cook pho")

	select {
	case <-time.After(1 * time.Second):
		chPho <- "Phở đã nấu xong"
	case <-ctx.Done():
		fmt.Println("Cancel!")
		return
	}
}

func cookBun(ctx context.Context, chBun chan<- string) {
	fmt.Println(" Start cook bún")

	select {
	case <-time.After(2 * time.Second):
		chBun <- "Bún đã nấu xong"
	case <-ctx.Done():
		fmt.Println("Cancel!")
		return
	}
}

func main() {

	chPho := make(chan string)
	chBun := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()

	go cookPho(ctx, chPho)
	go cookBun(ctx, chBun)

	for i := 1; i <= 2; i++ {
		select {
		case pho := <-chPho:
			fmt.Println("Nhan duoc: ", pho)
		case bun := <-chBun:
			fmt.Println("Nhan duoc: ", bun)
		case <-ctx.Done():
			fmt.Println("Ngưng nhận món")
			return
		}
	}
}
