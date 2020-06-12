package main

import "fmt"

//接口调用的过程细节
/**
开销在于：
初始化接口变量，以及接口变量的调用方法【itab中存放了对应方法集的数组指针，然后通过这样的委托调用】
整体开销会大一些，但是对性能的影响有限
*/

type Caler interface {
	Add(int, int) int
	Sub(int, int) int
}

type Adder struct {
	Id int
}

func (adder Adder) Add(a, b int) int {
	return a + b
}

func (adder Adder) Sub(a, b int) int {
	return a - b
}

func main() {
	var adder Caler = &Adder{
		Id: 1234,
	}

	fmt.Println(adder.Add(10, 20))
}
