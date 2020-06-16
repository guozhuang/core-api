package main

import (
	"context"
	"fmt"
	"time"
)

//context的实际使用
type otherContext struct {
	context.Context
}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			//退出通知
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			fmt.Printf("%s is running\n", name)
			time.Sleep(1 * time.Second)
		}
	}
}

func workWithValue(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			value := ctx.Value("key1").(string)
			fmt.Printf("%s get value=%s\n", name, value)
			time.Sleep(1 * time.Second)
		}
	}
}

//当前主体中就使用来context的三个主要的操作
func main() {

	//先进行基本的ctx构建
	ctxa, cancel := context.WithCancel(context.Background()) //创建root ctx

	go work(ctxa, "work1") //携带ctx创建goroutine

	tm := time.Now().Add(3 * time.Second)
	//进行deadline方式超时设置
	//这里可以看作是下一层的ctx包装【所以从这个ctx开始具备超时机制】【所以标准的输出中work1始终是在running】
	ctxb, _ := context.WithDeadline(ctxa, tm) //或者直接使用超时方法设置

	go work(ctxb, "work2")

	//传递数据
	oc := otherContext{ctxb}
	ctxc := context.WithValue(oc, "key1", "from main")

	go workWithValue(ctxc, "work3")

	time.Sleep(10 * time.Second)

	//进行退出通知
	cancel() //root ctx创建时返回的回调函数

	time.Sleep(5 * time.Second)
	fmt.Println("main stop")
}
