package main

import "fmt"

//工作池
//实践中往往通过工作池来减少性能开销【构建固定数目，使用完成不回收】，以及兼顾并发性能

//任务池内数量
const (
	NUMBER = 10
)

//
type task struct {
	begin  int
	end    int
	result chan<- int
}

//任务处理细节：
func (tsk *task) do() {
	sum := 0

	for i := tsk.begin; i <= tsk.end; i++ {
		sum += i
	}

	tsk.result <- sum
}

func InitTask(taskChan chan task, r chan int, p int) {
	qu := p / 10
	mod := p % 10
	high := qu * 10

	for j := 0; j < qu; j++ {
		b := 10*j + 1
		e := 10 * (j + 1)

		task := task{
			begin:  b,
			end:    e,
			result: r,
		}

		taskChan <- task
	}

	if mod != 0 {
		task := task{
			begin:  high + 1,
			end:    p,
			result: r,
		}

		taskChan <- task
	}

	close(taskChan)
}

//进行绑定【标准实现中，done参数应该前置】
func DistributeTask(taskChan chan task, workers int, done chan struct{}) {
	//进行统一分配:此处逻辑是同步的
	for i := 0; i < workers; i++ {
		go ProcessTask(taskChan, done)
	}
}

//此处的done描述的是当前任务完成的通知【通知外部】
func ProcessTask(taskChan chan task, done chan struct{}) {
	for tsk := range taskChan {
		tsk.do()
	}
	done <- struct{}{}
}

func closeResult(done chan struct{}, resultChan chan int, workers int) {
	for i := 0; i < workers; i++ {
		<-done //逐个输出，并且在此阻塞
	}
	//走到这里说明全部完成了通知
	close(done)
	close(resultChan)
}

func ProcessResult(resultChan chan int) int {
	sum := 0
	for v := range resultChan {
		fmt.Println(v)
		sum += v
	}

	return sum
}

func main() {
	workers := NUMBER

	taskChan := make(chan task, 10)

	resultChan := make(chan int, 10)

	done := make(chan struct{}, 10) //通知退出【需要逐个通知】

	go InitTask(taskChan, resultChan, 100)

	DistributeTask(taskChan, workers, done) //此处的逻辑就可以看作是启动固定数目的工作池
	//区别与之前的此处也是协程机制，而是固定数目进行统一约束，然后在每个内部任务池中进行done的基本通知
	//由此形成固定的生成和统一完成的done通知写入【连接池这种也可以进行设定，外部可以回收，内部出现异常也可以进行通知处理】

	go closeResult(done, resultChan, workers)

	sum := ProcessResult(resultChan)

	fmt.Println(sum)
}
