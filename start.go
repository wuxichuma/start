package main

import "fmt"

func main() {
	fmt.Print("wuxi\n")
	fmt.Print("wuxi2\n")
	for index := 0; index < 10; index++ {
		fmt.Print(index)
	}
	ch := make(chan int, 2)
	ch <- 2
	fmt.Println(<-ch)
}
