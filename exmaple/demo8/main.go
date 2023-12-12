package main

import (
	"fmt"
	"reflect"
)

func main() {
	num :=100
	reValue:=reflect.ValueOf(&num)
	reValue.Elem().SetInt(200)//得到地址对应的值进行修改
	fmt.Println(num)
}