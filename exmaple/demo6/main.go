package main

import (
	"fmt"
	"reflect"
)

//空接口没有任何方法，所以可以理解为所有类型都实现了空接口，也可以理解为我们可以把任何一个变量赋给空接口

func testReflect(i interface{}){
	value:= reflect.ValueOf(i)
	fmt.Println("value=",value)
	types:=reflect.TypeOf(i)
	fmt.Println("type=",types)
}

func main(){
	//对基本数据类型进行反射
	// 定义一个基本数据类型

	var num int=100
	//通过反射获取变量的类型和值
	testReflect(num)
	value:= reflect.ValueOf(num)
	fmt.Println("num=",value)

	num2:=value.Int()+100
	fmt.Println("num2=",num2)
	
}