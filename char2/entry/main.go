package main

import (
	"fmt"
)

func testDefer() int {
	a := 0
	//说明defer内的参数传递是在注册时进行的值传递
	defer func(data int) {
		fmt.Println(data) //0
	}(a)

	//但是如果不使用传参的话，就形成闭包的实现
	defer func() {
		fmt.Println(a) //1,因为是对外部的引用
	}()

	a++
	//panic("this is panic")//此处两个defer都正常执行
	//os.Exit(1)//此处两个defer都不会被执行
	return a
}

func main() {
	result := testDefer()
	fmt.Println(result)
}
