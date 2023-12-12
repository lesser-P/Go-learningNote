package main

import (
	"fmt"
	"reflect"
)
type Student struct{
	Name string
	Age int
}

func main() {
	stu :=Student{
		Name:"ye",
		Age:11,
	}
	value:=reflect.ValueOf(stu)

	value.Method(0).Call(nil)
	n:=value.Elem().NumField()
	fmt.Println(n)
	value.Elem().Field(0).SetString("zhangsan")
	fmt.Println(value)
}

func (s Student) Print(){
	fmt.Println("调用了Print()方法")
}