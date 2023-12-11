package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main(){
	fmt.Println("客户端启动")
	conn,err := net.Dial("tcp","127.0.0.1:8888")

	if err!=nil{
		fmt.Println("客户端连接失败",err)
		return
	}
	fmt.Println("连接成功，conn",conn.RemoteAddr().String())

	//通过客户端发送单行数据，然后退出
	reader :=bufio.NewReader(os.Stdin)//os.Stdin代表终端标准输入

	str,err:=reader.ReadString('\n')
	if err!=nil {
		fmt.Println("读取失败",err)
	}
	n,err:=conn.Write([]byte(str))
	if err!=nil {
		fmt.Println("连接失败",err)
	}
	fmt.Println("客户端发送了",n,"个字节的数据")
}