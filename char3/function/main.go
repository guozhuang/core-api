package main

import (
	"fmt"
	"net/http"
)

/*type Data struct {}

//直接使用命名类型来进行方法声明
func (Data) testM(){
//
}

//直接使用命名类型指针进行方法声明
func (*Data) testMPtr(){
//
}

func main(){

	//对上面不声明接收者变量名的方法调用【实践中几乎不会这样声明】
	//因为调用时，只能用方法表达式来使用【就和函数没区别了(虽然底层看来确实没区别)】
	Data.testM(struct{}{})//这种就是方法表达式的写法，传入的第一个参数就是对应方法声明的接收者的类型【因为方法中接收者限定的就是第一个参数】
	Data.testMPtr(struct{}{})//invalid method expression Data.testMPtr (needs pointer receiver: (*Data).testMPtr)
}*/

//函数实现接口并且进行中间层拦截使用【通过装饰器模式实现】
//借助golang的http包实现

//目前先使用函数作一个代理函数【体现的是函数作为参数的传递】
/*func mainHandleFunc (pattern string, handler func(http.ResponseWriter, *http.Request)) {
	fmt.Println("hello")//这就是中间层处理的逻辑，可以进行诸如验签或者log操作
	http.HandleFunc(pattern, handler)
}

func testHandler (w http.ResponseWriter, r *http.Request) {
	//先按照handleFunc标准参数进行设置
	fmt.Fprintf(w, "world")
}

func main(){
	mainHandleFunc("/", testHandler)

	http.ListenAndServe(":8000", nil)
}*/

//pipeline的方式实现中间层

/*func mainHandleFunc (handler func (http.ResponseWriter, *http.Request)) func (http.ResponseWriter, *http.Request) {
	//实现中间层
	fmt.Println("hi")
	return handler
}

func testHandler (w http.ResponseWriter, r *http.Request) {
	//先按照handleFunc标准参数进行设置
	fmt.Fprintf(w, "world")
}

func main (){

	http.HandleFunc("/", mainHandleFunc(testHandler))//使用pipeline的方式来编写
	//依旧表现的是函数作为参数已经返回值来传递的方式:这种方式初始化时已经调用mainhandleFunc

	http.ListenAndServe(":8000", nil)
}*/

//通过声明函数类型，来实现中间层的转发，将对应的handler
type MainHandle func(http.ResponseWriter, *http.Request) //函数命名类型

func (mainF MainHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mainF(w, r)
}

func mainHandleFunc(handler func(http.ResponseWriter, *http.Request)) http.Handler {
	fmt.Println("hi interface")
	return mainHandleFunc(handler) //这里的写法等价于下面，只是下面更加好理解
	//return (MainHandle)(handler) //因为Mainhandle实现来Http的Handler接口
	//问题在于此时的函数类型直接调用的形式难道不是函数调用么？
	//牛逼：此处相当于是函数handler（作为参数传递来的）进行类型转化【而不是函数调用，因为此处只有函数的命名类型，在golang的类型系统中
	//显然是属于类型转换：转换的底层依据是：又因为传入的handler本身的底层结构就是该命名类型的，所以直接能转换】
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	//先按照handleFunc标准参数进行设置
	fmt.Fprintf(w, "world")
}

func main() {
	//需要切换成handle的初始化
	http.Handle("/", mainHandleFunc(testHandler)) //此处使用handle而不是handleFunc
	//当前包相当于实现了http包的handleFunc的逻辑

	http.ListenAndServe(":8000", nil)
}
