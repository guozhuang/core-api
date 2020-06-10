package main

import "fmt"

type person struct {
	age  int
	name string
}

func New(name string) *person {
	person := new(person)
	person.name = name

	return person
}

func main() {
	person := New("张三")

	fmt.Printf("%s\n", person.name)
	/*a := person{
		age: 18,
		name: "hello",
	}

	//此处就是命名类型实例化的变量
	fmt.Printf("%T\n", a)//main.person

	b := struct {
		age int
		name string
	}{
		age:17,
		name:"hello",
	}

	//此处就是非命名类型的打印结果
	fmt.Printf("%T\n", b)//struct { age int; name string }*/
}
