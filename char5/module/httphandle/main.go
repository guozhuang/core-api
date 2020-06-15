package main

import (
	"fmt"
	"sync"
)

//匹配一个请求一个goroutine【因为本身请求除了请求内容标准构造之外，就是完全无状态的】
//单次请求匹配一个goroutine也是对应处理http请求的标准

//【核心点其实是将chan设置到struct的实现模式】

//而golang本身的http server也是这样实现的

//而当前示例使用的是：将任务拆分多个task，每个task匹配一个goroutine运行
//示例用来计算100以内的自然数之和

//将每个任务绑定对应的起始值和结果【】
type task struct {
	begin  int
	end    int
	result chan<- int //单向声明结构
}

//单条任务的具体执行，就是在任务指定的范围内，处理求和逻辑
func (t *task) do() {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i
	}

	t.result <- sum //此处的写入的结果就相当于写入resultChan中
}

//形成分阶段task与分片数据的关联关系【也就是task初始化的作用】
func InitTask(taskChan chan<- task, r chan int, p int) {
	qu := p / 10
	mod := p % 10
	high := qu * 10

	for j := 0; j < qu; j++ {
		b := 10*j + 1
		e := 10 * (j + 1)

		tsk := task{
			begin:  b,
			end:    e,
			result: r,
		}

		taskChan <- tsk
	}

	//剩余的最后一个处理task
	if mod != 0 {
		tsk := task{
			begin:  high,
			end:    p,
			result: r,
		}

		taskChan <- tsk
	}

	close(taskChan)
}

//进行任务的分发处理【单个任务匹配一个goroutine】
func DistributeTask(taskChan <-chan task,
	wait *sync.WaitGroup, resultChan chan int) {
	//
	for v := range taskChan {
		wait.Add(1)

		//并且通过wg来实现整体的结构管理
		go ProcessTask(v, wait)
	}

	wait.Wait()
	close(resultChan)
}

func ProcessTask(tsk task, wait *sync.WaitGroup) {
	tsk.do()

	wait.Done()
}

//对结果进行整合
func ProcessResult(resultChan chan int) int {
	sum := 0

	for v := range resultChan {
		//fmt.Println(v)
		sum += v
	}

	return sum
}

func main() {
	//创建任务列队【如果从http的视角来看，可以是拿到对应的请求socket，放入待处理队列】
	//进行结构的基础初始化
	taskChan := make(chan task, 10)

	//结果chan
	resultChan := make(chan int, 10)

	//wg:结合wg来实现同步？
	wait := &sync.WaitGroup{}

	go InitTask(taskChan, resultChan, 100) //限定范围

	go DistributeTask(taskChan, wait, resultChan)

	//resultChan就是被引用到task的result字段中，因为只有具体的do执行时，
	//才会进入t.result <- 写入，所以就是一个最简单的引用副本，只要正常写入【就相当于是给resultChan写入】
	//【实质上当前协程处理的结果并不会绑定到task这个结构体中，而是直接resultChan生效】
	//于是当前协程直接读取resultChan即可拿到最终的结果
	sum := ProcessResult(resultChan)

	fmt.Println(sum)
}
