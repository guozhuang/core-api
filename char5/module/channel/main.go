package main

import "fmt"

//管道相关的实践：实质上是pipeline的模式
//因为chan的结构是同样的，所以直接形成整一套的调用链

//显然根据同样chan类型，将chan的返回值重新作为参数来传递，就形成pipeline模式
func chain(in chan int) chan int {
	out := make(chan int)

	go func() {
		for v := range in {
			out <- v + 1
		}

		close(out)
	}()
	return out
}

//无阻塞chan的使用，形成stream
func main() {
	in := make(chan int) //因为是无缓冲的chan，所以可以进行同步（也就是阻塞）

	//形成初始化：逐个填充chan【并且逐个进行写入并且外部拿来使用】
	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}

		close(in)
	}()

	//形成pipeline的调用方式：利用的就是同类型的chan重新作为入参来进行调用
	out := chain(chain(chain(in)))

	//逐个读取
	for v := range out {
		fmt.Println(v)
	}
}
