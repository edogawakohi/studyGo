package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d done\n", id)

}

func main() {

	//bai 1
	// fmt.Println("Hello may thang nhoc")
	// go fmt.Println("cuc cu")
	// time.Sleep(2 * time.Second)

	//Bài 2 – Channel một chiều
	//👉 Tạo một goroutine gửi số 10 vào channel, goroutine chính in ra số đó.

	ch := make(chan int)

	go func() {
		ch <- 10
	}()

	go func() {
		fmt.Println("Day la so: ", <-ch)
	}()
	time.Sleep(2 * time.Second)

	/*🧩 Bài 3 – Deadlock

	👉 Viết chương trình gây deadlock*/

	// ch1 := make(chan int)

	// ch1 <- 1 // deadlock bị gây ra do go main đang chờ một goroutine nhận nạp 1 vào ch1,
	// //nhưng không có goroutine nào nhận nên bị block mãi mãi ở đây, gây ra chương trình deadlock

	// go func() {
	// 	fmt.Println(<-ch1)
	// }()

	/*
			🧩 Bài 4 – Channel có buffer

		👉 Tạo channel có buffer 2 phần tử, gửi 2 giá trị và in chúng ra.
	*/

	// ch2 := make(chan int, 2)

	// go func() {
	// 	ch2 <- 10
	// 	ch2 <- 25
	// }()

	// go func() {
	// 	fmt.Println(<-ch2)
	// 	fmt.Println(<-ch2)
	// }()

	// time.Sleep(2 * time.Second)

	/*
			🧩 Bài 5 – WaitGroup=
		👉 Tạo 3 goroutine, mỗi goroutine in “Task X done”, và đảm bảo chương trình chỉ thoát khi tất cả goroutine xong.
	*/

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()
	fmt.Println("Done")
}
