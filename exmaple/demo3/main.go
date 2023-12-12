package main

import (
	"fmt"
	"time"
)

func main(){
	//定义一个 int管道
	intChan :=make(chan int,1)
	go func ()  {
		time.Sleep(time.Second)
		intChan<-10
	}()

	stringChan :=make(chan string,1)
	go func ()  {
		time.Sleep(time.Second)
		stringChan<-"ss"	
	}()

	select{
	case v:=<-intChan:
		fmt.Println("intChan",v)
	case v:=<-stringChan:
		fmt.Println("stringChan",v)
	default:
		fmt.Println("防止select被阻塞")
	}
}