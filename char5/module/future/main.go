package main

import (
	"fmt"
	"time"
)

//一次流程中存在多个无依赖调用的并发实现

//demo

//多个无依赖实现的话，结合带缓冲的chan，并且构造一个hash来进行唯一标示来获取相应的结果
//制造一个对应的查询结构来作为请求和记录结果的chan
type query struct {
	//查询参数
	sql chan string

	//结果
	result chan string
}

//demo 伪装查询sql【实际中rpc的调用过程也可以这么操作】
func execQuery(q query) {

	//启动协程
	go func() {
		//
		sql := <-q.sql

		//进行查询

		//输出结果通道
		q.result <- "result from " + sql
	}()
}

//直接使用该协程来进行并发操作
func main() {

	//初始化query
	q := query{
		make(chan string, 1),
		make(chan string, 1),
	}

	go execQuery(q)

	q.sql <- "select * from table"

	time.Sleep(1 * time.Second)

	//q.sql <- "select tablename ...."//此处会形成死锁

	fmt.Println(<-q.result)
}
