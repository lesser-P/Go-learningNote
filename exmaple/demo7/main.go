package main

import (
	"fmt"
	"reflect"
)

//空接口没有任何方法，所以可以理解为所有类型都实现了空接口，也可以理解为我们可以把任何一个变量赋给空接口
type Student struct{
	Name string
	Age int
}

func main(){
	stu:=Student{
		Name:"tom",
		Age:20,
	}
	types:= reflect.TypeOf(stu)
	fmt.Println("type=",types)

	value:=reflect.ValueOf(stu)
	fmt.Println("value=",value)
	i2:=value.Interface()//转换回原始的接口类型
	n,flage:=i2.(Student)
	if flage ==true {
		fmt.Println(n.Name,n.Age)
	}

}