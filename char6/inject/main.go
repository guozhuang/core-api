package main

//实践使用inject，也是对反射实现的核心注意点
//inject源码包只有200行，需要分析这一实现来近一步理解。
//目前似乎看来主要是依赖原生的反射方法，然后实现注入map形式来进行挂载

import (
	"fmt"
	"github.com/codegangsta/inject"
)

//首先使用inject来实现函数的依赖注入
type S1 interface{}

type S2 interface{}

type Staff struct {
	Name    string `inject`
	Company S1     `inject` //注意这里声明的接口类型，注入时需要匹配：因为前面已经有string，同样类型会被直接覆盖掉，所以声明空接口来区别
	Level   S2     `inject`
	Age     int    `inject`
}

func main() {
	//使用struct注入的方式才是需要重点关注
	s := Staff{}

	inj := inject.New()
	inj.Map("tom")
	inj.MapTo("tencent", (*S1)(nil))
	inj.MapTo("T4", (*S2)(nil))
	inj.Map(23)

	inj.Apply(&s) //相当于将s指针进行重新实例化【似乎应该使用这一方式来进行处理，每个组件加载先使用这一方式进行注入】

	fmt.Printf("s=%v\n", s) //s={tom tencent T4 23}
}

/*func Format(name string, company S1, level S2, age int) {
	fmt.Printf("name = %s, company=%s, level=%s, age=%d\n", name, company, level, age)
}

func main (){
	inj := inject.New()

	inj.Map("tom")
	inj.MapTo("tencent", (*S1)(nil))
	inj.MapTo("T4", (*S2)(nil))
	inj.Map(23)

	inj.Invoke(Format)//实现了函数的委托调用【但是实际上直接使用call的函数式编程也能轻松实现这一逻辑】
	//name = tom, company=tencent, level=T4, age=23
}*/
