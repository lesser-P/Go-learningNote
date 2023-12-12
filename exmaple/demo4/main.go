package main

import (
	"fmt"
	"time"
)

func printNum(){
	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func devide(){
	defer func(){
		error:=recover()
		if error !=nil {
			fmt.Println("error",error)
		}
	}()
	num1:=10
	num2:=0
	result:=num1/num2
	fmt.Println(result)

}

func main(){
	go printNum()
	go devide()

	time.Sleep(time.Second*5)
}