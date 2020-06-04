package main

import "fmt"

type User struct {
	age  int
	name string
}

func main() {

	ma := make(map[int]User)

	agens := User{
		age:  18,
		name: "agens",
	}

	ma[1] = agens

	agens.age = 19
	//修改值之后重新赋值
	ma[1] = agens

	fmt.Printf("%v\n", ma)
}
