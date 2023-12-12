package main

import (
	"fmt"
	"net"
)
func process(conn net.Conn){
	defer conn.Close()

	for{
		//创建一个切片，准备将读取的数据放入切片
		buf := make([]byte,1024)
		n,err:=conn.Read(buf)
		if err!=nil {
			return
		}
		fmt.Print(string(buf[0:n]))
	}
}

func main(){
	fmt.Println("服务器端启动")
	listerner,err:= net.Listen("tcp","127.0.0.1:8888")
	if err!=nil{
		fmt.Println("服务器端监听失败",err)
		return 
	}
	//循环等待
	for{
		conn,err2:=listerner.Accept()
	if err2!=nil {
		fmt.Println("服务器端接收失败",err2)
		return
	}

	fmt.Println("服务器端接收成功",conn.RemoteAddr().String())
	//准备一个协程，协程处理客户端服务的请求
	go process(conn)
	}
	
}