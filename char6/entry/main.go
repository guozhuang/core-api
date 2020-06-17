package main

import (
	"fmt"
	"reflect"
)

//反射demo

type Student struct {
	Name string "姓名"
	Age  int    `a:"111"b:"333"` //tag的填充处理可以用于editor的处理使用
}

func (s Student) method1() {
	//
}

func main() {
	s := Student{}
	rt := reflect.TypeOf(s) //此处获取的接口变量，并且赋值给rt。通过rt来获取基础信息
	fieldName, ok := rt.FieldByName("Name")

	//取tag数据
	if ok {
		fmt.Println(fieldName.Tag)
	}

	fieldAge, ok := rt.FieldByName("Age")

	if ok {
		fmt.Println(fieldAge.Tag.Get("a"))
		fmt.Println(fieldAge.Tag.Get("b"))
	}

	fmt.Println("type_name:", rt.Name())
	fmt.Println("type_NumField:", rt.NumField())
	fmt.Println("type_PkgPath:", rt.PkgPath())
	fmt.Println("type_String:", rt.String()) //具名类型
	//fmt.Println("type_NumMethod:", rt.NumMethod())//为何返回的是0？【似乎应该使用reflect.Value来使用？】

	fmt.Println("type_Kind:", rt.Kind().String()) //kind的使用

	//获取结构体类型的字段名称[此处可以在批量化初始化时使用]
	for i := 0; i < rt.NumField(); i++ {
		fmt.Printf("type.Field[%d].Name:=%v\n", i, rt.Field(i).Name)
	}

	sc := make([]int, 10)
	sc = append(sc, 1, 2, 3)
	sct := reflect.TypeOf(sc)

	//对slice的反射获取
	scet := sct.Elem() //Elem仅能用于array、chan、map、ptr、slice

	fmt.Println("slice element type.Kind() =:", scet.Kind())
	fmt.Printf("slice element type.Kind()=%d\n", scet.Kind())

	fmt.Println("slice element type.Name()", scet.Name())
	fmt.Println("slice type.NumMethod() = ", scet.NumMethod())
}
