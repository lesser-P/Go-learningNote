package main

func main(){
	//默认情况下管道是双向的，可读可写
	
	//声明为只读
	var intChan1 chan<- int // 管道只写
	var intChan2 <-chan int // 管道只读
	intChan1=make(chan int,3)
	intChan2=make(chan int,3)
	intChan1<- 100
	intChan2<- 200
}