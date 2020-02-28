package main

import "fmt"

type person struct {
	name string
	age int
}
// 构造函数：约定函数用new开头
// 返回的是结构体的指针，如果在go当中不返回指针，会造成在函数外来回复制的情况，造成滥用内存
// 当结构体比较大的时候尽量使用结构体指针，减少程序的内存开销
func newPerson(name string, age int)  *person{    // 函数使用*返回的是指针
	return &person{    // go语言中&符号返回的是数据的指针
		name: name,
		age:  age,
	}
	
}

func main()  {
	p1 := newPerson("元帅", 10)
	p2 := newPerson("周林", 20)
	fmt.Printf("%p\n",p1)
	fmt.Printf("%p\n",p2)
}
