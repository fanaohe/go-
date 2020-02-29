package main

import (
	"fmt"
)

type person struct {
	name string
	age int
	gender string
	hobby []string
}

func f(x *person)  {
	x.name = "梦想"
	//fmt.Println(x)    // 说明在go语言中在从外部传参到函数内部，是copy一个副本过来的，底层是说，重新声明了一个变量
	                  // 修改函数内部的值，对函数外部的值不会进行修改

}

//func f2(x *person)  {
//	x.name = "梦想2"
//	//fmt.Println(x)       // 如果想要修改原有数据，需要根据源数据指针进行更改数据
//}



func main()  {
	var p  person
	p.name = "用户名称"
	p.age = 18
	p.gender = "男"
	p.hobby = []string{"篮球","足球"}
	//f2(&p)
	//fmt.Println(p)
	//// 匿名结构体
	//var s struct {
	//	name string
	//	age int
	//}
	//s.name = "名称"
	//s.age = 19
	//fmt.Printf("type:%T, value:%v", s, s)
	// 结构体指针1
	var p2 = new(person)
	p2.name = "理想"

	fmt.Println(p2)
	// 结构体指针2
	var p3 = person{
		name:   "名称",
		age:    10,
		gender: "男",
		hobby:  []string{
			"篮球","乒乓球",
		},
	}
	f(&p3)
	fmt.Printf("%p\n", &p3)
	m := make(map[string][]string, 10)
	m["名称"] = []string{"李云鹏"}


	fmt.Println(m)

	s := [...]int{1,2,3,4,5}

	for _,v := range s{
		fmt.Println(v)
	}


}


