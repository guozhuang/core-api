package main

import "fmt"

type IStructer interface {
	InitStruct(int)
}

type myStruct struct {
	Age int
}

func (my *myStruct) InitStruct(age int) {
	my.Age = age
}

//直接使用一个函数作为类型
type Show func(int)

func (Show) InitStruct(age int) {
	//
}

/*func main(){
	var in IStructer

	ageMember := &myStruct{
		Age: 18,
	}

	in = ageMember

	in.InitStruct(19)//接口实现的委托调用

	fmt.Println(ageMember.Age)
}*/

//type Assertion
/*func main(){
	//对类型断言进行实现
	var in IStructer

	var funcD Show

	in = funcD

	if _, ok := in.(Show); ok {
		fmt.Println("this is func type")
	}

	var in2 IStructer
	ageMember := &myStruct{}
	ageMember.InitStruct(20)

	in2 = ageMember
	in2.InitStruct(25)

	fmt.Println(ageMember.Age)
	//这里必须匹配对应的指针类型来进行类型断言【注意】
	if o, ok := in2.(*myStruct); ok {
		fmt.Println(o.Age)
	}
}*/

//type switches
func main() {
	var in IStructer

	in = &myStruct{}
	in.InitStruct(20)

	var funcD Show
	in = &funcD //此处使用【指针类型】赋值给接口变量，则判断类型
	in = funcD  //此处接口变量匹配的【值类型】接口变量

	/**接口变量的初始化也进行一次值拷贝
	原理：本身经过初始化的接口的底层结构存在data，如果赋值的类型实例是值类型，那么接口变量的data就是对应值类型的副本
	如果是赋值的类型变量是指针类型，那么对应接口变量的data就是相应指针的副本
	【除了data之外，接口变量的底层结构还存在tab：*itab：存放相应类型以及函数指针（同样跟struct将数据和方法进行正交化设计一样）也可以看作本身就是一个struct】

	//itab结构说明：【5个字段】
	inner：指向接口类型元数据指针
	_type:绑定对应类型元数据指针【动态绑定的核心实现】
	hash:快速进行类型比对
	_:匿名字段
	func：指向对应的类型数据的指针数组【初始化长度是1，但是会动态扩展（编译器负责扩充）】

	itab数据会被编译器存放在可执行文件对应的静态空间中，不会被gc检查并且回收

	_type字段详情：
	当前结构存放的字段包含：类型名称，类型kind（反射时有所体现），类型元数据指针的起始位置，gc信息，函数指针表，size，ptr以及hash（这个字段是itab中冗余在外部方便快速比对）

	inner:接口类型元数据指针细节：
	包含通用的类型、path以及对应的方法集数组【这里数据又存在方法的名称以及方法的type描述的偏移量】


	以上就是interface基础结构的细节内容
	*/

	switch o := in.(type) {
	case *myStruct:
		fmt.Printf("this is struct %T\n", o)
	case *Show:
		fmt.Printf("this is func ptr %T\n", o)
	case Show:
		fmt.Printf("this is func value %T\n", o)
	}
}
