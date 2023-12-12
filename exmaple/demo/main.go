package main

import (
	"fmt"
	"sync"
	"time"
)

var intChan chan int
var wg sync.WaitGroup

func main(){
	intChan=make(chan int,50)
	wg.Add(2)
	go writeData()
	go readData()
	//主线程阻塞等待协程完成
	wg.Wait()

}

func writeData(){
	defer wg.Done()
	for i := 0; i < 50; i++ {
		intChan <- i
		fmt.Println("writeData",i)
		time.Sleep(time.Second)
	}
	close(intChan)
}

func readData(){
	defer wg.Done()
	for i := 0; i < 50; i++ {
		num:=<-intChan
		fmt.Println("readData",num)
	}
}