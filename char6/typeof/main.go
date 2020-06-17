package main

import (
	"fmt"
	"reflect"
)

//reflect.typeOf传入参数的情况
//1.传入具体类型的实例：返回该类型的反射信息
//2.传入绑定来具体类型的接口变量：返回接口绑定的具体类型的信息
//3.传入的就是单纯的接口变量【未绑定具体类型】：返回的是接口的静态信息

type Int int

type A struct {
	a int
}

type B struct {
	b string
}

//定义接口进行反射类型获取
type Ita interface {
	String() string
}

//使某个类型实现相应的接口
func (b B) String() string {
	return b.b
}

func main() {
	//首先验证反射对类型变量的使用【此处说明的是type声明的新类型，虽然底层类型是一致的】
	var a Int = 12
	var b int = 14

	ra := reflect.TypeOf(a)
	rb := reflect.TypeOf(b)

	if ra == rb {
		//判定获取的反射类型是否一致
		fmt.Println("ra == rb")
	} else {
		fmt.Println("ra != rb") //被输出，说明底层类型相同的不同命名类型获得的反射接口类型不一致
		//显然
	}

	fmt.Println(ra.Name()) //获取对应的类型名称
	fmt.Println(rb.Name())

	fmt.Println(ra.Kind()) //获取相应类型的类型【此时底层类型相同获得kind也相同】
	fmt.Println(rb.Kind())

	s1 := A{
		1,
	}
	s2 := B{
		b: "structb",
	}

	rs1 := reflect.TypeOf(s1)
	rs2 := reflect.TypeOf(s2)

	fmt.Println(rs1.Name())
	fmt.Println(rs2.Name())

	fmt.Println(rs1.Kind().String())
	fmt.Println(rs2.Kind().String())

	//对接口进行处理
	ita := new(Ita)  //直接使用接口
	var itb Ita = s2 //此处接口绑定具体实例

	var itc Ita = &s2 //利用&操作符返回相应指针的形式

	//对接口类型反射输出
	fmt.Println(reflect.TypeOf(ita).Elem().Name())
	fmt.Println(reflect.TypeOf(ita).Elem().Kind().String()) //返回interface

	//对绑定类型的接口输出
	fmt.Println(reflect.TypeOf(itb).Name())
	fmt.Println(reflect.TypeOf(itb).Kind().String()) //返回struct
	//注意绑定类型的接口变量不能直接使用Elem来查看

	fmt.Println(reflect.TypeOf(itc).Elem().Name()) //但是因为ptr可以用elem访问，
	// 所以此处就可以
}
