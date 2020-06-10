package main

import "fmt"

type IDemo interface {
	demo()
}

type Demo struct {
	//
}

//此处使用指针接收者，对应的接口变量赋值时，就只能传入指针【需要&或者new来进行创建】
func (d *Demo) demo() {
	fmt.Println("demo")
}

func main() {
	var d IDemo

	var dStruct Demo

	//d = dStruct//此处编译不通过
	d = &dStruct

	d.demo()
}
