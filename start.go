package main

import "fmt"
import "rsc.io/quote"
func main() {
	fmt.Print("wuxi\n")
	fmt.Print("wuxi2\n")
	for index := 0; index < 10; index++ {
		fmt.Print(index)
	}
	ch := make(chan int, 2)
	ch <- 3
	fmt.Println(<-ch)

	fmt.Println(quote.Hello())
}
