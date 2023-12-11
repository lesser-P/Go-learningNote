package main

import "fmt"

// func main() {
// 	fmt.Println("Hello World")
// 	var intChan chan int
// 	intChan = make(chan int, 3)

// 	//向管道存放数据
// 	intChan <- 10
// 	num := 20
// 	intChan <- num
// 	num1 := <- intChan
// 	num2 := <- intChan
// 	fmt.Println(len(intChan), cap(intChan))
// 	fmt.Println(num1)
// 	fmt.Println(num2)

// }

func main(){
	var intChan chan int 
	intChan=make(chan int,100)
	for i := 0; i < 100; i++ {
		intChan<- i
	}
	close(intChan)
	for e := range intChan {
		fmt.Println(e)
	}
}