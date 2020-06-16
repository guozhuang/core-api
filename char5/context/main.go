package main

import "time"

//对context原理进行描述
//痛点：因为golang中的goroutine本身是平行的，不存在父子关系。【显然MPG模型中，P上挂在的G是平行的】
//但是在比较复杂的并发模型中，显然并发单元应该是能够被组织起来的。【例如可以进行退出通知回收处理的】
//context的核心作用就是退出通知和元数据传递

//而在编程模型中的作用就是用来跟踪goroutine调用，在其内部维护一个调用树，在这个链条上传递通知和元数据。

//context包的整体工作机制：】
//第一个创建context的goroutine被成为root节点。
//【该节点负责创建一个实现Context接口的具体对象，并将该对象作为参数传递到其新拉起的goroutine】
//下游节点可以继续封装该对象，再传递到更下面一层
//由此形成了一条整体的链条，通知消息就能够从root一直传递到下游。
//【显然，就像自己常见的task对象，内部字段带chan类型】【这样传递的实例就能起到相应通知的作用】

//context基本数据结构：
//context接口：
type Context interface {
	//超时设置
	Deadline(deadline time.Time, ok bool)

	//完成通知
	Done() <-chan struct{}

	//出错查询
	Err() error

	//上下游传递的数据访问
	Value(key interface{}) interface{}
}

//canceler接口
//canceler接口是context的扩展接口：限定取消通知的context需要实现的方法
type canceler interface {
	//调用cancel方法之后，即可通知后面创建的ctx进行退出
	cancel(removeFromParent bool, err error)

	Done() <-chan struct{} //写入该chan，即可进行取消通知
}

//（1）empty context结构：实现了context接口的空对象，作为root节点的ctx
//从root创建出的goroutine通过具体化包装ctx来形成一个具备相应功能的context实例
//进而具备了相应能力的树状结构。

//（2）cancelCtx 实现了context接口之外，还实现了canceler接口【具备了退出设置】

//（3）timeCtx 实现了context接口的实例，内部封装了cancelCtx类型，并且具备deadline变量【实现定时退出】

//（4）valueCtx：实现了context接口，并且提供了String和Value方法来实现传递信息的公共获取

//context的操作相关：
//实际中context实例在整个调用的链路中，都是作为第一个参数来传递
//每一层中，如果需要对context进行包装，则调用相应的函数进行处理，并且接着传递。
//这些函数有：Background, TODO,
//下面包含具体业务的创建新的子context的函数
//（1）WithCancel(包装成带退出的context)
//（2）WithDeadline(包装成带终止通知)
//（3）WithTimeout（包装成带超时通知）
//（4）WithValue（包装成可以传递数据）//这里的kv结构是并发安全的【但是不应该拿来传递业务数据】【往往传递退出信号，日志信息，调试信息】
//由上面的基本操作来进行ctx树的构建
