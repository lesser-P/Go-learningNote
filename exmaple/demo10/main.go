package main

import (
	"fmt"
	"time"
)

func main(){
	timer:=time.NewTimer(time.Second*3)
	t:= <-timer.C //是一个只读的管道
	fmt.Println(t)
}