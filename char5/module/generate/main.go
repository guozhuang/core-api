package main

import (
	"fmt"
	"math/rand"
	"time"
)

//生成器的解耦实现
//常见的标准随机数生成器的标准构造

//最简单的生成器
//直接生成对应结果的chan进行返回
/*func generateSimple() chan int {
	ch := make (chan int, 10)

	//内部实现协程:协程调用
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

func main(){
	ch := generateSimple()

	//于是使用时，持续进行运行获取，一直没有退出main，所以能持续读取，并且生成器可以持续生成结果
	for {
		fmt.Println(<-ch)
	}
	//持续的阻塞，导致生成器的标准方法能够一直被调度，用来同步对应的chan的数据
	//这也就是通过通道来同步状态
}*/

//多路生成器：扇入
/*func generateIntA() chan int{
	ch := make(chan int, 10)

	go func (){
		for {
			ch <- rand.Int()
		}
	}()

	return ch
}

func generateIntB() chan int {
	ch := make (chan int, 10)

	go func (){
		for {
			ch <- rand.Int()
		}
	}()

	return ch
}

//进行合并生成【也就是扇入思想】
func generateInt() chan int{
	ch := make (chan int, 20)

	go func() {
		for {
			select {
			case ch <- <-generateIntA()://将其中一路的chan中取数据，并且写入总体chan中
			case ch <- <-generateIntB():
			}
		}
	}()

	return ch
}

func main(){
	ch := generateInt()

	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)
	}
}*/

//生成器接收对应的退出信号，来进行生成退出
/*func generateA(done chan struct{}) chan int{
	ch := make (chan int, 10)

	go func() {
	Label:
		for {
			select {
			case <- done:
				break Label
			case ch <- rand.Int():
			}
		}
		close(ch)
	}()

	return ch
}


func main(){
	done := make(chan struct{})

	ch := generateA(done)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	close(done)//进行关闭广播

	for v := range ch {
		fmt.Println(v)
	}
}*/

//最终一个较为完备的生成器工具[结合缓冲，并发，以及扇入思想和退出机制]

func generateIntA(done chan struct{}) chan int {
	ch := make(chan int, 10)

	go func() {
	Label:
		for {
			select {
			case <-done:
				fmt.Println("done generateA")
				break Label
			case ch <- rand.Int():
			}
		}

		close(ch)
	}()
	return ch
}

func generateIntB(done chan struct{}) chan int {
	ch := make(chan int, 10)

	go func() {
	Label:
		for {
			select {
			case <-done:
				fmt.Println("done generateB")
				break Label
			case ch <- rand.Int():
			}
		}

		close(ch)
	}()
	return ch
}

func generateInt(done chan struct{}) chan int {
	ch := make(chan int, 20)

	sendChild := make(chan struct{})
	go func() {
	Label:
		for {
			select {
			case <-done:
				sendChild <- struct{}{} //struct{}的chan不占任何内存，就是专门用来close的
				sendChild <- struct{}{}
				fmt.Println("done generate")
				break Label
			case ch <- <-generateIntA(sendChild):
			case ch <- <-generateIntB(sendChild):
			}
		}

		close(ch)
	}()

	return ch
}

func main() {
	done := make(chan struct{})

	ch := generateInt(done)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)
	}

	close(done)

	fmt.Println("stop generate")

	//形成相应的回收逻辑
	/**【因为当前main存在持续运行的结果，所以能正常打印，如果直接退出，则没有后面的goroutine的运行打印】
	当前示例下，停止的generate的打印结果为：
	stop generate
	done generate
	done generateA
	done generateB
	*/

	for {
		time.Sleep(1 * time.Millisecond)
	}
}
