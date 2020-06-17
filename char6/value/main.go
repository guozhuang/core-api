package main

import (
	"fmt"
	"reflect"
)

//使用reflect.Value来实践
//reflect.Value是一个struct【直接返回的就是struct实例】，并且提供了相应的方法来进行使用
//底层结构为：
/*type Value struct {
	typ *rtype//此处就是动态类型的基本结构
	ptr:指向相应值
	flag：标记字段
}*/

type User struct {
	Id   int
	Name string
	Age  int
}

func (user User) String() {
	fmt.Println("User:", user.Id, user.Name, user.Age)
}

//增加动态适配的功能
func Info(o interface{}) {
	v := reflect.ValueOf(o)

	t := v.Type()

	fmt.Println("type:", t.Name())

	//访问传入的字段名、类型和字段值信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		//进行字段的类型检查【既可以捕获对应类型又可以转化为相应类型？】
		//为何此处value可以重新赋值？【并且只能使用:=来赋值？】【因为switch的作用域进行了value的覆盖：但是代码编写中一定注意不这么使用】
		switch valueType := value.(type) {
		case int:
			fmt.Printf("%s: %v = %d\n", field.Name, field.Type, valueType) //
		case string:
			fmt.Printf("%s: %v = %s\n", field.Name, field.Type, valueType)
		default:
			fmt.Printf("%s: %v = %s\n", field.Name, field.Type, valueType)
		}

		/**
		此处反而引出了相应类型断言和类型查询中一直没有注意到的点：
		1.  o := i.(typeName)：判定是否符合该类型，如果符合，则返回相应类型（接口变量i绑定）的实例（指针）【显然，因为等同于传参】
			【这也是为何函数实现接口之后中转的那个函数中这样写的原因】
		2.  switch v := i.(type)
			此处i必须是接口变量，说明i.(type)的写法是确定的
			值得注意的是v这里必然使用:=进行赋值使用，尽量取不同的名称。
			此时的v的值：一旦case匹配成功，v的值就是i绑定的实例的副本，所以上面的例子中，v可以被直接当值使用。
			【而v不是标注了对应类型名称来匹配case，而是v本身就是值。所以golang的switch语句中的case就是检查v的类型（所以被叫做是类型查询）】
			case中，后面跟随的类型可以是接口类型，也可以是具体的类型
			【注意】【判定规则是，实现了case中指定的接口类型，或者i绑定的具体类型匹配case中的具体类型。就走到case内部】
		*/
	}
}

func main() {
	u := User{
		1,
		"gz",
		30,
	}
	Info(u) //将实例进行传输
}
